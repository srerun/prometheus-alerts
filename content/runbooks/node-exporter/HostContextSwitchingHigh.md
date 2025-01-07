---
title: HostContextSwitchingHigh
description: Troubleshooting for alert HostContextSwitchingHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostContextSwitchingHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Context switching is growing on the node (twice the daily average during the last 15m)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostContextSwitchingHigh
expr: '(rate(node_context_switches_total[15m])/count without(mode,cpu) (node_cpu_seconds_total{mode="idle"})) / (rate(node_context_switches_total[1d])/count without(mode,cpu) (node_cpu_seconds_total{mode="idle"})) > 2 '
for: 0m
labels:
    severity: warning
annotations:
    summary: Host context switching high (instance {{ $labels.instance }})
    description: |-
        Context switching is growing on the node (twice the daily average during the last 15m)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostContextSwitchingHigh

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
