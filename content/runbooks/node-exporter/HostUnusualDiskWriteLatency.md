---
title: HostUnusualDiskWriteLatency
description: Troubleshooting for alert HostUnusualDiskWriteLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualDiskWriteLatency

Disk latency is growing (write operations > 100ms)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostUnusualDiskWriteLatency" %}}

{{% comment %}}

```yaml
alert: HostUnusualDiskWriteLatency
expr: (rate(node_disk_write_time_seconds_total[1m]) / rate(node_disk_writes_completed_total[1m]) > 0.1 and rate(node_disk_writes_completed_total[1m]) > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host unusual disk write latency (instance {{ $labels.instance }})
    description: |-
        Disk latency is growing (write operations > 100ms)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostUnusualDiskWriteLatency.md

```

{{% /comment %}}

</details>


Here is a runbook for the `HostUnusualDiskWriteLatency` alert rule:

## Meaning

This alert is triggered when the average disk write latency on a host exceeds 100ms. This indicates that the disk is taking longer than usual to complete write operations, which can impact the performance and responsiveness of applications running on the host.

## Impact

High disk write latency can cause:

* Slow application responses
* Delayed data processing
* Increased latency in database queries
* Reduced system throughput
* Potential data loss or corruption in extreme cases

## Diagnosis

To diagnose the cause of high disk write latency, follow these steps:

1. Check the disk usage and available space on the host to ensure it is not running low on disk space.
2. Investigate the disk I/O usage and identify any resource-intensive processes or applications that may be contributing to the high latency.
3. Verify that the disk is not experiencing any hardware issues, such as failures or bad sectors.
4. Check the system logs for any errors or warnings related to disk I/O or storage.
5. Review the node's configuration and ensure that it is properly optimized for disk I/O performance.

## Mitigation

To mitigate the effects of high disk write latency, follow these steps:

1. Free up disk space by deleting unnecessary files or expanding the disk capacity.
2. Identify and optimize resource-intensive processes or applications to reduce their disk I/O usage.
3. Consider replacing the disk with a faster one, such as an SSD, if the current disk is old or slow.
4. Implement disk I/O caching or other optimization techniques to improve disk performance.
5. Monitor the host's disk I/O performance and adjust the system configuration as needed to prevent future occurrences of high disk write latency.