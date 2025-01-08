---
title: ApcUpsHighLoad
description: Troubleshooting for alert ApcUpsHighLoad
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsHighLoad

UPS load is > 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsHighLoad" %}}

{{% comment %}}

```yaml
alert: ApcUpsHighLoad
expr: apcupsd_ups_load_percent > 80
for: 0m
labels:
    severity: warning
annotations:
    summary: APC UPS high load (instance {{ $labels.instance }})
    description: |-
        UPS load is > 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsHighLoad.md

```

{{% /comment %}}

</details>


Here is a runbook for the ApcUpsHighLoad alert rule:

## Meaning

The ApcUpsHighLoad alert is triggered when the load on an APC UPS (Uninterruptible Power Supply) exceeds 80%. This indicates that the UPS is being utilized at a level that may compromise its ability to provide reliable power to connected devices in the event of a power outage.

## Impact

A high load on the UPS can lead to reduced battery life, increased risk of power outages, and potential data loss or corruption. This can have significant consequences for critical systems and services that rely on the UPS for power backup.

## Diagnosis

To diagnose the cause of the high load on the UPS:

1. Review the `apcupsd_ups_load_percent` metric to determine the current load on the UPS.
2. Investigate the devices connected to the UPS to identify any high-power draws or unusual usage patterns.
3. Check the UPS's power capacity and ensure it is sufficient to support the connected devices.
4. Verify that the UPS is properly configured and maintained, including ensuring that the batteries are functioning correctly.

## Mitigation

To mitigate the high load on the UPS:

1. Identify and address any high-power draws from connected devices, such as servers or storage systems.
2. Consider upgrading the UPS to a model with higher power capacity or adding additional UPS units to distribute the load.
3. Implement power-saving measures, such as reducing the power consumption of connected devices or implementing energy-efficient practices.
4. Schedule a maintenance window to perform UPS battery replacement or other necessary maintenance tasks.

Remember to review the alert annotations for specific details about the affected instance and load percentage to inform your diagnosis and mitigation efforts.