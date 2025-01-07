---
title: RedisReplicationBroken
description: Troubleshooting for alert RedisReplicationBroken
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisReplicationBroken

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Redis instance lost a slave

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RedisReplicationBroken
expr: delta(redis_connected_slaves[1m]) < 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis replication broken (instance {{ $labels.instance }})
    description: |-
        Redis instance lost a slave
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RedisReplicationBroken

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
