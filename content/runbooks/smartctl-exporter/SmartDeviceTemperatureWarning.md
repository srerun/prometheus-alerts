---
title: SmartDeviceTemperatureWarning
description: Troubleshooting for alert SmartDeviceTemperatureWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SmartDeviceTemperatureWarning

Device temperature  warning (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "smart-device-monitoring/smartctl-exporter.yml" "SmartDeviceTemperatureWarning" %}}

{{% comment %}}

```yaml
alert: SmartDeviceTemperatureWarning
expr: smartctl_device_temperature > 60
for: 2m
labels:
    severity: warning
annotations:
    summary: Smart device temperature warning (instance {{ $labels.instance }})
    description: |-
        Device temperature  warning (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/smartctl-exporter/SmartDeviceTemperatureWarning.md

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
