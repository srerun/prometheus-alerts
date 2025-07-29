---
title: ProxmoxMemoryAllocationHigh
description: Troubleshooting for alert ProxmoxMemoryAllocationHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxMemoryAllocationHigh

It is recommended to keep more of your node's memory unallocated for use by PVE and other server applications your Proxmox node runs

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxMemoryAllocationHigh" %}}

{{% comment %}}

```yaml
alert: ProxmoxMemoryAllocationHigh
expr: |
    100 * (proxmox_node_memory_allocated_bytes / proxmox_node_memory_total_bytes) > {{ .Values.prometheusRule.threshold_ProxmoxMemoryAllocationHigh | default 90 }}
for: 5m
labels:
    severity: critical
annotations:
    summary: Proxmox node {{ printf "{{ $labels.node }}" }} has {{ printf "{{ $value }}" }}% of its memory allocated to guests
    description: It is recommended to keep more of your node's memory unallocated for use by PVE and other server applications your Proxmox node runs
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxmemoryallocationhigh/

```

{{% /comment %}}

</details>


Here is a runbook for the ProxmoxMemoryAllocationHigh alert:

## Meaning

The ProxmoxMemoryAllocationHigh alert is triggered when the percentage of memory allocated to guests on a Proxmox node exceeds a configurable threshold (default 90%) for 5 minutes or more. This indicates that a significant portion of the node's memory is being utilized by virtual machines, leaving limited resources for other system processes and applications.

## Impact

If left unaddressed, high memory allocation can lead to performance degradation, increased latency, and potential crashes on the Proxmox node. This may result in:

* Slow response times for virtual machines and applications
* Increased risk of node crashes and downtime
* Decreased overall system reliability and performance

## Diagnosis

To diagnose the issue, follow these steps:

1. Log in to the Proxmox node using SSH or the WebUI.
2. Check the current memory allocation using the `pvesh get /nodes/{node}/memory` command or the WebUI's "Nodes" > "Summary" page.
3. Verify that the allocated memory percentage is above the threshold.
4. Identify which virtual machines or applications are consuming the most memory using `pvesh get /nodes/{node}/qemu` or the WebUI's "Nodes" > "QEMU" page.
5. Review the node's configuration, including resource allocation and quotas, to determine if any adjustments are necessary.

## Mitigation

To mitigate the issue, take the following steps:

1. Identify and shut down or migrate virtual machines that are not critical or are consuming excessive memory.
2. Adjust resource allocation and quotas for virtual machines to ensure a more balanced distribution of resources.
3. Consider adding more physical memory to the node if possible.
4. Implement memory-efficient configurations and optimize virtual machine settings to reduce memory usage.
5. Monitor the node's memory allocation and overall performance to ensure the issue is resolved and does not recur.

Remember to adapt the mitigation steps to your specific environment and requirements. It is essential to strike a balance between allocating sufficient resources to virtual machines and maintaining a healthy amount of free memory for system processes and applications.