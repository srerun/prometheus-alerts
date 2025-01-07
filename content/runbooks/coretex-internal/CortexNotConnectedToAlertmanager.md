---
title: CortexNotConnectedToAlertmanager
description: Troubleshooting for alert CortexNotConnectedToAlertmanager
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexNotConnectedToAlertmanager

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cortex not connected to Alertmanager (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CortexNotConnectedToAlertmanager
expr: cortex_prometheus_notifications_alertmanagers_discovered < 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Cortex not connected to Alertmanager (instance {{ $labels.instance }})
    description: |-
        Cortex not connected to Alertmanager (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CortexNotConnectedToAlertmanager

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
