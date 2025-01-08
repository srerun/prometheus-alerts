---
title: RedisOutOfConfiguredMaxmemory
description: Troubleshooting for alert RedisOutOfConfiguredMaxmemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisOutOfConfiguredMaxmemory

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Redis is running out of configured maxmemory (> 90%)

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisOutOfConfiguredMaxmemory" %}}

{{% comment %}}

```yaml
alert: RedisOutOfConfiguredMaxmemory
expr: redis_memory_used_bytes / redis_memory_max_bytes * 100 > 90 and on(instance) redis_memory_max_bytes > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Redis out of configured maxmemory (instance {{ $labels.instance }})
    description: |-
        Redis is running out of configured maxmemory (> 90%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisOutOfConfiguredMaxmemory.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
