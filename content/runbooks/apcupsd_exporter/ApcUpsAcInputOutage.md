---
title: ApcUpsAcInputOutage
description: Troubleshooting for alert ApcUpsAcInputOutage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApcUpsAcInputOutage

UPS now running on battery (since {{$value | humanizeDuration}})

<details>
  <summary>Alert Rule</summary>

{{% rule "apc-ups/apcupsd_exporter.yml" "ApcUpsAcInputOutage" %}}

{{% comment %}}

```yaml
alert: ApcUpsAcInputOutage
expr: apcupsd_battery_time_on_seconds > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: APC UPS AC input outage (instance {{ $labels.instance }})
    description: |-
        UPS now running on battery (since {{$value | humanizeDuration}})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/apcupsd_exporter/ApcUpsAcInputOutage.md

```

{{% /comment %}}

</details>


Here is a runbook for the ApcUpsAcInputOutage alert rule:

## Meaning
The ApcUpsAcInputOutage alert is triggered when the APC UPS (Uninterruptible Power Supply) detects a loss of AC input power and switches to battery mode. This means that the UPS is no longer receiving power from the primary AC source and is now running on battery backup.

## Impact
The impact of this alert is that the system or equipment connected to the UPS may experience a power outage if the battery backup is depleted before the AC power is restored. This can lead to data loss, system downtime, and other disruptions.

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the UPS status: Verify the UPS status using the apcupsd command-line tool or the web interface to confirm that the UPS is indeed running on battery power.
2. Check the AC input: Verify that the AC input power is not present or is faulty.
3. Check the battery level: Check the remaining battery life to estimate how much time is left before the UPS shuts down.
4. Check the system logs: Review system logs to identify any other errors or issues that may be related to the power outage.

## Mitigation
To mitigate the impact of this alert, follow these steps:

1. Investigate and correct the AC input issue: Identify and resolve the issue causing the loss of AC input power.
2. Switch to a backup power source: If available, switch to a backup power source, such as a generator or another UPS, to minimize downtime.
3. Notify stakeholders: Inform stakeholders of the potential power outage and estimated downtime.
4. Prepare for possible shutdown: Prepare for a possible shutdown of the system or equipment connected to the UPS if the battery backup is depleted.

Remember to follow your organization's procedures for handling power outages and UPS failures.