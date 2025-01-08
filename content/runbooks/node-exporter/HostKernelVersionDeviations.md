---
title: HostKernelVersionDeviations
description: Troubleshooting for alert HostKernelVersionDeviations
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostKernelVersionDeviations

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Different kernel versions are running

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostKernelVersionDeviations" %}}

{{% comment %}}

```yaml
alert: HostKernelVersionDeviations
expr: (count(sum(label_replace(node_uname_info, "kernel", "$1", "release", "([0-9]+.[0-9]+.[0-9]+).*")) by (kernel)) > 1) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 6h
labels:
    severity: warning
annotations:
    summary: Host kernel version deviations (instance {{ $labels.instance }})
    description: |-
        Different kernel versions are running
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostKernelVersionDeviations.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
