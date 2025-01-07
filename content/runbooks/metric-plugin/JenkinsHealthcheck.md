---
title: JenkinsHealthcheck
description: Troubleshooting for alert JenkinsHealthcheck
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsHealthcheck

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Jenkins healthcheck score: {{$value}}. Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: JenkinsHealthcheck
expr: jenkins_health_check_score < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Jenkins healthcheck (instance {{ $labels.instance }})
    description: |-
        Jenkins healthcheck score: {{$value}}. Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/JenkinsHealthcheck

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
