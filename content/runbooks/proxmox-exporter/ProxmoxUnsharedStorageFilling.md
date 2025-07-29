---
title: ProxmoxUnsharedStorageFilling
description: Troubleshooting for alert ProxmoxUnsharedStorageFilling
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxUnsharedStorageFilling

Volume of type {{ printf "{{ $labels.type }}" }} is {{ printf "{{ $value }}" }}% full

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxUnsharedStorageFilling" %}}

{{% comment %}}

```yaml
alert: ProxmoxUnsharedStorageFilling
expr: |
    100 * (sum by (storage,node) (proxmox_node_storage_used_bytes{shared="false"}) / sum by (storage,node) (proxmox_node_storage_total_bytes)) > {{ .Values.prometheusRule.threshold_ProxmoxUnsharedStorageFilling | default 80 }}
for: 5m
labels:
    severity: warning
annotations:
    summary: Proxmox storage volume {{ printf "{{ $labels.node }}" }}/{{ printf "{{ $labels.storage }}" }} nearly full
    description: Volume of type {{ printf "{{ $labels.type }}" }} is {{ printf "{{ $value }}" }}% full
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxunsharedstoragefilling/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "ProxmoxUnsharedStorageFilling":

## Meaning

The ProxmoxUnsharedStorageFilling alert is triggered when the used storage capacity of a non-shared storage volume on a Proxmox node exceeds a certain threshold (default 80%). This alert is critical as it can lead to data loss, service disruptions, and system crashes.

## Impact

The impact of this alert is high, as it can:

* Cause data loss or corruption due to lack of available storage space
* Disrupt services running on the Proxmox node
* Lead to system crashes or instability
* Require immediate attention to prevent further damage

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected Proxmox node and storage volume using the alert labels (`node` and `storage`).
2. Check the current storage usage and available capacity on the affected storage volume.
3. Verify that the storage volume is not shared (i.e., `shared="false"`).
4. Investigate the cause of the rapid storage growth (e.g., data influx, software updates, etc.).
5. Check the Proxmox node's system logs for any errors or warnings related to storage or disk usage.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the root cause of the storage growth and address it (e.g., stop data influx, roll back software updates, etc.).
2. Free up storage space on the affected volume by:
	* Deleting unnecessary files or data
	* Moving data to a different storage volume
	* Expanding the storage volume (if possible)
3. Monitor the storage usage closely to ensure the issue does not recur.
4. Consider adjusting the alert threshold or implementing additional monitoring and alerting for storage usage.
5. If the issue persists, consider involving Proxmox support or an experienced system administrator for further assistance.