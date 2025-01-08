---
title: ApcUpsLowBatteryVoltage
description: Troubleshooting for alert ApcUpsLowBatteryVoltage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsLowBatteryVoltage

Battery voltage is lower than nominal (< 95%)

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsLowBatteryVoltage" %}}

{{% comment %}}

```yaml
alert: ApcUpsLowBatteryVoltage
expr: (apcupsd_battery_volts / apcupsd_battery_nominal_volts) < 0.95
for: 0m
labels:
    severity: warning
annotations:
    summary: APC UPS low battery voltage (instance {{ $labels.instance }})
    description: |-
        Battery voltage is lower than nominal (< 95%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsLowBatteryVoltage.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
