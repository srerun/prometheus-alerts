---
title: CephOsdReweighted
description: Troubleshooting for alert CephOsdReweighted
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephOsdReweighted

Ceph Object Storage Daemon takes too much time to resize.

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephOsdReweighted" %}}

{{% comment %}}

```yaml
alert: CephOsdReweighted
expr: ceph_osd_weight < 1
for: 2m
labels:
    severity: warning
annotations:
    summary: Ceph OSD reweighted (instance {{ $labels.instance }})
    description: |-
        Ceph Object Storage Daemon takes too much time to resize.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephOsdReweighted.md

```

{{% /comment %}}

</details>


Here is the runbook for the CephOsdReweighted alert rule:

## Meaning

The CephOsdReweighted alert rule triggers when a Ceph Object Storage Daemon (OSD) has its weight reduced to less than 1. This means that the OSD is taking too much time to resize, which can lead to performance issues and impact the overall health of the Ceph cluster.

## Impact

The impact of this alert is that the Ceph cluster may experience reduced performance, slower data replication, and increased latency. This can lead to:

* Slower data access times for clients
* Increased risk of data loss or corruption
* Decreased overall cluster reliability

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the Ceph cluster status using `ceph -s` or `ceph status`
2. Verify the OSD's weight using `ceph osd tree` or `ceph osd dump`
3. Check the OSD's logs for any error messages related to resizing or slow performance
4. Verify that the OSD is not overloaded or experiencing high latency
5. Check if there are any ongoing maintenance tasks or upgrades that may be causing the issue

## Mitigation

To mitigate the issue, follow these steps:

1. Check if the OSD can be restarted or remapped to improve performance
2. Verify that the OSD's disk is not full or experiencing high latency
3. Check if there are any configuration issues or misconfigurations that may be causing the slow performance
4. Consider increasing the OSD's weight or adjusting the crush map to improve performance
5. If the issue persists, consider consulting the Ceph documentation or seeking assistance from a Ceph expert or support team.

Remember to always follow proper procedures and testing before making any changes to the Ceph cluster.