package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/magicx-ai/groq-go/groq"
	"gopkg.in/yaml.v3"
)

type Rules struct {
	Groups []struct {
		Name  string `yaml:"name"`
		Rules []Rule `yaml:"rules"`
	} `yaml:"groups"`
}

type Rule struct {
	Source string `yaml:"-"`
	Alert  string `yaml:"alert"`
	Expr   string `yaml:"expr"`
	For    string `yaml:"for"`
	Labels struct {
		Severity string `yaml:"severity"`
	} `yaml:"labels"`
	Annotations *struct {
		Summary     string `yaml:"summary"`
		Description string `yaml:"description"`
		Runbook     string `yaml:"runbook,omitempty"`
	} `yaml:"annotations"`
}

type Data struct {
	Path    string
	Rule    Rule
	Content string
}

const groqModel groq.ModelID = "llama-3.3-70b-versatile"

// const filename = "blackbox.yaml"
const runbook_url = "https://srerun.github.io/prometheus-alerts/runbooks"
const runbookPath = "content/runbooks"
const rulesPath = "rules"
const contentTmpl = `## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
`

//go:embed template.tmpl
var templates embed.FS
var groqKey = os.Getenv("GROQ_API_KEY")

func main() {
	tmpl, err := template.New("template.tmpl").Funcs(template.FuncMap{
		"to_yaml": ToYaml, "first_line": FirstLine,
	}).Delims("{{{", "}}}").ParseFS(templates, "template.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if groqKey == "" {
		slog.Warn("GROQ_API_KEY is not set.  AI-generated runbooks will not be created.")
	}

	fmt.Printf("Reading rules files in %s\n", rulesPath)
	files, err := filepath.Glob(fmt.Sprintf("%s/*/*.yml", rulesPath))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("found %d files\n", len(files))
	for _, filename := range files {
		fd, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		content, err := io.ReadAll(fd)
		if err != nil {
			log.Fatal(err)
		}
		rules := &Rules{}
		if err := yaml.Unmarshal(content, rules); err != nil {
			log.Fatal(err)
		}
		path := buildPathname(filename)
		for _, group := range rules.Groups {
			for _, rule := range group.Rules {
				rulePath := filepath.Dir(filename)
				rule.Source = fmt.Sprintf("%s/%s", buildPathname(rulePath), filepath.Base(filename))
				rule.Annotations.Runbook = fmt.Sprintf("%s/%s/%s/", runbook_url, path, strings.ToLower(rule.Alert))
				createRunbook(path, tmpl, rule)
			}
		}
		out, err := yaml.Marshal(rules)
		if err != nil {
			log.Fatal(err)
		}
		fd.Close()
		fd, err = os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		fd.Write(out)
		fd.Close()
		time.Sleep(2 * time.Second)
	}
}

// createRunbook will create a new runbook based on the alert
func createRunbook(path string, t *template.Template, rule Rule) {
	filepath := fmt.Sprintf("%s/%s", runbookPath, path)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if err = os.MkdirAll(filepath, 0755); err != nil {
			slog.Error("could not create diretory", "path", filepath, "error", err)
			return
		}
	}
	if _, err := os.Stat(fmt.Sprintf("%s/_index.md", filepath)); os.IsNotExist(err) {
		createIndex(filepath, path)
	}
	filename := fmt.Sprintf("%s/%s.md", filepath, rule.Alert)
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return
	}
	fmt.Printf("Adding runbook for %s\n", rule.Alert)

	out, err := os.Create(filename)
	if err != nil {
		slog.Error("could not create file", "filename", filename, "error", err)
		return
	}
	defer out.Close()
	if err = t.Execute(out, Data{path, rule, generateContent(rule)}); err != nil {
		slog.Error("error rendering template", "filename", filename, "error", err)
	}
}

func createIndex(path string, section string) {
	out, err := os.Create(fmt.Sprintf("%s/_index.md", path))
	if err != nil {
		slog.Error("could not create _index.md", "error", err)
		return
	}
	defer out.Close()
	out.WriteString(fmt.Sprintf("---\ntitle: %s\n", section))
	out.WriteString(`
bookCollapseSection: true
bookFlatSection: true
weight: 1
---
`)

}

func buildPathname(filename string) string {
	file := strings.Split(filename, ".")[0]
	paths := strings.Split(file, "/")
	return paths[len(paths)-1]
}

func ToYaml(data interface{}) (string, error) {
	// implementation to convert date string to Unix timestamp
	yml, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(yml), nil
}

func FirstLine(lines string) (string, error) {
	return strings.Split(lines, "\n")[0], nil
}

func generateContent(rule Rule) string {
	if groqKey == "" {
		return contentTmpl
	}
	ruleB, err := yaml.Marshal(rule)
	if err != nil {
		return contentTmpl
	}
	cli := groq.NewClient(groqKey, &http.Client{})

	req := groq.ChatCompletionRequest{
		Messages: []groq.Message{
			{
				Role: "user",
				Content: fmt.Sprintf(`Write a runbook using level 2 headers (##) for the sections with the 
				sections meaning, impact, diagnosis, and mitigation for the following prometheus alert rule.
				%s`, string(ruleB)),
			},
		},
		Model:       groqModel,
		MaxTokens:   1500,
		Temperature: 1,
		TopP:        1,
		NumChoices:  1,
		Stream:      false,
	}

	var resp *groq.ChatCompletionResponse
	for retries := 0; retries < 3; retries++ {
		resp, err = cli.CreateChatCompletion(req)
		if err != nil {
			if strings.Contains(err.Error(), "invalid status code: 429") {
				fmt.Println(err.Error())
				sleepTime := retries + 5
				fmt.Printf("Rate Limit exceeded...Sleeping for %d seconds", sleepTime)
				time.Sleep(time.Duration(sleepTime) * time.Second)
				fmt.Printf("\nretrying...attempt #%d\n", retries+2)
				continue
			}
			fmt.Println(fmt.Errorf("error occurred: %v", err))
			return contentTmpl
		} else {
			return resp.Choices[0].Message.Content
		}
	}
	log.Fatal("Retries exceeded")
	return contentTmpl
}
