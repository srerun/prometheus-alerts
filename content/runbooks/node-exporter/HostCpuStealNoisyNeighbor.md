---
title: HostCpuStealNoisyNeighbor
description: Troubleshooting for alert HostCpuStealNoisyNeighbor
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostCpuStealNoisyNeighbor

## Meaning
[//]: # "Short paragraph that explains what the alert means"
CPU steal is > 10%. A noisy neighbor is killing VM performances or a spot instance may be out of credit.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostCpuStealNoisyNeighbor" %}}

<!-- Rule when generated

```yaml
alert: HostCpuStealNoisyNeighbor
expr: (avg by(instance) (rate(node_cpu_seconds_total{mode="steal"}[5m])) * 100 > 10) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host CPU steal noisy neighbor (instance {{ $labels.instance }})
    description: |-
        CPU steal is > 10%. A noisy neighbor is killing VM performances or a spot instance may be out of credit.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostCpuStealNoisyNeighbor.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
