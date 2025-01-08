---
title: NetdataLowDiskSpace
description: Troubleshooting for alert NetdataLowDiskSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataLowDiskSpace

Netdata low disk space (> 80%)

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataLowDiskSpace" %}}

{{% comment %}}

```yaml
alert: NetdataLowDiskSpace
expr: 100 / netdata_disk_space_GB_average * netdata_disk_space_GB_average{dimension=~"avail|cached"} < 20
for: 5m
labels:
    severity: warning
annotations:
    summary: Netdata low disk space (instance {{ $labels.instance }})
    description: |-
        Netdata low disk space (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataLowDiskSpace.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the NetdataLowDiskSpace alert:

## Meaning

The NetdataLowDiskSpace alert is triggered when the available disk space on a Netdata instance falls below 20%. This means that the disk space is critically low, and immediate action is required to prevent data loss or system crashes.

## Impact

The impact of low disk space on a Netdata instance can be severe:

* Data loss: If the disk space is completely depleted, Netdata may not be able to write new data, leading to data loss.
* System crashes: Low disk space can cause system crashes or freezes, leading to downtime and disruption to monitoring and alerting capabilities.
* Performance degradation: Low disk space can also cause performance degradation, leading to slower response times and decreased system performance.

## Diagnosis

To diagnose the cause of the low disk space, follow these steps:

1. Check the disk usage: Run the command `df -h` to check the disk usage on the Netdata instance.
2. Identify the cause: Check the output of the `df -h` command to identify the cause of the low disk space. Common causes include:
	* Log file growth: Check the log file sizes and rotate or delete unnecessary logs.
	* Data file growth: Check the sizes of data files and consider archiving or deleting unnecessary data.
	* Software updates: Check if any software updates have caused the disk space to decrease.
3. Check the Netdata configuration: Check the Netdata configuration to ensure that it is set up correctly and that disk space is being monitored correctly.

## Mitigation

To mitigate the low disk space issue, follow these steps:

1. Free up disk space: Immediately free up disk space by deleting unnecessary files, logs, and data.
2. Configure log rotation: Configure log rotation to ensure that log files are regularly rotated and deleted.
3. Configure data retention: Configure data retention policies to ensure that unnecessary data is deleted or archived.
4. Increase disk space: Consider increasing the disk space on the Netdata instance or moving the data to a larger storage device.
5. Monitor disk space: Closely monitor the disk space on the Netdata instance to prevent similar issues in the future.