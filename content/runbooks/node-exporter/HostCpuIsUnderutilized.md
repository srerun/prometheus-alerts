---
title: HostCpuIsUnderutilized
description: Troubleshooting for alert HostCpuIsUnderutilized
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostCpuIsUnderutilized

CPU load is < 20% for 1 week. Consider reducing the number of CPUs.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostCpuIsUnderutilized" %}}

{{% comment %}}

```yaml
alert: HostCpuIsUnderutilized
expr: (100 - (rate(node_cpu_seconds_total{mode="idle"}[30m]) * 100) < 20) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 1w
labels:
    severity: info
annotations:
    summary: Host CPU is underutilized (instance {{ $labels.instance }})
    description: |-
        CPU load is < 20% for 1 week. Consider reducing the number of CPUs.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostCpuIsUnderutilized.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
