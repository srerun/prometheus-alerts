---
title: JenkinsBuildsHealthScore
description: Troubleshooting for alert JenkinsBuildsHealthScore
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JenkinsBuildsHealthScore

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})

<details>
  <summary>Alert Rule</summary>

{{% rule "jenkins/metric-plugin.yml" "JenkinsBuildsHealthScore" %}}

{{% comment %}}

```yaml
alert: JenkinsBuildsHealthScore
expr: default_jenkins_builds_health_score < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Jenkins builds health score (instance {{ $labels.instance }})
    description: |-
        Healthcheck failure for `{{$labels.instance}}` in realm {{$labels.realm}}/{{$labels.env}} ({{$labels.region}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/metric-plugin/JenkinsBuildsHealthScore.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
