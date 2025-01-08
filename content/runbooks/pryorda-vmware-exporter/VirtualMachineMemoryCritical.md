---
title: VirtualMachineMemoryCritical
description: Troubleshooting for alert VirtualMachineMemoryCritical
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# VirtualMachineMemoryCritical

High memory usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "vmware/pryorda-vmware-exporter.yml" "VirtualMachineMemoryCritical" %}}

{{% comment %}}

```yaml
alert: VirtualMachineMemoryCritical
expr: vmware_vm_mem_usage_average / 100 >= 90
for: 1m
labels:
    severity: critical
annotations:
    summary: Virtual Machine Memory Critical (instance {{ $labels.instance }})
    description: |-
        High memory usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pryorda-vmware-exporter/VirtualMachineMemoryCritical.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "VirtualMachineMemoryCritical":

## Meaning

The VirtualMachineMemoryCritical alert is triggered when a virtual machine's memory usage averages 90% or higher over a 1-minute period. This alert indicates a critical memory usage situation that requires immediate attention to prevent potential crashes or performance degradation.

## Impact

If left unaddressed, high memory usage can lead to:

* Virtual machine crashes or freezes
* Performance degradation and slow response times
* Increased risk of data loss or corruption
* Decreased system availability and uptime

## Diagnosis

To diagnose the root cause of high memory usage, follow these steps:

1. Check the virtual machine's current memory usage and trends using the Prometheus console or other monitoring tools.
2. Review the virtual machine's configuration, including the allocated memory, CPU, and disk resources.
3. Investigate recent changes or updates to the virtual machine, guest operating system, or applications that may be contributing to increased memory usage.
4. Check for any memory leaks or misconfigured applications that may be causing excessive memory consumption.
5. Verify that the virtual machine is not experiencing any resource contention or bottlenecks that could be contributing to high memory usage.

## Mitigation

To mitigate high memory usage and resolve the VirtualMachineMemoryCritical alert, follow these steps:

1. Increase the virtual machine's allocated memory, if possible, to reduce the usage percentage.
2. Identify and terminate any unnecessary or resource-intensive applications or processes consuming excessive memory.
3. Optimize memory usage by implementing memory-efficient configurations, such as adjusting JVM heap sizes or tweaking application settings.
4. Consider migrating the virtual machine to a more powerful host or cluster with additional resources.
5. Implement monitoring and alerting to detect potential memory usage issues earlier, and schedule regular maintenance and optimization tasks to prevent recurrence.

Remember to update the runbook with the specific steps and procedures relevant to your environment and infrastructure.