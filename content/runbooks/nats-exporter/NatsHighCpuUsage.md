---
title: NatsHighCpuUsage
description: Troubleshooting for alert NatsHighCpuUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighCpuUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
NATS server is using more than 80% CPU for the last 5 minutes

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NatsHighCpuUsage
expr: rate(gnatsd_varz_cpu[5m]) > 0.8
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high CPU usage (instance {{ $labels.instance }})
    description: |-
        NATS server is using more than 80% CPU for the last 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NatsHighCpuUsage

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
