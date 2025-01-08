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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
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


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
