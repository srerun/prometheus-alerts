package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Rules struct {
	Groups []struct {
		Name  string `yaml:"name"`
		Rules []Rule `yaml:"rules"`
	} `yaml:"groups"`
}

type Rule struct {
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

// const filename = "blackbox.yaml"
const runbook_url = "https://github.com/srerun/prometheus-alerts/content/runbooks"
const runbookPath = "/Users/tdavis/src/srerun/prometheus-alerts/content/runbooks"
const rulesPath = "/Users/tdavis/src/srerun/prometheus-alerts/rules"

//go:embed template.tmpl
var templates embed.FS

func main() {
	tmpl, err := template.New("template.tmpl").Funcs(template.FuncMap{
		"to_yaml": ToYaml, "first_line": FirstLine,
	}).Delims("{{{", "}}}").ParseFS(templates, "template.tmpl")
	if err != nil {
		log.Fatal(err)
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
				fmt.Printf("Adding runbook for %s\n", rule.Alert)
				rule.Annotations.Runbook = fmt.Sprintf("%s/%s", runbook_url, rule.Alert)
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
		createIndex(filepath, path)
	}
	filename := fmt.Sprintf("%s/%s.md", filepath, rule.Alert)
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return
	}
	out, err := os.Create(filename)
	if err != nil {
		slog.Error("could not create file", "filename", filename, "error", err)
		return
	}
	defer out.Close()
	if err = t.Execute(out, rule); err != nil {
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
