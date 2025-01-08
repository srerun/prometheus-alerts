---
title: HostUnusualDiskReadLatency
description: Troubleshooting for alert HostUnusualDiskReadLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualDiskReadLatency

Disk latency is growing (read operations > 100ms)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostUnusualDiskReadLatency" %}}

{{% comment %}}

```yaml
alert: HostUnusualDiskReadLatency
expr: (rate(node_disk_read_time_seconds_total[1m]) / rate(node_disk_reads_completed_total[1m]) > 0.1 and rate(node_disk_reads_completed_total[1m]) > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host unusual disk read latency (instance {{ $labels.instance }})
    description: |-
        Disk latency is growing (read operations > 100ms)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostUnusualDiskReadLatency.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the `HostUnusualDiskReadLatency` alert:

## Meaning

The `HostUnusualDiskReadLatency` alert is triggered when the average disk read latency on a host exceeds 100ms. This indicates that disk I/O operations are taking longer than usual to complete, which can impact system performance and responsiveness.

## Impact

High disk read latency can cause:

* Slow application performance
* Increased load times
* Decreased system responsiveness
* Potential data loss or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the system logs for any errors or warnings related to disk I/O operations.
2. Verify that the disk is not experiencing any hardware issues, such as high disk usage or hardware failures.
3. Check the disk queue length and disk utilization to identify if there are any bottlenecks.
4. Investigate if there are any resource-intensive applications or processes running on the system that may be contributing to the high disk read latency.

## Mitigation

To mitigate the issue, follow these steps:

1. Check and clean up any unnecessary files or data on the disk to free up space.
2. Consider upgrading to faster disk hardware, such as SSDs, if possible.
3. Optimize disk configuration and settings, such as adjusting disk scheduling algorithms or enabling write caching.
4. Identify and optimize resource-intensive applications or processes that may be contributing to the high disk read latency.
5. Consider implementing disk I/O monitoring and alerting to catch potential issues before they become critical.

Additionally, review system configuration and tuning options, such as adjusting disk timeouts or queue lengths, to optimize disk performance.