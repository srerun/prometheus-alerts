---
title: CephPgBackfillFull
description: Troubleshooting for alert CephPgBackfillFull
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephPgBackfillFull

Some Ceph placement groups are located on full Object Storage Daemon on cluster. Those PGs can be unavailable shortly. Please check OSDs, change weight or reconfigure CRUSH rules.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephPgBackfillFull" %}}

{{% comment %}}

```yaml
alert: CephPgBackfillFull
expr: ceph_pg_backfill_toofull > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph PG backfill full (instance {{ $labels.instance }})
    description: |-
        Some Ceph placement groups are located on full Object Storage Daemon on cluster. Those PGs can be unavailable shortly. Please check OSDs, change weight or reconfigure CRUSH rules.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephPgBackfillFull.md

```

{{% /comment %}}

</details>


Here is a runbook for the CephPgBackfillFull alert:

## Meaning

The CephPgBackfillFull alert is triggered when one or more Ceph placement groups (PGs) are located on a full Object Storage Daemon (OSD) on the cluster. This means that the OSD has reached its maximum capacity, and the PGs on that OSD are at risk of becoming unavailable shortly.

## Impact

The impact of this alert is that some PGs may become unavailable, leading to:

* Reduced performance and availability of the Ceph cluster
* Potential data loss or corruption
* Disruption to applications and services that rely on the Ceph cluster

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the OSD utilization: Verify that the OSD utilization is indeed high and identify the specific OSD(s) that are full.
2. Check the CRUSH map: Review the CRUSH (Controlled Replication Under Scalable Hashing) map to ensure that it is correctly configured and up-to-date.
3. Check the PG distribution: Verify that the PGs are evenly distributed across the OSDs and that no single OSD is overloaded.
4. Check for any OSD issues: Investigate if there are any issues with the OSDs, such as disk failures or network connectivity problems.

## Mitigation

To mitigate this alert, follow these steps:

1. Reduce OSD utilization: Free up space on the full OSD by deleting unnecessary data, migrating data to other OSDs, or adding more OSDs to the cluster.
2. Reconfigure CRUSH rules: Update the CRUSH map to ensure that it is correctly configured and balanced.
3. Rebalance PGs: Rebalance the PGs across the OSDs to ensure even distribution and prevent overload on any single OSD.
4. Monitor OSD utilization: Closely monitor OSD utilization to prevent similar issues in the future.

Note: The above steps are general guidelines and may vary depending on your specific Ceph cluster configuration and deployment.