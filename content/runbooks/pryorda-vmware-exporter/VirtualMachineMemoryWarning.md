---
title: VirtualMachineMemoryWarning
description: Troubleshooting for alert VirtualMachineMemoryWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# VirtualMachineMemoryWarning

High memory usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%

<details>
  <summary>Alert Rule</summary>

{{% rule "vmware/pryorda-vmware-exporter.yml" "VirtualMachineMemoryWarning" %}}

{{% comment %}}

```yaml
alert: VirtualMachineMemoryWarning
expr: vmware_vm_mem_usage_average / 100 >= 80 and vmware_vm_mem_usage_average / 100 < 90
for: 5m
labels:
    severity: warning
annotations:
    summary: Virtual Machine Memory Warning (instance {{ $labels.instance }})
    description: |-
        High memory usage on {{ $labels.instance }}: {{ $value | printf "%.2f"}}%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pryorda-vmware-exporter/VirtualMachineMemoryWarning.md

```

{{% /comment %}}

</details>


Here is a runbook for the VirtualMachineMemoryWarning alert rule:

## Meaning

The VirtualMachineMemoryWarning alert is triggered when the average memory usage of a virtual machine (VM) exceeds 80% and is less than 90%. This indicates that the VM is experiencing high memory usage, which may lead to performance issues or even crashes if left unattended.

## Impact

If this alert is not addressed, the VM may experience:

* Slowed performance
* Increased swap usage
* Potential crashes or freezing
* Degraded user experience

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the VM's current memory usage and historical trends using the `vmware_vm_mem_usage_average` metric.
2. Verify that the VM's memory resources are sufficient for its workload.
3. Check for any memory-intensive processes or applications running on the VM.
4. Review the VM's configuration and ensure that memory allocation is set correctly.

## Mitigation

To mitigate this issue, follow these steps:

1. **Increase the VM's memory resources**: Allocate more memory to the VM to alleviate the high usage.
2. **Optimize resource utilization**: Identify and terminate any unnecessary or resource-intensive processes or applications on the VM.
3. **Monitor memory usage**: Closely monitor the VM's memory usage to ensure it does not exceed the warning threshold again.
4. **Consider upgrading the VM's hardware**: If the VM's workload consistently requires high memory usage, consider upgrading the underlying hardware to prevent future issues.

Remember to update the alert annotations and runbook to reflect any changes made to the mitigation steps.