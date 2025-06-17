---
title: ProxmoxBackupSnapshotVerifyFailed
description: Troubleshooting for alert ProxmoxBackupSnapshotVerifyFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupSnapshotVerifyFailed

Last verified snapshot of vm {{ $labels.vm_name }} is older than 2 days.

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupSnapshotVerifyFailed" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupSnapshotVerifyFailed
expr: sum by (vm_id) (max_over_time(pbs_snapshot_vm_last_verify[2d]) and pbs_snapshot_vm_count > 1) == 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Last verified Proxmox snapshot older than 2 days
    description: Last verified snapshot of vm {{ $labels.vm_name }} is older than 2 days.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackupsnapshotverifyfailed/

```

{{% /comment %}}

</details>


## Meaning

The ProxmoxBackupSnapshotVerifyFailed alert is triggered when the last verified snapshot of a Proxmox machine (VM) is more than 2 days old. This alert indicates that the backup verification process for the VM has failed or not been performed recently.

## Impact

The impact of this alert is a potential loss of data in case of a VM failure or data corruption. If the last verified snapshot is older than 2 days, it means that the backup system has not been able to verify the integrity of the VM's data recently. This increases the risk of data loss in case of a VM failure or data corruption, the backup system may not be able to restore the data to a consistent state.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox Backup Server (PBS) logs for the affected VM to identify any errors or issues.
2. Verify that the PBS is configured correctly and the VM is properly registered.
3. Check the PBS dashboard to ensure that the VM is being backed up regularly.
4. Verify that the network connection between the PBS and the Proxmox cluster is stable and not experiencing any connectivity problems.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately trigger a manual backup verification for the VM using the PBS dashboard.
2. Verify that the PBS is configured correctly and the VM is properly registered.
3. Check the PBS logs for any errors or issues and address them as needed.
4. Consider increasing the frequency of backup verifications or implementing additional monitoring to detect potential issues earlier.

By following this runbook, you should be able to identify and resolve the root cause of the ProxmoxBackupSnapshotVerifyFailed alert, and ensure the integrity and availability of your Proxmox VMs.