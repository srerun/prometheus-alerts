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

The SmartNvmeWearoutIndicator alert is triggered when the available spare blocks on an NVMe device falls below a certain threshold. This indicates that the device is worn out and needs to be replaced to prevent data loss or corruption.

## Impact

* Data loss or corruption due to worn-out NVMe device
* Potential downtime or system unavailability
* Reduced performance and increased latency

## Diagnosis

1. Check the affected instance and NVMe device using the `instance` and `device` labels.
2. Verify the current available spare blocks value using the `smartctl_device_available_spare` metric.
3. Check the device's wear level and health status using other smartctl metrics, such as `smartctl_device_wear_level` and `smartctl_device_health_status`.
4. Consult the device's documentation and manufacturer's guidelines for replacement or maintenance.

## Mitigation

1. Immediately replace the worn-out NVMe device to prevent data loss or corruption.
2. Perform a thorough backup of critical data to ensure business continuity.
3. Consider migrating to a more reliable or redundant storage solution.
4. Monitor the device's health status and wear level regularly to prevent future wear-out issues.
5. Update the `smartctl_device_available_spare_threshold` value if necessary, based on the device's manufacturer's guidelines and your organization's requirements.

Note: For more detailed steps and specific instructions, refer to the linked runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/smartctl-exporter/SmartNvmeWearoutIndicator.md