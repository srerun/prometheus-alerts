---
title: EtcdHighCommitDurations
description: Troubleshooting for alert EtcdHighCommitDurations
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighCommitDurations

Etcd commit duration increasing, 99th percentile is over 0.25s

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighCommitDurations" %}}

{{% comment %}}

```yaml
alert: EtcdHighCommitDurations
expr: histogram_quantile(0.99, rate(etcd_disk_backend_commit_duration_seconds_bucket[1m])) > 0.25
for: 2m
labels:
    severity: warning
annotations:
    summary: Etcd high commit durations (instance {{ $labels.instance }})
    description: |-
        Etcd commit duration increasing, 99th percentile is over 0.25s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighCommitDurations.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `EtcdHighCommitDurations`:

## Meaning

The `EtcdHighCommitDurations` alert is triggered when the 99th percentile of etcd commit durations exceeds 0.25 seconds over a 1-minute period. This indicates that etcd is taking longer than expected to commit data to disk, which can lead to performance issues and increased latency in the system.

## Impact

A high etcd commit duration can have several negative impacts on the system:

* Increased latency: As etcd takes longer to commit data, requests may be delayed, leading to increased latency and slower system response times.
* Performance degradation: High commit durations can lead to reduced system performance, as etcd may become bottlenecked by slow disk I/O.
* Data consistency issues: In extreme cases, high commit durations can lead to data consistency issues, as etcd may struggle to keep up with the rate of incoming requests.

## Diagnosis

To diagnose the root cause of high etcd commit durations, follow these steps:

1. Check etcd disk usage: Verify that etcd has sufficient disk space available. High disk usage can cause slow commit durations.
2. Investigate disk I/O performance: Check disk I/O performance metrics to determine if the disk is experiencing high latency or throughput issues.
3. Review etcd configuration: Verify that etcd is configured correctly, including settings such as `sync-commit` and `fsync-duration`.
4. Check for system resource contention: Verify that etcd has sufficient system resources (e.g., CPU, memory) and is not contending with other processes for resources.
5. Check etcd logs: Review etcd logs for errors or warnings related to commit durations.

## Mitigation

To mitigate high etcd commit durations, follow these steps:

1. Increase disk space: If disk usage is high, consider increasing available disk space or implementing disk cleanup mechanisms.
2. Optimize disk I/O performance: Consider upgrading disk hardware or implementing disk I/O optimization techniques, such as caching or parallelizing I/O operations.
3. Adjust etcd configuration: Consider adjusting etcd configuration settings, such as `sync-commit` and `fsync-duration`, to optimize commit performance.
4. Resource allocation: Ensure that etcd has sufficient system resources (e.g., CPU, memory) and consider adjusting resource allocation if necessary.
5. Monitor and analyze commit durations: Continue to monitor and analyze commit durations to identify trends and patterns, and adjust mitigation strategies as needed.