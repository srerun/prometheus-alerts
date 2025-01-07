---
title: JenkinsRunFailureTotal
description: Troubleshooting for alert JenkinsRunFailureTotal
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsRunFailureTotal

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Job run failures: ({{$value}}) {{$labels.jenkins_job}}. Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: JenkinsRunFailureTotal
expr: delta(jenkins_runs_failure_total[1h]) > 100
for: 0m
labels:
    severity: warning
annotations:
    summary: Jenkins run failure total (instance {{ $labels.instance }})
    description: |-
        Job run failures: ({{$value}}) {{$labels.jenkins_job}}. Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/JenkinsRunFailureTotal

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
