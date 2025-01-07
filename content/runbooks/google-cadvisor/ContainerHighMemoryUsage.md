---
title: ContainerHighMemoryUsage
description: Troubleshooting for alert ContainerHighMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerHighMemoryUsage

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Container Memory usage is above 80%

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ContainerHighMemoryUsage
expr: (sum(container_memory_working_set_bytes{name!=""}) BY (instance, name) / sum(container_spec_memory_limit_bytes > 0) BY (instance, name) * 100) > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: Container High Memory usage (instance {{ $labels.instance }})
    description: |-
        Container Memory usage is above 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ContainerHighMemoryUsage

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
