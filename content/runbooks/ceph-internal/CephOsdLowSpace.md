---
title: CephOsdLowSpace
description: Troubleshooting for alert CephOsdLowSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephOsdLowSpace

Ceph Object Storage Daemon is going out of space. Please add more disks.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephOsdLowSpace" %}}

{{% comment %}}

```yaml
alert: CephOsdLowSpace
expr: ceph_osd_utilization > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph OSD low space (instance {{ $labels.instance }})
    description: |-
        Ceph Object Storage Daemon is going out of space. Please add more disks.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephOsdLowSpace.md

```

{{% /comment %}}

</details>


Here is a runbook for the CephOsdLowSpace alert:

## Meaning

The CephOsdLowSpace alert is triggered when the utilization of a Ceph Object Storage Daemon (OSD) exceeds 90% for more than 2 minutes. This indicates that the OSD is running low on available space, which can lead to reduced performance and increased risk of data loss.

## Impact

If left unaddressed, a Ceph OSD with low available space can cause:

* Reduced write performance and increased latency
* Increased risk of data loss or corruption
* Potential for the OSD to become unavailable, leading to a reduction in overall cluster capacity
* Increased risk of cascading failures in the cluster

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Ceph cluster's overall health using the `ceph -s` command
2. Identify the specific OSD that triggered the alert using the `instance` label in the alert notification
3. Check the OSD's utilization and available space using the `ceph osd df` command
4. Verify that the OSD is not experiencing any other issues, such as high latency or errors

## Mitigation

To mitigate the issue, follow these steps:

1. Add additional disks to the OSD to increase its available space
2. Consider rebalancing the data in the cluster to reduce the load on the affected OSD
3. Monitor the OSD's utilization and available space to ensure the issue is resolved
4. Consider adjusting the alert threshold or configuration to prevent false positives or unnecessary alerts