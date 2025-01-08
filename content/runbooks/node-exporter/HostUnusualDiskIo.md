---
title: HostUnusualDiskIo
description: Troubleshooting for alert HostUnusualDiskIo
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualDiskIo

Time spent in IO is too high on {{ $labels.instance }}. Check storage for issues.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostUnusualDiskIo" %}}

{{% comment %}}

```yaml
alert: HostUnusualDiskIo
expr: (rate(node_disk_io_time_seconds_total[1m]) > 0.5) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 5m
labels:
    severity: warning
annotations:
    summary: Host unusual disk IO (instance {{ $labels.instance }})
    description: |-
        Time spent in IO is too high on {{ $labels.instance }}. Check storage for issues.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostUnusualDiskIo.md

```

{{% /comment %}}

</details>


## Meaning

The `HostUnusualDiskIo` alert is triggered when the rate of disk I/O time on a host exceeds 0.5 seconds per minute for a sustained period of 5 minutes. This indicates that the disk I/O subsystem on the host is experiencing high levels of utilization, which can lead to performance degradation and potential outages.

## Impact

* High disk I/O utilization can cause slower response times and increased latency for applications and services running on the host.
* Prolonged high disk I/O can lead to disk failures, data corruption, and potential data loss.
* This can result in downtime, revenue loss, and a negative impact on business operations.

## Diagnosis

* Investigate the host's disk usage and I/O patterns to identify the root cause of the high disk I/O.
* Check the disk storage configuration, including disk type, capacity, and usage.
* Review system logs for errors or warnings related to disk I/O or storage subsystems.
* Use tools like `iostat` or `iotop` to analyze disk I/O statistics and identify which processes or applications are contributing to the high disk I/O.
* Verify that disk storage is properly configured and optimized for the host's workload.

## Mitigation

* Immediately investigate and address any disk storage issues, such as low disk space, disk failures, or misconfigured storage settings.
* Identify and terminate or throttle any resource-intensive processes or applications contributing to high disk I/O.
* Consider upgrading or replacing disk storage to improve performance and capacity.
* Implement disk I/O optimization techniques, such as disk caching, spindle allocation, or storage tiering.
* Consider migrating workloads to alternative storage solutions, such as cloud-based storage or distributed file systems.
* Verify that system logs are properly configured and monitored to detect and alert on disk I/O issues.
* Update the alert threshold or adjust the alerting rules as needed based on further analysis and investigation.