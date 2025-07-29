---
title: ProxmoxNodeDown
description: Troubleshooting for alert ProxmoxNodeDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxNodeDown

Check the alerting Proxmox host

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxNodeDown" %}}

{{% comment %}}

```yaml
alert: ProxmoxNodeDown
expr: |
    proxmox_node_up == 0
for: 1m
labels:
    severity: critical
annotations:
    summary: Proxmox node {{ printf "{{ $labels.node }}" }} is down
    description: Check the alerting Proxmox host
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxnodedown/

```

{{% /comment %}}

</details>


## Meaning

The ProxmoxNodeDown alert is triggered when a Proxmox node is down, indicating that the node is no longer responding or is unavailable. This alert is critical in nature, as it can impact the overall availability and reliability of the Proxmox cluster.

## Impact

The impact of a Proxmox node being down can be significant, as it can:

* Cause virtual machines to become unavailable or crash
* Disrupt critical services and applications relying on the node
* Lead to data loss or corruption
* Increase the risk of data breaches or security vulnerabilities
* Affect the overall performance and redundancy of the Proxmox cluster

## Diagnosis

To diagnose the ProxmoxNodeDown alert, follow these steps:

1. Check the Proxmox node's status in the Proxmox Web UI or using the `pvesh` command-line tool.
2. Verify that the node is not experiencing any hardware or software issues, such as disk failures, network connectivity problems, or kernel crashes.
3. Review the node's system logs for any errors or warnings that may indicate the cause of the issue.
4. Check the Proxmox cluster's configuration and ensure that it is correctly configured and that there are no misconfigurations that could be causing the node to be down.
5. Consult the Proxmox documentation and online resources for troubleshooting guides and known issues related to node downtime.

## Mitigation

To mitigate the ProxmoxNodeDown alert, follow these steps:

1. Immediately investigate the cause of the node downtime and take corrective action to resolve the issue.
2. If the node is experiencing a hardware failure, replace the faulty component and restore the node to a working state.
3. If the node is experiencing software issues, restart the node or perform a rolling update to the affected component.
4. If the node is experiencing configuration issues, correct the misconfiguration and restore the node to a working state.
5. Consider implementing redundant configurations, such as high availability (HA) clustering or load balancing, to minimize the impact of node downtime in the future.
6. Update the Proxmox cluster's configuration to prevent similar node downtime issues from occurring in the future.
7. Monitor the node's status and performance closely to ensure that it remains stable and available.