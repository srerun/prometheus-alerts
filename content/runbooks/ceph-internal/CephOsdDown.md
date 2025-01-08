---
title: CephOsdDown
description: Troubleshooting for alert CephOsdDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CephOsdDown

Ceph Object Storage Daemon Down

<details>
  <summary>Alert Rule</summary>

{{% rule "ceph/ceph-internal.yml" "CephOsdDown" %}}

{{% comment %}}

```yaml
alert: CephOsdDown
expr: ceph_osd_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Ceph OSD Down (instance {{ $labels.instance }})
    description: |-
        Ceph Object Storage Daemon Down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ceph-internal/CephOsdDown.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the CephOsdDown alert rule:

## Meaning

The CephOsdDown alert rule is triggered when a Ceph Object Storage Daemon (OSD) is down. Ceph OSDs are responsible for storing data in a Ceph cluster, and a down OSD can lead to data unavailability and reduced redundancy.

## Impact

The impact of a down Ceph OSD can be severe, as it can lead to:

* Data unavailability: If the OSD is down, the data stored on it may not be accessible, leading to service disruptions.
* Reduced redundancy: A down OSD can reduce the overall redundancy of the Ceph cluster, making it more vulnerable to data loss in case of further failures.
* Performance degradation: The remaining OSDs may need to handle additional workload, leading to performance degradation and slower write performance.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Ceph cluster status using `ceph -s` command.
2. Verify the OSD status using `ceph osd tree` command.
3. Check the OSD log files for any error messages.
4. Verify the OSD is not in a maintenance mode.
5. Check the network connectivity to the OSD node.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the OSD service on the affected node.
2. Check for any software or hardware issues on the OSD node and resolve them.
3. If the OSD is faulty, replace it with a new one and re-add it to the Ceph cluster.
4. Run `ceph osd repair` command to repair any damaged or inconsistent OSDs.
5. Verify the Ceph cluster status and OSD status after the mitigation steps to ensure the issue is resolved.

Remember to follow your organization's specific procedures and guidelines for troubleshooting and resolving Ceph OSD issues. This runbook is meant to provide general guidance and may need to be customized to fit your specific environment.