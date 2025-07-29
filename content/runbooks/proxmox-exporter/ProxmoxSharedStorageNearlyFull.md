---
title: ProxmoxSharedStorageNearlyFull
description: Troubleshooting for alert ProxmoxSharedStorageNearlyFull
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxSharedStorageNearlyFull

Volume of type {{ printf "{{ $labels.type }}" }} is {{ printf "{{ $value }}" }}% full

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxSharedStorageNearlyFull" %}}

{{% comment %}}

```yaml
alert: ProxmoxSharedStorageNearlyFull
expr: |
    100 * (sum by (storage) (proxmox_node_storage_used_bytes{shared="true"}) / sum by (storage) (proxmox_node_storage_total_bytes)) > {{ .Values.prometheusRule.threshold_ProxmoxSharedStorageNearlyFull | default 90 }}
for: 5m
labels:
    severity: critical
annotations:
    summary: Proxmox shared storage volume {{ printf "{{ $labels.storage }}" }} nearly full
    description: Volume of type {{ printf "{{ $labels.type }}" }} is {{ printf "{{ $value }}" }}% full
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxsharedstoragenearlyfull/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ProxmoxSharedStorageNearlyFull`:

## Meaning

The `ProxmoxSharedStorageNearlyFull` alert is triggered when the shared storage volume on a Proxmox node reaches a certain threshold (default 90%) of its total capacity. This alert indicates that the storage volume is nearly full and requires immediate attention to prevent data loss or system instability.

## Impact

If left unaddressed, a full shared storage volume can lead to:

* Data loss or corruption
* System crashes or instability
* Inability to perform backups or snapshots
* Increased risk of data loss or corruption due to failed writes

## Diagnosis

To diagnose the issue, follow these steps:

1. Log in to the Proxmox node and check the storage volume usage.
2. Identify the type of storage volume that is nearly full (e.g., HDD, SSD, NVMe).
3. Verify that the storage volume is properly configured and mounted.
4. Check the storage volume for any signs of file system corruption or errors.
5. Review the system logs for any related error messages or warnings.

## Mitigation

To mitigate the issue, follow these steps:

1. **Immediate action**: Stop any non-essential services or applications that are writing data to the shared storage volume.
2. **Free up disk space**: Delete unnecessary files, logs, or data to free up disk space.
3. **Expand storage capacity**: Consider adding more storage capacity to the shared storage volume or migrating data to a larger storage volume.
4. **Configure disk quotas**: Implement disk quotas to prevent individual users or applications from consuming too much disk space.
5. **Schedule a maintenance window**: Schedule a maintenance window to perform a thorough cleaning and optimization of the shared storage volume.
6. **Monitor storage usage**: Set up monitoring tools to track storage usage and receive early warnings for potential storage capacity issues.

Remember to update the runbook URL in the alert rule to point to this runbook document.