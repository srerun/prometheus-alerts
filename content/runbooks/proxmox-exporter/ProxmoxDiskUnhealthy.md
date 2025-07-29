---
title: ProxmoxDiskUnhealthy
description: Troubleshooting for alert ProxmoxDiskUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxDiskUnhealthy

The disk {{ printf "{{ $labels.devpath }}" }} in node {{ printf "{{ $labels.node }}" }} is reporting unhealthy in SMART tests

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxDiskUnhealthy" %}}

{{% comment %}}

```yaml
alert: ProxmoxDiskUnhealthy
expr: |
    proxmox_node_disk_smart_status == 0
for: 1m
labels:
    severity: critical
annotations:
    summary: Proxmox disk {{ printf "{{ $labels.devpath }}" }} is unhealthy
    description: The disk {{ printf "{{ $labels.devpath }}" }} in node {{ printf "{{ $labels.node }}" }} is reporting unhealthy in SMART tests
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxdiskunhealthy/

```

{{% /comment %}}

</details>


## Meaning

The ProxmoxDiskUnhealthy alert is triggered when a disk in a Proxmox node reports an unhealthy status in its SMART (Self-Monitoring, Analysis and Reporting Technology) tests. This alert indicates a potential issue with the disk's reliability and may lead to data loss or system crashes if left unattended.

## Impact

* Data loss or corruption
* System crashes or instability
* Downtime for the affected node
* Potential impact on hosted virtual machines or containers

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox Web Interface or the `pvecli` command-line tool to verify the disk's status and identify the specific error.
2. Review the disk's SMART attributes to determine the cause of the unhealthy status.
3. Check the system logs for any errors or warnings related to the disk.
4. Verify that the disk is properly seated and connected to the node.

## Mitigation

To mitigate the issue, follow these steps:

1. Replace the unhealthy disk with a functioning one, if possible.
2. Run a thorough disk check and repair any errors found.
3. Update the disk's firmware to the latest version.
4. Consider migrating critical data to a redundant storage system or backup.
5. Implement additional monitoring and alerting for disk health to detect potential issues earlier.

Note: The specific mitigation steps may vary depending on the specific error and the node's configuration. It is recommended to consult the Proxmox documentation and the disk manufacturer's guidelines for further guidance.