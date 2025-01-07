---
title: SmartCriticalWarning
description: Troubleshooting for alert SmartCriticalWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SmartCriticalWarning

## Meaning
[//]: # "Short paragraph that explains what the alert means"
device has critical warning (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SmartCriticalWarning
expr: smartctl_device_critical_warning > 0
for: 15m
labels:
    severity: critical
annotations:
    summary: Smart critical warning (instance {{ $labels.instance }})
    description: |-
        device has critical warning (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/SmartCriticalWarning

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
