---
title: OutdatedSnapshots
description: Troubleshooting for alert OutdatedSnapshots
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OutdatedSnapshots

Outdated snapshots on {{ $labels.instance }}: {{ $value | printf "%.0f"}} days

<details>
  <summary>Alert Rule</summary>

{{% rule "vmware/pryorda-vmware-exporter.yml" "OutdatedSnapshots" %}}

{{% comment %}}

```yaml
alert: OutdatedSnapshots
expr: (time() - vmware_vm_snapshot_timestamp_seconds) / (60 * 60 * 24) >= 3
for: 5m
labels:
    severity: warning
annotations:
    summary: Outdated Snapshots (instance {{ $labels.instance }})
    description: |-
        Outdated snapshots on {{ $labels.instance }}: {{ $value | printf "%.0f"}} days
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pryorda-vmware-exporter/OutdatedSnapshots.md

```

{{% /comment %}}

</details>


Here is a runbook for the OutdatedSnapshots alert rule:

## Meaning

This alert is triggered when a VMware VM snapshot is outdated, meaning it has not been updated in 3 days or more. This can be a problem because outdated snapshots can cause issues with VM performance, backup and restore operations, and can also lead to storage capacity issues.

## Impact

The impact of this alert is moderate to high. Outdated snapshots can cause a range of problems, including:

* Performance degradation: Outdated snapshots can cause VMs to slow down or become unresponsive, leading to downtime and impacting business operations.
* Backup and restore issues: Outdated snapshots can make it difficult or impossible to restore VMs from backups, leading to data loss and business disruption.
* Storage capacity issues: Outdated snapshots can consume large amounts of storage space, leading to capacity issues and impacting the overall performance of the VMware environment.

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the VMWare vCenter Server or ESXi host to identify the affected VM and snapshot.
2. Verify that the snapshot is indeed outdated and not being updated regularly.
3. Check the VMware logs for any errors or issues related to snapshot creation or updating.
4. Check the storage capacity and availability to ensure that there are no issues with storage space or performance.

## Mitigation

To mitigate this alert, follow these steps:

1. Update the outdated snapshot to the latest version.
2. Verify that the snapshot is being updated regularly to prevent future issues.
3. Consider implementing a snapshot management policy to ensure that snapshots are regularly updated and deleted when no longer needed.
4. Review and adjust storage capacity and performance to ensure that there are no issues with storage space or performance.
5. Consider automating snapshot management using VMware APIs or third-party tools to simplify the process and reduce the risk of human error.