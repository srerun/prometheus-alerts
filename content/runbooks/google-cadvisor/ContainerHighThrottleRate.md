---
title: ContainerHighThrottleRate
description: Troubleshooting for alert ContainerHighThrottleRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerHighThrottleRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Container is being throttled

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ContainerHighThrottleRate
expr: sum(increase(container_cpu_cfs_throttled_periods_total{container!=""}[5m])) by (container, pod, namespace) / sum(increase(container_cpu_cfs_periods_total[5m])) by (container, pod, namespace) > ( 25 / 100 )
for: 5m
labels:
    severity: warning
annotations:
    summary: Container high throttle rate (instance {{ $labels.instance }})
    description: |-
        Container is being throttled
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ContainerHighThrottleRate

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
