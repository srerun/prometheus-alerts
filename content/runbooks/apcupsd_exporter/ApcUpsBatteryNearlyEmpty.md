---
title: ApcUpsBatteryNearlyEmpty
description: Troubleshooting for alert ApcUpsBatteryNearlyEmpty
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsBatteryNearlyEmpty

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Battery is almost empty (< 10% left)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ApcUpsBatteryNearlyEmpty
expr: apcupsd_battery_charge_percent < 10
for: 0m
labels:
    severity: critical
annotations:
    summary: APC UPS Battery nearly empty (instance {{ $labels.instance }})
    description: |-
        Battery is almost empty (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ApcUpsBatteryNearlyEmpty

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
