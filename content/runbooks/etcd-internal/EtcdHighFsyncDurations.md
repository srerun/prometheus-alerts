---
title: EtcdHighFsyncDurations
description: Troubleshooting for alert EtcdHighFsyncDurations
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighFsyncDurations

Etcd WAL fsync duration increasing, 99th percentile is over 0.5s

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighFsyncDurations" %}}

{{% comment %}}

```yaml
alert: EtcdHighFsyncDurations
expr: histogram_quantile(0.99, rate(etcd_disk_wal_fsync_duration_seconds_bucket[1m])) > 0.5
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high fsync durations (instance {{ $labels.instance }})
    description: |-
        Etcd WAL fsync duration increasing, 99th percentile is over 0.5s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighFsyncDurations.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the EtcdHighFsyncDurations alert:

## Meaning

The EtcdHighFsyncDurations alert is triggered when the 99th percentile of etcd's WAL fsync duration exceeds 0.5 seconds over a 1-minute period. This indicates that etcd is experiencing high latency when writing to disk, which can impact the overall performance and reliability of the system.

## Impact

The impact of high fsync durations can be significant, leading to:

* Slow etcd write performance
* Increased latency for etcd operations
* Potential for etcd to fall behind, leading to inconsistencies and errors
* Increased risk of data loss or corruption in the event of a failure

## Diagnosis

To diagnose the root cause of high fsync durations, follow these steps:

1. Check etcd disk usage and available space to ensure that the disk is not full or nearly full.
2. Verify that the disk is properly configured and healthy (e.g., check for disk errors, firmware issues, etc.).
3. Investigate system resource utilization (CPU, memory, I/O) to identify potential bottlenecks.
4. Check etcd logs for errors or warnings related to disk I/O or fsync operations.
5. Verify that etcd is properly configured and optimized for the underlying storage system.

## Mitigation

To mitigate the effects of high fsync durations, follow these steps:

1. Immediately investigate and address any underlying disk or system resource issues.
2. Consider increasing etcd's disk write buffer size to reduce the frequency of fsync operations.
3. Optimize etcd's configuration for the underlying storage system (e.g., adjust sync interval, etc.).
4. Consider migrating etcd to a faster storage system or optimizing the existing storage configuration.
5. Implement additional monitoring and alerting to detect potential issues before they become critical.