---
title: HostHighCpuLoad
description: Troubleshooting for alert HostHighCpuLoad
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostHighCpuLoad

## Meaning
[//]: # "Short paragraph that explains what the alert means"
CPU load is > 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostHighCpuLoad" %}}

<!-- Rule when generated

```yaml
alert: HostHighCpuLoad
expr: (sum by (instance) (avg by (mode, instance) (rate(node_cpu_seconds_total{mode!="idle"}[2m]))) > 0.8) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 10m
labels:
    severity: warning
annotations:
    summary: Host high CPU load (instance {{ $labels.instance }})
    description: |-
        CPU load is > 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostHighCpuLoad.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
