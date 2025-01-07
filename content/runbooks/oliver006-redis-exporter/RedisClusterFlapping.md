---
title: RedisClusterFlapping
description: Troubleshooting for alert RedisClusterFlapping
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisClusterFlapping

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Changes have been detected in Redis replica connection. This can occur when replica nodes lose connection to the master and reconnect (a.k.a flapping).

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RedisClusterFlapping
expr: changes(redis_connected_slaves[1m]) > 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Redis cluster flapping (instance {{ $labels.instance }})
    description: |-
        Changes have been detected in Redis replica connection. This can occur when replica nodes lose connection to the master and reconnect (a.k.a flapping).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RedisClusterFlapping

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
