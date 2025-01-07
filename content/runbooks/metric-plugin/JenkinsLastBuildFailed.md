---
title: JenkinsLastBuildFailed
description: Troubleshooting for alert JenkinsLastBuildFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsLastBuildFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Last build failed: {{$labels.jenkins_job}}. Failed build for job `{{$labels.jenkins_job}}` on {{$labels.instance}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: JenkinsLastBuildFailed
expr: default_jenkins_builds_last_build_result_ordinal == 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Jenkins last build failed (instance {{ $labels.instance }})
    description: |-
        Last build failed: {{$labels.jenkins_job}}. Failed build for job `{{$labels.jenkins_job}}` on {{$labels.instance}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/JenkinsLastBuildFailed

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
