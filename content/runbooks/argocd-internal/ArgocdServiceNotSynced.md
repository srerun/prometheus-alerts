---
title: ArgocdServiceNotSynced
description: Troubleshooting for alert ArgocdServiceNotSynced
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ArgocdServiceNotSynced

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Service {{ $labels.name }} run by argo is currently not in sync.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ArgocdServiceNotSynced
expr: argocd_app_info{sync_status!="Synced"} != 0
for: 15m
labels:
    severity: warning
annotations:
    summary: ArgoCD service not synced (instance {{ $labels.instance }})
    description: |-
        Service {{ $labels.name }} run by argo is currently not in sync.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ArgocdServiceNotSynced

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
