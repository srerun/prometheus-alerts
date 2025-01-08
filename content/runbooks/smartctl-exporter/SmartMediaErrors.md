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


Here is a sample runbook for the Prometheus alert rule:

## Meaning
This alert is triggered when the `smartctl_device_media_errors` metric exceeds 0, indicating that the monitored device has experienced media errors. This is a critical alert as it may indicate a potential failure of the storage device, leading to data loss or corruption.

## Impact
The impact of this alert is high, as media errors can cause:

* Data loss or corruption
* System crashes or instability
* Downtime and reduced productivity
* Potential loss of critical business data

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the device logs for any error messages related to the media errors.
2. Run the `smartctl` command on the affected device to gather more detailed information about the errors.
3. Verify that the device is properly configured and that the firmware is up-to-date.
4. Check the device's SMART (Self-Monitoring, Analysis and Reporting Technology) attributes to determine the cause of the media errors.

## Mitigation
To mitigate the issue, follow these steps:

1. Immediately backup critical data to prevent potential data loss.
2. Restart the affected device to attempt to recover from the error.
3. Run a thorough diagnostic test on the device using `smartctl` or other diagnostic tools.
4. Consider replacing the device if the errors persist or if the device is approaching its end-of-life.
5. Perform regular maintenance on the device, such as firmware updates and disk checks, to prevent future errors.

Note: This is just a sample runbook and may need to be customized to fit your specific use case and environment.