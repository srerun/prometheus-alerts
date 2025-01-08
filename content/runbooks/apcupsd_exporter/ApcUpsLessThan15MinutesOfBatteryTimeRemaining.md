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


Here is a runbook for the Prometheus alert rule:

## Meaning

This alert is triggered when the APC UPS battery time remaining falls below 15 minutes, indicating that the battery is almost empty. This is a critical alert as it can lead to unexpected downtime or data loss if the UPS is unable to provide power to the system.

## Impact

* Unexpected downtime or data loss due to UPS failure
* Potential loss of business-critical services or operations
* Increased risk of equipment damage or data corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the UPS status and battery levels using the apcupsd exporter dashboard or CLI.
2. Verify that the UPS is properly configured and receiving power from the grid.
3. Check for any signs of UPS malfunction or failure, such as alarms or error messages.
4. Review system logs for any errors or warnings related to the UPS or power supply.

## Mitigation

To mitigate the issue, follow these steps:

1. **Immediately switch to backup power source**: If possible, switch to a backup power source, such as a generator or secondary UPS, to prevent downtime or data loss.
2. **Notify system administrators and stakeholders**: Inform relevant teams and stakeholders of the issue and expected downtime, if applicable.
3. **Replace or repair the UPS**: Arrange for the UPS to be replaced or repaired as soon as possible to ensure continued power supply to the system.
4. **Perform a thorough system check**: Once the UPS is replaced or repaired, perform a thorough system check to ensure that all systems are functioning correctly and data integrity is maintained.

Remember to also investigate the root cause of the issue to prevent future occurrences.