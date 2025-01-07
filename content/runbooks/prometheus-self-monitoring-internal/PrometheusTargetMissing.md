---
title: PrometheusTargetMissing
description: Troubleshooting for alert PrometheusTargetMissing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTargetMissing

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A Prometheus target has disappeared. An exporter might be crashed.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusTargetMissing
expr: up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus target missing (instance {{ $labels.instance }})
    description: |-
        A Prometheus target has disappeared. An exporter might be crashed.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PrometheusTargetMissing

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
