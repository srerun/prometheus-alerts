---
title: ContainerLowCpuUtilization
description: Troubleshooting for alert ContainerLowCpuUtilization
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerLowCpuUtilization

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Container CPU utilization is under 20% for 1 week. Consider reducing the allocated CPU.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ContainerLowCpuUtilization
expr: (sum(rate(container_cpu_usage_seconds_total{container!=""}[5m])) by (pod, container) / sum(container_spec_cpu_quota{container!=""}/container_spec_cpu_period{container!=""}) by (pod, container) * 100) < 20
for: 7d
labels:
    severity: info
annotations:
    summary: Container Low CPU utilization (instance {{ $labels.instance }})
    description: |-
        Container CPU utilization is under 20% for 1 week. Consider reducing the allocated CPU.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ContainerLowCpuUtilization

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
