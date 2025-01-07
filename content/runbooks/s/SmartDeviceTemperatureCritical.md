---
title: SmartDeviceTemperatureCritical
description: Troubleshooting for alert SmartDeviceTemperatureCritical
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SmartDeviceTemperatureCritical

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Device temperature critical  (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SmartDeviceTemperatureCritical
expr: smartctl_device_temperature > 80
for: 2m
labels:
    severity: critical
annotations:
    summary: Smart device temperature critical (instance {{ $labels.instance }})
    description: |-
        Device temperature critical  (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/SmartDeviceTemperatureCritical

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
