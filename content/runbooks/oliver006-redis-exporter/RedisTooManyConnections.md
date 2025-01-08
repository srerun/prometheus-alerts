---
title: RedisTooManyConnections
description: Troubleshooting for alert RedisTooManyConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisTooManyConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Redis is running out of connections (> 90% used)

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisTooManyConnections" %}}

{{% comment %}}

```yaml
alert: RedisTooManyConnections
expr: redis_connected_clients / redis_config_maxclients * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Redis too many connections (instance {{ $labels.instance }})
    description: |-
        Redis is running out of connections (> 90% used)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisTooManyConnections.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
