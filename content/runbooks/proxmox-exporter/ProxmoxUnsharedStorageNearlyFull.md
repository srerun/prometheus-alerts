---
title: ProxmoxUnsharedStorageNearlyFull
description: Troubleshooting for alert ProxmoxUnsharedStorageNearlyFull
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxUnsharedStorageNearlyFull

Volume of type {{ printf "{{ $labels.type }}" }} is {{ printf "{{ $value }}" }}% full

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxUnsharedStorageNearlyFull" %}}

{{% comment %}}

```yaml
alert: ProxmoxUnsharedStorageNearlyFull
expr: |
    100 * (sum by (storage,node) (proxmox_node_storage_used_bytes{shared="false"}) / sum by (storage,node) (proxmox_node_storage_total_bytes)) > {{ .Values.prometheusRule.threshold_ProxmoxUnsharedStorageNearlyFull | default 90 }}
for: 5m
labels:
    severity: critical
annotations:
    summary: Proxmox storage volume {{ printf "{{ $labels.node }}" }}/{{ printf "{{ $labels.storage }}" }} nearly full
    description: Volume of type {{ printf "{{ $labels.type }}" }} is {{ printf "{{ $value }}" }}% full
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxunsharedstoragenearlyfull/

```

{{% /comment %}}

</details>


## Meaning

The `ProxmoxUnsharedStorageNearlyFull` alert is triggered when the usage of unshared storage on a Proxmox node exceeds a certain threshold (default 90%). This alert indicates that the available storage capacity is running low, which can lead to performance issues, data loss, or even complete system failure.

## Impact

The impact of this alert is high, as it can cause:

* Performance degradation: Low storage capacity can lead to slow disk I/O, causing sluggish system performance and affecting VMs and applications running on the node.
* Data loss: If the storage becomes completely full, new data cannot be written, and existing data may be lost or corrupted.
* System downtime: In extreme cases, a full storage volume can cause the entire system to become unavailable, leading to downtime and revenue loss.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected node and storage volume using the `{{ $labels.node }}` and `{{ $labels.storage }}` labels.
2. Check the storage usage graph for the affected node and volume to identify the rate of growth and potential causes.
3. Verify that there are no ongoing backup or data transfer processes that could be contributing to the high storage usage.
4. Check the Proxmox node's disk usage using the `proxmox_node_disk.Usage` metric to identify any other potential storage issues.

## Mitigation

To mitigate the issue, follow these steps:

1. **Free up storage space**: Immediately remove any unnecessary files, logs, or data from the affected storage volume to free up space.
2. **Identify and address the root cause**: Analyze the storage usage graph and system logs to identify the root cause of the high storage usage. This could be due to a misconfigured backup job, a data-intensive application, or a storage-hungry VM.
3. **Expand storage capacity**: Consider adding more storage capacity to the affected node or migrating VMs and data to a node with more available storage.
4. **Implement monitoring and alerting**: Set up additional monitoring and alerting rules to detect early signs of storage capacity issues, allowing for more proactive remediation.

Remember to update the `threshold_ProxmoxUnsharedStorageNearlyFull` value in your Prometheus configuration file to adjust the alert threshold according to your specific storage requirements.