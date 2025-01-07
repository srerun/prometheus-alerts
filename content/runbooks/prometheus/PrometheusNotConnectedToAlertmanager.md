---
title: PrometheusNotConnectedToAlertmanager
description: Troubleshooting for alert PrometheusNotConnectedToAlertmanager
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusNotConnectedToAlertmanager

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus cannot connect the alertmanager

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusNotConnectedToAlertmanager
expr: prometheus_notifications_alertmanagers_discovered < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus not connected to alertmanager (instance {{ $labels.instance }})
    description: |-
        Prometheus cannot connect the alertmanager
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/PrometheusNotConnectedToAlertmanager

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
