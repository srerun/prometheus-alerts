---
title: SmartCriticalWarning
description: Troubleshooting for alert SmartCriticalWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SmartCriticalWarning

device has critical warning (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "smart-device-monitoring/smartctl-exporter.yml" "SmartCriticalWarning" %}}

{{% comment %}}

```yaml
alert: SmartCriticalWarning
expr: smartctl_device_critical_warning > 0
for: 15m
labels:
    severity: critical
annotations:
    summary: Smart critical warning (instance {{ $labels.instance }})
    description: |-
        device has critical warning (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/smartctl-exporter/SmartCriticalWarning.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The `SmartCriticalWarning` alert is triggered when the `smartctl_device_critical_warning` metric indicates that a critical warning has been detected on a disk drive. This warning is generated by the SMART (Self-Monitoring, Analysis and Reporting Technology) system, which is a built-in feature of most modern hard drives. A critical warning typically indicates a serious issue with the drive, such as a high temperature, bad sectors, or other potential failures.

## Impact

If this alert is triggered, it is essential to investigate and take immediate action to prevent data loss or system downtime. A critical warning on a disk drive can lead to:

* Data loss or corruption
* System crashes or instability
* Downtime for critical services or applications
* Potential hardware failure

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the alert annotations for the specific error message and device instance.
2. Review the SMART logs to identify the specific warning or error that triggered the alert.
3. Check the system logs for any related errors or warnings.
4. Run `smartctl` commands to gather more information about the drive's health.
5. Verify that the drive is not already in a failed state.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately backup critical data from the affected drive to a safe location.
2. Identify the underlying cause of the critical warning and take corrective action (e.g., replace the drive, adjust system settings, etc.).
3. Run `smartctl` commands to clear the warning and reset the drive's health status (if possible).
4. Monitor the drive's health closely to ensure the issue is resolved.
5. Consider scheduling a maintenance window to replace the drive if it is deemed necessary.

Remember to always follow your organization's specific procedures and guidelines for handling critical warnings on disk drives.