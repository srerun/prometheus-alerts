---
title: ProxmoxDiskTargetLost
description: Troubleshooting for alert ProxmoxDiskTargetLost
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxDiskTargetLost

The disk {{ printf "{{ $labels.devpath }}" }} in node {{ printf "{{ $labels.node }}" }} is not showing up in metrics from Proxmox anymore

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxDiskTargetLost" %}}

{{% comment %}}

```yaml
alert: ProxmoxDiskTargetLost
expr: |
    absent_over_time(proxmox_node_disk_smart_status[1h])
for: 1m
labels:
    severity: critical
annotations:
    summary: Lost metrics for Proxmox disk {{ printf "{{ $labels.devpath }}" }}
    description: The disk {{ printf "{{ $labels.devpath }}" }} in node {{ printf "{{ $labels.node }}" }} is not showing up in metrics from Proxmox anymore
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxdisktargetlost/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "ProxmoxDiskTargetLost":

## Meaning

This runbook is triggered by the "ProxmoxDiskTargetLost" alert rule, which indicates that Prometheus has stopped receiving disk metrics from a Proxmox node. This means that the node is no longer sending disk data to Prometheus, which can lead to incomplete monitoring and potential issues with disk health.

## Impact

The impact of this alert is critical, as it affects the monitoring and health checks of the Proxmox disk. Without disk metrics, it is difficult to detect potential issues, such as disk failures or high usage, which can lead to:

* Data loss or corruption
* System crashes or downtime
* Performance degradation
* Incomplete monitoring and alerting

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox node logs for any errors or issues related to disk monitoring.
2. Verify that the Proxmox node is still sending metrics to Prometheus by checking the Prometheus UI or running a `promtool query` command.
3. Check the disk status on the Proxmox node using the `pvesh` command or the Proxmox Web UI.
4. Verify that the disk is properly configured and mounted on the Proxmox node.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Proxmox node to ensure that the disk monitoring is reinitialized.
2. Check the Proxmox node configuration to ensure that disk monitoring is enabled and properly configured.
3. Verify that the disk is properly connected and configured on the Proxmox node.
4. If the issue persists, consider resetting the Proxmox node's disk monitoring or seeking assistance from the Proxmox support team.

Remember to update the runbook link in the alert annotation to point to this runbook.