---
title: RedisDown
description: Troubleshooting for alert RedisDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Redis instance is down

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisDown" %}}

<!-- Rule when generated

```yaml
alert: RedisDown
expr: redis_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis down (instance {{ $labels.instance }})
    description: |-
        Redis instance is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisDown.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
