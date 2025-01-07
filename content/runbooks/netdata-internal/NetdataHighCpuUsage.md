---
title: NetdataHighCpuUsage
description: Troubleshooting for alert NetdataHighCpuUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataHighCpuUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Netdata high CPU usage (> 80%)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NetdataHighCpuUsage
expr: rate(netdata_cpu_cpu_percentage_average{dimension="idle"}[1m]) > 80
for: 5m
labels:
    severity: warning
annotations:
    summary: Netdata high cpu usage (instance {{ $labels.instance }})
    description: |-
        Netdata high CPU usage (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NetdataHighCpuUsage

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
