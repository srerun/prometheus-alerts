---
title: JenkinsBuildTestsFailing
description: Troubleshooting for alert JenkinsBuildTestsFailing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsBuildTestsFailing

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Last build tests failed: {{$labels.jenkins_job}}. Failed build Tests for job `{{$labels.jenkins_job}}` on {{$labels.instance}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

{{% rule "jenkins/metric-plugin.yml" "JenkinsBuildTestsFailing" %}}

{{% comment %}}

```yaml
alert: JenkinsBuildTestsFailing
expr: default_jenkins_builds_last_build_tests_failing > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Jenkins build tests failing (instance {{ $labels.instance }})
    description: |-
        Last build tests failed: {{$labels.jenkins_job}}. Failed build Tests for job `{{$labels.jenkins_job}}` on {{$labels.instance}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsBuildTestsFailing.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
