---
title: ProxmoxBackupRootDiskFillingUpInFourDays
description: Troubleshooting for alert ProxmoxBackupRootDiskFillingUpInFourDays
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupRootDiskFillingUpInFourDays

Based on recent sampling, the root disk {{ $labels.instance }} is expected to fill up withing four days. Currently {{ $value | humanizePercentage }} is available.

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupRootDiskFillingUpInFourDays" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupRootDiskFillingUpInFourDays
expr: pbs_host_disk_used / pbs_host_disk_total > 0.7 and pbs_host_disk_used > 0 and predict_linear(pbs_host_disk_available[6h], 4 * 24 * 3600) < 0
for: 30m
labels:
    severity: warning
annotations:
    summary: Proxmox root disk is filling up.
    description: Based on recent sampling, the root disk {{ $labels.instance }} is expected to fill up withing four days. Currently {{ $value | humanizePercentage }} is available.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuprootdiskfillingupinfourdays/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ProxmoxBackupRootDiskFillingUpInFourDays`:

## Meaning

This alert indicates that the root disk of a Proxmox node is filling up and is expected to reach full capacity within four days.

## Impact

If no action is taken, the root disk will become full, leading to:

* Unavailability of the Proxmox node
* Potential data loss or corruption
* Increased risk of node crashes or instability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the alert description for the specific instance and current disk usage percentage.
2. Log in to the Proxmox node and check the disk usage using the `df` command or a similar tool.
4. Identify the files or directories consuming the most disk space.
5. Check the backup and retention policies to determine if they need to be adjusted.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and remove any unnecessary files or data on the root disk.
2. Adjust the backup and retention policies to ensure that backups are not consuming excessive disk space.
3. Consider increasing the disk space available for the Proxmox node.
4. Monitor the disk usage and alert thresholds to ensure that the issue does not reoccur.

Remember to update the runbook and annotations as necessary to fit your specific use case.