---
title: ProxmoxBackupHostDown
description: Troubleshooting for alert ProxmoxBackupHostDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupHostDown

Proxmox Backup Server {{ $labels.instance }} is not available.

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupHostDown" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupHostDown
expr: absent(pbs_available)
for: 2m
labels:
    severity: warning
annotations:
    summary: Proxmox Backup Server down
    description: Proxmox Backup Server {{ $labels.instance }} is not available.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuphostdown/

```

{{% /comment %}}

</details>


## Meaning

The ProxmoxBackupHostDown alert is triggered when the Proxmox Backup Server is not available. This means that the PBS (Proxmox Backup Server) exporter is not reporting the availability of the Proxmox Backup Server, which is crucial for ensuring the integrity of backups.

## Impact

The impact of this alert is moderate to high, as it directly affects the reliability of backups. If the Proxmox Backup Server is not available, backups may not be executed correctly, resulting in potential data loss. This can lead to:

* Incomplete or failed backups
* Data loss in case of a disaster
* Increased risk of data corruption
* Difficulty in recovering data in case of a recovery scenario

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox Backup Server logs for any errors or anomalies.
3. Verify that the PBS exporter is running and configured correctly.
2. Check the network connectivity between the Prometheus instance and the Proxmox Backup Server.
4. Check the Proxmox Backup Server status to ensure it is running and available.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Proxmox Backup Server service to ensure it is running correctly.
2. Check and update the PBS exporter configuration to ensure it is correctly reporting the availability of the Proxmox Backup Server.
3. Verify that the network connectivity between the Prometheus instance and the Proxmox Backup Server is stable.
4. If the issue persists, consider restarting the Proxmox Backup Server itself.

Remember to update the runbook according to your specific environment and needs.