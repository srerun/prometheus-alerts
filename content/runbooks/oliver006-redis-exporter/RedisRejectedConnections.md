---
title: RedisRejectedConnections
description: Troubleshooting for alert RedisRejectedConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisRejectedConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some connections to Redis has been rejected

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RedisRejectedConnections
expr: increase(redis_rejected_connections_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis rejected connections (instance {{ $labels.instance }})
    description: |-
        Some connections to Redis has been rejected
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RedisRejectedConnections

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
