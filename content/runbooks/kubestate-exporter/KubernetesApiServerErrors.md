---
title: KubernetesApiServerErrors
description: Troubleshooting for alert KubernetesApiServerErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesApiServerErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Kubernetes API server is experiencing high error rate

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: KubernetesApiServerErrors
expr: sum(rate(apiserver_request_total{job="apiserver",code=~"(?:5..)"}[1m])) by (instance, job) / sum(rate(apiserver_request_total{job="apiserver"}[1m])) by (instance, job) * 100 > 3
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes API server errors (instance {{ $labels.instance }})
    description: |-
        Kubernetes API server is experiencing high error rate
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/KubernetesApiServerErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
