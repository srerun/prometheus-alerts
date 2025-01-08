---
title: RedisNotEnoughConnections
description: Troubleshooting for alert RedisNotEnoughConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisNotEnoughConnections

Redis instance should have more connections (> 5)

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisNotEnoughConnections" %}}

{{% comment %}}

```yaml
alert: RedisNotEnoughConnections
expr: redis_connected_clients < 5
for: 2m
labels:
    severity: warning
annotations:
    summary: Redis not enough connections (instance {{ $labels.instance }})
    description: |-
        Redis instance should have more connections (> 5)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisNotEnoughConnections.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
