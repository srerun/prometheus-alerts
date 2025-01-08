---
title: ApcUpsBatteryNearlyEmpty
description: Troubleshooting for alert ApcUpsBatteryNearlyEmpty
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsBatteryNearlyEmpty

Battery is almost empty (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsBatteryNearlyEmpty" %}}

{{% comment %}}

```yaml
alert: ApcUpsBatteryNearlyEmpty
expr: apcupsd_battery_charge_percent < 10
for: 0m
labels:
    severity: critical
annotations:
    summary: APC UPS Battery nearly empty (instance {{ $labels.instance }})
    description: |-
        Battery is almost empty (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsBatteryNearlyEmpty.md

```

{{% /comment %}}

</details>


Here is a runbook for the ApcUpsBatteryNearlyEmpty alert:

## Meaning

The ApcUpsBatteryNearlyEmpty alert indicates that the battery level of an APC UPS (Uninterruptible Power Supply) has fallen below 10%. This critical alert suggests that the UPS is at risk of shutting down due to low battery power, which can lead to unexpected downtime and potential data loss.

## Impact

If left unattended, a low battery level can cause the UPS to shut down, leading to:

* Unexpected downtime of critical systems and services
* Potential data loss or corruption
* Disruption of business operations and revenue loss
* Increased risk of hardware damage due to sudden power loss

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the current battery level and voltage: Use the `apcupsd` command-line tool or a monitoring dashboard to check the current battery level and voltage.
2. Verify the UPS is properly connected and configured: Ensure the UPS is correctly connected to the device it is supposed to protect, and that all cables are securely connected.
3. Check for any alarms or error messages: Review the UPS logs and event history to identify any error messages or alarm conditions that may be contributing to the low battery level.
4. Perform a self-test: Initiate a self-test on the UPS to verify its functionality and identify any potential issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Replace the battery: If the battery is old or has been in use for an extended period, consider replacing it with a new one.
2. Perform a battery calibration: Calibrate the battery to ensure it is accurately reporting its charge level.
3. Reduce the load: Identify any unnecessary devices connected to the UPS and disconnect them to reduce the load and conserve battery power.
4. Implement a backup plan: Develop a backup plan to ensure business continuity in the event of an extended power outage or UPS failure.
5. Schedule a maintenance window: Schedule a maintenance window to perform further diagnostic checks and repairs as needed.

Remember to always follow proper safety precautions when working with electrical equipment, and consult the UPS user manual or manufacturer's instructions if you are unsure about any of the mitigation steps.