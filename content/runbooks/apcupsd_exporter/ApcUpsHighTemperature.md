---
title: ApcUpsHighTemperature
description: Troubleshooting for alert ApcUpsHighTemperature
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsHighTemperature

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Internal temperature is high ({{$value}}°C)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ApcUpsHighTemperature
expr: apcupsd_internal_temperature_celsius >= 40
for: 2m
labels:
    severity: warning
annotations:
    summary: APC UPS high temperature (instance {{ $labels.instance }})
    description: |-
        Internal temperature is high ({{$value}}°C)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ApcUpsHighTemperature

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
