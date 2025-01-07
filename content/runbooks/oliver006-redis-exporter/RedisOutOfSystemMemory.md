---
title: RedisOutOfSystemMemory
description: Troubleshooting for alert RedisOutOfSystemMemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisOutOfSystemMemory

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Redis is running out of system memory (> 90%)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RedisOutOfSystemMemory
expr: redis_memory_used_bytes / redis_total_system_memory_bytes * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Redis out of system memory (instance {{ $labels.instance }})
    description: |-
        Redis is running out of system memory (> 90%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RedisOutOfSystemMemory

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
