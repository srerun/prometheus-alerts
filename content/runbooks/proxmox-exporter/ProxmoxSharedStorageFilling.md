---
title: ProxmoxSharedStorageFilling
description: Troubleshooting for alert ProxmoxSharedStorageFilling
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxSharedStorageFilling

Volume of type {{ printf "{{ $labels.type }}" }} is {{ printf "{{ $value }}" }}% full

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxSharedStorageFilling" %}}

{{% comment %}}

```yaml
alert: ProxmoxSharedStorageFilling
expr: |
    100 * (sum by (storage) (proxmox_node_storage_used_bytes{shared="true"}) / sum by (storage) (proxmox_node_storage_total_bytes)) > {{ .Values.prometheusRule.threshold_ProxmoxSharedStorageFilling | default 80 }}
for: 5m
labels:
    severity: warning
annotations:
    summary: Proxmox shared storage volume {{ printf "{{ $labels.storage }}" }} nearly full
    description: Volume of type {{ printf "{{ $labels.type }}" }} is {{ printf "{{ $value }}" }}% full
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxsharedstoragefilling/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "ProxmoxSharedStorageFilling":

## Meaning
This alert is triggered when the shared storage volume in a Proxmox cluster is nearly full. This means that the available storage space is running low, and action is required to free up space or add more storage capacity to prevent data loss or system instability.

## Impact
The impact of this alert is high, as a full shared storage volume can cause the following issues:

* Data loss or corruption
* System instability or crashes
* Unavailability of virtual machines or services
* Inability to perform backups or snapshots

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Proxmox web interface or CLI to confirm the shared storage volume is running low on space.
2. Identify which storage volume is affected and how much space is available.
3. Review disk usage reports to determine which VMs or services are consuming the most storage space.
4. Check for any recent changes or updates that may have caused the storage usage to increase.

## Mitigation
To mitigate the issue, follow these steps:

1. Immediately free up space on the shared storage volume by deleting unnecessary files, logs, or snapshots.
2. Identify and migrate or delete unnecessary VMs or services that are consuming excessive storage space.
3. Consider adding more storage capacity to the Proxmox cluster or upgrading the existing storage hardware.
4. Implement a monitoring and alerting system to track storage usage and prevent future occurrences of this issue.
5. Review and adjust the storage allocation and quotas for VMs and services to prevent over-usage.

Remember to update the threshold value in the Prometheus alert rule to a suitable value for your environment, and to test the alert and runbook regularly to ensure they are working correctly.