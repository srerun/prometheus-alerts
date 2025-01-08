---
title: CephHighOsdLatency
description: Troubleshooting for alert CephHighOsdLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephHighOsdLatency

Ceph Object Storage Daemon latency is high. Please check if it doesn't stuck in weird state.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephHighOsdLatency" %}}

{{% comment %}}

```yaml
alert: CephHighOsdLatency
expr: ceph_osd_perf_apply_latency_seconds > 5
for: 1m
labels:
    severity: warning
annotations:
    summary: Ceph high OSD latency (instance {{ $labels.instance }})
    description: |-
        Ceph Object Storage Daemon latency is high. Please check if it doesn't stuck in weird state.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephHighOsdLatency.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `CephHighOsdLatency`:

## Meaning
The `CephHighOsdLatency` alert is triggered when the average apply latency of Ceph Object Storage Daemons (OSDs) exceeds 5 seconds over a 1-minute period. This indicates that the OSDs are experiencing high latency, which can lead to slower write performance and potentially impact the overall health of the Ceph cluster.

## Impact
High OSD latency can have several negative consequences:

* Slower write performance, leading to increased latency for applications writing data to the Ceph cluster
* Increased risk of data loss or corruption due to delayed writes
* Potential for OSDs to become stuck in a weird state, leading to further performance degradation and instability in the cluster

## Diagnosis
To diagnose the root cause of high OSD latency, follow these steps:

1. Check the Ceph cluster's overall health using the `ceph -s` command or a monitoring dashboard.
2. Investigate the OSD's performance metrics, such as `ceph_osd_perf_apply_latency_seconds`, to identify which OSDs are experiencing high latency.
3. Check the OSD's debug logs for any errors or warnings related to high latency.
4. Verify that the OSDs have sufficient resources (e.g., CPU, memory, disk space) to operate effectively.
5. Check for any network connectivity issues or congestion that may be contributing to high latency.

## Mitigation
To mitigate high OSD latency, follow these steps:

1. Check for any recent changes to the Ceph cluster's configuration or deployment that may be contributing to high latency.
2. Increase the OSD's resources (e.g., CPU, memory, disk space) if necessary to improve performance.
3. Implement load balancing or distribute writes across multiple OSDs to reduce the load on individual OSDs.
4. Adjust the Ceph cluster's configuration to optimize performance, such as adjusting the `osd_target_transaction_size` or `osd_client_message_size` settings.
5. Consider upgrading the Ceph cluster to a newer version, which may include performance improvements or bug fixes related to OSD latency.