---
title: SmartNvmeWearoutIndicator
description: Troubleshooting for alert SmartNvmeWearoutIndicator
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SmartNvmeWearoutIndicator

NVMe device is wearing out (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "smart-device-monitoring/smartctl-exporter.yml" "SmartNvmeWearoutIndicator" %}}

{{% comment %}}

```yaml
alert: SmartNvmeWearoutIndicator
expr: smartctl_device_available_spare{device=~"nvme.*"} < smartctl_device_available_spare_threshold{device=~"nvme.*"}
for: 15m
labels:
    severity: critical
annotations:
    summary: Smart NVME Wearout Indicator (instance {{ $labels.instance }})
    description: |-
        NVMe device is wearing out (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/smartctl-exporter/SmartNvmeWearoutIndicator.md

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
