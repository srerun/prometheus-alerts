---
title: ApcUpsLowBatteryVoltage
description: Troubleshooting for alert ApcUpsLowBatteryVoltage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsLowBatteryVoltage

Battery voltage is lower than nominal (< 95%)

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsLowBatteryVoltage" %}}

{{% comment %}}

```yaml
alert: ApcUpsLowBatteryVoltage
expr: (apcupsd_battery_volts / apcupsd_battery_nominal_volts) < 0.95
for: 0m
labels:
    severity: warning
annotations:
    summary: APC UPS low battery voltage (instance {{ $labels.instance }})
    description: |-
        Battery voltage is lower than nominal (< 95%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsLowBatteryVoltage.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the ApcUpsLowBatteryVoltage alert rule:

## Meaning

The ApcUpsLowBatteryVoltage alert is triggered when the battery voltage of an APC UPS device falls below 95% of its nominal voltage. This indicates that the battery is not fully charged and may not be able to provide sufficient power in the event of a power failure.

## Impact

If this alert is not addressed, the APC UPS device may not be able to provide sufficient power to critical systems in the event of a power failure, leading to downtime and potential data loss.

## Diagnosis

To diagnose the issue, check the APC UPS device's battery voltage and charging status using the `apcupsd` command-line tool or the web interface. Verify that the battery is properly connected and that there are no issues with the charging circuit.

Check the Prometheus metrics `apcupsd_battery_volts` and `apcupsd_battery_nominal_volts` to confirm that the battery voltage is indeed below 95% of its nominal value.

## Mitigation

To mitigate this issue, perform the following steps:

1. Check the APC UPS device's battery health and charging status to ensure that it is functioning properly.
2. Ensure that the battery is properly connected and that there are no issues with the charging circuit.
3. If the battery is old or damaged, consider replacing it with a new one.
4. Verify that the APC UPS device is properly configured and that the battery is set to charge correctly.
5. Consider increasing the frequency of battery testing and maintenance to prevent future issues.
6. If the issue persists, consider contacting the APC UPS device's manufacturer or a qualified technician for further assistance.