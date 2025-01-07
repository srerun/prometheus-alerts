---
title: ArgocdServiceUnhealthy
description: Troubleshooting for alert ArgocdServiceUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ArgocdServiceUnhealthy

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Service {{ $labels.name }} run by argo is currently not healthy.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ArgocdServiceUnhealthy
expr: argocd_app_info{health_status!="Healthy"} != 0
for: 15m
labels:
    severity: warning
annotations:
    summary: ArgoCD service unhealthy (instance {{ $labels.instance }})
    description: |-
        Service {{ $labels.name }} run by argo is currently not healthy.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ArgocdServiceUnhealthy

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
