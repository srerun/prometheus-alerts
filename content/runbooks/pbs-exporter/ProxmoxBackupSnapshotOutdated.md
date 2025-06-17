---
title: ProxmoxBackupSnapshotOutdated
description: Troubleshooting for alert ProxmoxBackupSnapshotOutdated
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupSnapshotOutdated

Last snapshot of vm {{ $labels.vm_name }} is older than 2 days.

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupSnapshotOutdated" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupSnapshotOutdated
expr: (time() - avg_over_time(pbs_snapshot_vm_last_timestamp[5m])) / 3600 / 24 > 2
for: 2m
labels:
    severity: warning
annotations:
    summary: Last snapshot of vm is older than 2 days
    description: Last snapshot of vm {{ $labels.vm_name }} is older than 2 days.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackupsnapshotoutdated/

```

{{% /comment %}}

</details>


## Meaning

The ProxmoxBackupSnapshotOutdated` alert is triggered when the average time since the last snapshot of a virtual machine (VM) exceeds 2 days. This alert is critical as it may indicate a failure in the backup process, which can result in catastrophic consequences, including data loss, in the event of a system failure or data corruption.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus query `avg_over_time(pbs_snapshot_vm_last_timestamp[5m])` to verify the timestamp of the last snapshot.
2. Investigate the Proxmox backup logs to identify any errors or issues with the backup process.
3. Verify the VM's configuration and ensure that it is correctly set up for backup process.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify the Proxmox backup job is running successfully and scheduled correctly.
2. Take an immediate snapshot of the affected machine(s) to ensure data integrity.
3. Check and resolve any errors or issues found in the backup logs.
4. Verify the snapshot is completed successfully.
5. Consider adjusting the backup schedule to ensure more frequent snapshots are taken.
6. Consider implementing additional monitoring and alerting for backup process to detect potential issues earlier.

Remember to update the runbook with any additional information or steps specific to your environment.