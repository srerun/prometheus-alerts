---
title: HostOutOfDiskSpace
description: Troubleshooting for alert HostOutOfDiskSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostOutOfDiskSpace

Disk is almost full (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostOutOfDiskSpace" %}}

{{% comment %}}

```yaml
alert: HostOutOfDiskSpace
expr: ((node_filesystem_avail_bytes * 100) / node_filesystem_size_bytes < 10 and ON (instance, device, mountpoint) node_filesystem_readonly == 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host out of disk space (instance {{ $labels.instance }})
    description: |-
        Disk is almost full (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostOutOfDiskSpace.md

```

{{% /comment %}}

</details>


## Meaning

The HostOutOfDiskSpace alert is triggered when a host's available disk space falls below 10% and the file system is not read-only. This alert indicates that the host is at risk of running out of disk space, which can cause service disruptions, data loss, and system crashes.

## Impact

If left unaddressed, a host running out of disk space can lead to:

* Service disruptions and downtime
* Data loss or corruption
* System crashes or freezes
* Increased latency and decreased performance
* Inability to write logs, causing debugging and troubleshooting difficulties

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the disk usage of the affected host using tools like `df` or `du`.
2. Identify the file system that is running low on disk space.
3. Investigate the cause of the disk space usage, such as:
	* Large log files or directories.
	* Unused or unnecessary files and directories.
	* Disk-hungry applications or processes.
4. Check the Prometheus metrics, such as `node_filesystem_avail_bytes` and `node_filesystem_size_bytes`, to confirm the disk space usage.

## Mitigation

To mitigate the issue, follow these steps:

1. Free up disk space by deleting unnecessary files and directories.
2. compress or rotate log files to reduce their size.
3. Identify and terminate disk-hungry applications or processes.
4. Consider increasing the disk space or migrating data to a larger storage solution.
5. Implement disk space monitoring and alerting to catch low disk space issues before they become critical.
6. Schedule regular disk cleanups and maintenance tasks to prevent disk space issues from recurring.

Additional resources:

* Refer to the official Node Exporter documentation for more information on disk space metrics.
* Consult the Prometheus Alerting documentation for configuration and customization options.
* Review the runbook for this alert on GitHub for more detailed instructions and troubleshooting steps.