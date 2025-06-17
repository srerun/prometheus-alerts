---
title: ProxmoxBackupPoolDiskFillingUpInFourDays
description: Troubleshooting for alert ProxmoxBackupPoolDiskFillingUpInFourDays
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupPoolDiskFillingUpInFourDays

Based on recent sampling, the datastore {{ $labels.datastore }} is expected to fill up withing four days. Currently {{ $value | humanizePercentage }} is available.

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupPoolDiskFillingUpInFourDays" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupPoolDiskFillingUpInFourDays
expr: pbs_used / pbs_size > 0.7 and pbs_used > 0 and predict_linear(pbs_available[6h], 4 * 24 * 3600) < 0
for: 30m
labels:
    severity: warning
annotations:
    summary: Proxmox datastore is filling up.
    description: Based on recent sampling, the datastore {{ $labels.datastore }} is expected to fill up withing four days. Currently {{ $value | humanizePercentage }} is available.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuppooldiskfillingupinfourdays/

```

{{% /comment %}}

</details>


## Meaning
This alert is triggered when the Proxmox backup pool disk is predicted to fill up within the next 4 days. The alert is warning severity and indicates that the current available disk space is decreasing rapidly.

## Impact
If no action is taken, the Proxmox datastore will run out of disk space, leading to:

* Backup failures
* Data loss
* Downtime of critical systems

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the current disk usage and available space using the `pbs_used` and `pbs_available` metrics.
2. Verify the predicted fill-up time using the `predict_linear` function.
3. Check the Proxmox backup pool disk configuration and usage patterns.
4. Identify any recent changes or unusual usage patterns that may be contributing to the disk filling up.

## Mitigation
To mitigate the issue, follow these steps:

1. Free up disk space by deleting unnecessary files or backups.
2. Adjust the backup retention policies to reduce the amount of data being stored.
3. Consider adding more disk space to the Proxmox backup pool.
4. Monitor the disk usage and adjust the backup pool configuration as needed to prevent future fill-up predictions.

Note: Refer to the provided runbook link for more information and detailed instructions on how to resolve this issue.