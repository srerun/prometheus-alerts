---
title: ProxmoxBackupPoolDiskOutOfSpace
description: Troubleshooting for alert ProxmoxBackupPoolDiskOutOfSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupPoolDiskOutOfSpace

Disk is almost full (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupPoolDiskOutOfSpace" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupPoolDiskOutOfSpace
expr: pbs_used / pbs_size * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Proxmox datastore out of space (datastore {{ $labels.datastore }})
    description: |-
        Disk is almost full (< 10% left)
          VALUE = {{ $value }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuppooldiskoutofspace/

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule:

## Meaning
This alert is triggered when a Proxmox backup pool disk is almost out of space, with less than 10% of free space left. This can cause backups to fail, leading to potential data loss.

## Impact
The impact of this alert is not addressed, backups may fail, resulting in:

* Data loss due to failed backups
* Increased risk of data corruption
* Inability to restore data in case of a disaster

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Proxmox backup pool disk usage using the `pbs_used` and `pbs_size` metrics in Prometheus.
2. Verify that the disk is indeed running low on space.
3. Check the backup pool configuration to ensure it is set up correctly.

## Mitigation
To mitigate the issue, follow these steps:

1. Increase the size of the Proxmox backup pool disk.
2. Remove unnecessary files or backups to free up space on the disk.
3. Consider implementing a more efficient backup retention policy to minimize storage usage.
4. Verify that backups are completing successfully and that data is being properly stored.

Additional Resources:

* [Proxmox Backup Pool Documentation](https://pve.proxmox.com/wiki/Backup)
* [PBS Exporter Documentation](https://github.com/prometheus-community/pbs-exporter)