---
title: KubernetesCronjobSuspended
description: Troubleshooting for alert KubernetesCronjobSuspended
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesCronjobSuspended

## Meaning
[//]: # "Short paragraph that explains what the alert means"
CronJob {{ $labels.namespace }}/{{ $labels.cronjob }} is suspended

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: KubernetesCronjobSuspended
expr: kube_cronjob_spec_suspend != 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes CronJob suspended ({{ $labels.namespace }}/{{ $labels.cronjob }})
    description: |-
        CronJob {{ $labels.namespace }}/{{ $labels.cronjob }} is suspended
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/KubernetesCronjobSuspended

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
