---
title: SidekiqSchedulingLatencyTooHigh
description: Troubleshooting for alert SidekiqSchedulingLatencyTooHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SidekiqSchedulingLatencyTooHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Sidekiq jobs are taking more than 1min to be picked up. Users may be seeing delays in background processing.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SidekiqSchedulingLatencyTooHigh
expr: max(sidekiq_queue_latency) > 60
for: 0m
labels:
    severity: critical
annotations:
    summary: Sidekiq scheduling latency too high (instance {{ $labels.instance }})
    description: |-
        Sidekiq jobs are taking more than 1min to be picked up. Users may be seeing delays in background processing.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/SidekiqSchedulingLatencyTooHigh

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
