---
title: SmartMediaErrors
description: Troubleshooting for alert SmartMediaErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SmartMediaErrors

device has media errors (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "smart-device-monitoring/smartctl-exporter.yml" "SmartMediaErrors" %}}

{{% comment %}}

```yaml
alert: SmartMediaErrors
expr: smartctl_device_media_errors > 0
for: 15m
labels:
    severity: critical
annotations:
    summary: Smart media errors (instance {{ $labels.instance }})
    description: |-
        device has media errors (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/smartctl-exporter/SmartMediaErrors.md

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
