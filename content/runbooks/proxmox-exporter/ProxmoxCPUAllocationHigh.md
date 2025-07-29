---
title: ProxmoxCPUAllocationHigh
description: Troubleshooting for alert ProxmoxCPUAllocationHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxCPUAllocationHigh

It is recommended to keep more of your node's CPU unallocated for use by PVE and other server applications your Proxmox node runs

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxCPUAllocationHigh" %}}

{{% comment %}}

```yaml
alert: ProxmoxCPUAllocationHigh
expr: |
    100 * (proxmox_node_cpus_allocated / proxmox_node_cpus_total) > {{ .Values.prometheusRule.threshold_ProxmoxCPUAllocationHigh | default 90 }}
for: 5m
labels:
    severity: critical
annotations:
    summary: Proxmox node {{ printf "{{ $labels.node }}" }} has {{ printf "{{ $value }}" }}% of its CPU allocated to guests
    description: It is recommended to keep more of your node's CPU unallocated for use by PVE and other server applications your Proxmox node runs
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxcpuallocationhigh/

```

{{% /comment %}}

</details>


## Meaning

The ProxmoxCPUAllocationHigh alert is triggered when the percentage of allocated CPU resources on a Proxmox node exceeds a configured threshold (defaulting to 90%) for a duration of 5 minutes or more. This alert indicates that the node is heavily utilized and may not have sufficient CPU resources available for the Proxmox Virtual Environment (PVE) and other server applications running on the node.

## Impact

If left unaddressed, high CPU allocation can lead to:

* Performance degradation of virtual machines and server applications
* Increased latency and response times
* Potential for node crashes or instability
* Inability to allocate sufficient resources to new virtual machines or applications
* Negative impact on overall system reliability and availability

## Diagnosis

To diagnose the root cause of the high CPU allocation, follow these steps:

1. Log in to the Proxmox web interface and navigate to the node exhibiting high CPU allocation.
2. Check the node's resource utilization (CPU, memory, and disk) to identify any bottlenecks or resource contention.
3. Examine the list of running virtual machines and their resource allocations to identify potential culprits.
4. Verify that the node is not experiencing any hardware or software issues that could be contributing to the high CPU allocation.
5. Review system logs for errors or warnings related to resource allocation or node performance.

## Mitigation

To mitigate the high CPU allocation, consider the following steps:

1. **Rebalance virtual machine resources**: Re-allocate CPU resources from heavily utilized virtual machines to less utilized ones, ensuring a more balanced distribution of resources.
2. **Migrate virtual machines**: Migrate virtual machines to other nodes with available resources, if possible, to alleviate the load on the affected node.
3. **Right-size virtual machines**: Review virtual machine configurations to ensure they are correctly sized for their workloads, and adjust as necessary to reduce CPU allocation.
4. **Implement resource limits**: Set resource limits for virtual machines to prevent over-allocation and ensure sufficient resources are available for the node and other applications.
5. **Upgrade node resources**: Consider upgrading the node's hardware resources (e.g., adding more CPU cores or upgrading to faster processors) to increase available capacity.
6. **Investigate node optimization**: Investigate opportunities to optimize node performance, such as tuning PVE settings or adjusting system configurations.