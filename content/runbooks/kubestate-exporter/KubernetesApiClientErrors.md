---
title: KubernetesApiClientErrors
description: Troubleshooting for alert KubernetesApiClientErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesApiClientErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Kubernetes API client is experiencing high error rate

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: KubernetesApiClientErrors
expr: (sum(rate(rest_client_requests_total{code=~"(4|5).."}[1m])) by (instance, job) / sum(rate(rest_client_requests_total[1m])) by (instance, job)) * 100 > 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes API client errors (instance {{ $labels.instance }})
    description: |-
        Kubernetes API client is experiencing high error rate
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/KubernetesApiClientErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
