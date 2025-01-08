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

CPU steal is > 10%. A noisy neighbor is killing VM performances or a spot instance may be out of credit.

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "HostCpuStealNoisyNeighbor" %}}

{{% comment %}}

```yaml
alert: HostCpuStealNoisyNeighbor
expr: rate(netdata_cpu_cpu_percentage_average{dimension="steal"}[1m]) > 10
for: 5m
labels:
    severity: warning
annotations:
    summary: Host CPU steal noisy neighbor (instance {{ $labels.instance }})
    description: |-
        CPU steal is > 10%. A noisy neighbor is killing VM performances or a spot instance may be out of credit.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/HostCpuStealNoisyNeighbor.md

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
