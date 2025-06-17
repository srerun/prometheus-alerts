---
title: ProxmoxBackupRootDiskOutOfSpace
description: Troubleshooting for alert ProxmoxBackupRootDiskOutOfSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupRootDiskOutOfSpace

Disk is almost full (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupRootDiskOutOfSpace" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupRootDiskOutOfSpace
expr: pbs_host_disk_used / pbs_host_disk_total * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Proxmox root disk out of space (instance {{ $labels.instance }})
    description: |-
        Disk is almost full (< 10% left)
          VALUE = {{ $value }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuprootdiskoutofspace/

```

{{% /comment %}}

</details>


Here is a sample runbook for the ProxmoxBackupRootDiskOutOfSpace alert:

## Meaning

The ProxmoxBackupRootDiskOutOfSpace alert is triggered when the used disk space on the root disk of a Proxmox host exceeds 90% capacity. This alert indicates that the disk is almost full, with less than 10% of available space remaining.

## Impact

A full root disk can cause stability issues, and may lead to system crashes or corruption. If left unattended, this issue may result in:

* System downtime
* Loss of data
* Increased troubleshooting and repair time

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the disk usage of the Proxmox host using the `pbs_host_disk_used` metric.
2. Verify the total disk capacity of the root disk using the `pbs_host_disk_total` metric.
3. Calculate the percentage of used disk space using the formula `(pbs_host_disk_used / pbs_host_disk_total) * 100`.
4. Check the instance label `{{ $labels.instance }}` to identify the affected Proxmox host.

## Mitigation

To mitigate the issue, perform the following steps:

1. ** Immediately free up disk space **: Remove unnecessary files, logs, and snapshots to free up space.
2. Configure disk usage monitoring to prevent similar issues in the future.
3. Consider expanding the root disk or adding additional storage capacity to the Proxmox host.
4. Document the resolution in the incident report.

Note: This runbook provides general guidance and may need to be tailored to your specific environment and requirements.