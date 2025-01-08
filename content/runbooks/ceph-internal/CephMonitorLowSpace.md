---
title: CephMonitorLowSpace
description: Troubleshooting for alert CephMonitorLowSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephMonitorLowSpace

Ceph monitor storage is low.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephMonitorLowSpace" %}}

{{% comment %}}

```yaml
alert: CephMonitorLowSpace
expr: ceph_monitor_avail_percent < 10
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph monitor low space (instance {{ $labels.instance }})
    description: |-
        Ceph monitor storage is low.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephMonitorLowSpace.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The CephMonitorLowSpace alert is triggered when the available storage space on a Ceph monitor node falls below 10%. This alert indicates that the storage capacity of the monitor node is running low, which can lead to issues with the overall health and stability of the Ceph cluster.

## Impact

If left unaddressed, this issue can have significant impacts on the Ceph cluster, including:

* Reduced performance and availability of the cluster
* Increased risk of data loss or corruption
* Inability to store new data or metadata, leading to errors and failures
* Potential for cascading failures throughout the cluster

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. Verify the alert by checking the current available storage space on the monitor node using the `ceph df` command.
2. Check the monitoring dashboard for any signs of disk usage trends or anomalies.
3. Review the system logs for any errors or warnings related to storage or disk usage.
4. Investigate any recent changes to the cluster, such as increased data ingestion or node additions/removals.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately free up disk space on the monitor node by deleting unnecessary files or data.
2. Identify and address any underlying issues causing the storage usage to increase, such as data growth or inefficient storage practices.
3. Consider adding additional storage capacity to the monitor node or rebalancing data across the cluster.
4. Implement long-term monitoring and alerting to prevent similar issues from occurring in the future.

Remember to carefully evaluate and address the root cause of the issue to prevent recurrence and ensure the long-term health of the Ceph cluster.