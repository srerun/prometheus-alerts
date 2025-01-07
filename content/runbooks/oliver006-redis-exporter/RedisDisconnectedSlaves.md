---
title: RedisDisconnectedSlaves
description: Troubleshooting for alert RedisDisconnectedSlaves
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisDisconnectedSlaves

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Redis not replicating for all slaves. Consider reviewing the redis replication status.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RedisDisconnectedSlaves
expr: count without (instance, job) (redis_connected_slaves) - sum without (instance, job) (redis_connected_slaves) - 1 > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis disconnected slaves (instance {{ $labels.instance }})
    description: |-
        Redis not replicating for all slaves. Consider reviewing the redis replication status.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RedisDisconnectedSlaves

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
