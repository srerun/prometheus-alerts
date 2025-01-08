---
title: ApcUpsLessThan15MinutesOfBatteryTimeRemaining
description: Troubleshooting for alert ApcUpsLessThan15MinutesOfBatteryTimeRemaining
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsLessThan15MinutesOfBatteryTimeRemaining

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Battery is almost empty (< 15 Minutes remaining)

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsLessThan15MinutesOfBatteryTimeRemaining" %}}

{{% comment %}}

```yaml
alert: ApcUpsLessThan15MinutesOfBatteryTimeRemaining
expr: apcupsd_battery_time_left_seconds < 900
for: 0m
labels:
    severity: critical
annotations:
    summary: APC UPS Less than 15 Minutes of battery time remaining (instance {{ $labels.instance }})
    description: |-
        Battery is almost empty (< 15 Minutes remaining)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsLessThan15MinutesOfBatteryTimeRemaining.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
