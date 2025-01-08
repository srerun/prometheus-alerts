---
title: HostNodeOvertemperatureAlarm
description: Troubleshooting for alert HostNodeOvertemperatureAlarm
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNodeOvertemperatureAlarm

Physical node temperature alarm triggered

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostNodeOvertemperatureAlarm" %}}

{{% comment %}}

```yaml
alert: HostNodeOvertemperatureAlarm
expr: ((node_hwmon_temp_crit_alarm_celsius == 1) or (node_hwmon_temp_alarm == 1)) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: critical
annotations:
    summary: Host node overtemperature alarm (instance {{ $labels.instance }})
    description: |-
        Physical node temperature alarm triggered
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostNodeOvertemperatureAlarm.md

```

{{% /comment %}}

</details>


Here is a runbook for the `HostNodeOvertemperatureAlarm` alert rule:

## Meaning

The `HostNodeOvertemperatureAlarm` alert is triggered when a node's temperature exceeds a critical or warning threshold, indicating a potential overheating issue. This alert is critical because high temperatures can cause damage to the node's hardware, leading to system failures or data loss.

## Impact

* The node may experience reduced performance or shut down unexpectedly, leading to service disruptions and potential data loss.
* Prolonged high temperatures can cause permanent damage to the node's hardware, requiring costly repairs or replacement.
* If the node is part of a critical service, the downtime can have significant business and financial impacts.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the node's temperature readings in Prometheus to determine the current temperature and trend.
2. Verify that the temperature sensor is functioning correctly and not reporting false values.
3. Investigate recent changes to the node's environment, such as changes in air flow, cooling systems, or nearby heat sources.
4. Check the node's system logs for any error messages related to temperature or cooling systems.
5. Perform a visual inspection of the node to ensure that it is properly ventilated and that all fans are functioning correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately shut down any non-essential workloads on the node to reduce heat generation.
2. Verify that the node's cooling system is functioning correctly and adjust as needed.
3. Implement temporary cooling measures, such as directing cool air towards the node or using temporary cooling devices.
4. Schedule a maintenance window to inspect and clean the node's air vents and fans.
5. Consider relocating the node to a cooler location or providing additional cooling solutions, such as air conditioning or liquid cooling systems.
6. Monitor the node's temperature readings closely to ensure that the issue is resolved and prevent further overheating events.