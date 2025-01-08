---
title: ApcUpsAcInputOutage
description: Troubleshooting for alert ApcUpsAcInputOutage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsAcInputOutage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
UPS now running on battery (since {{$value | humanizeDuration}})

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsAcInputOutage" %}}

{{% comment %}}

```yaml
alert: ApcUpsAcInputOutage
expr: apcupsd_battery_time_on_seconds > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: APC UPS AC input outage (instance {{ $labels.instance }})
    description: |-
        UPS now running on battery (since {{$value | humanizeDuration}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsAcInputOutage.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
