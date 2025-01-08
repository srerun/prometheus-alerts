---
title: RedisTooManyMasters
description: Troubleshooting for alert RedisTooManyMasters
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisTooManyMasters

Redis cluster has too many nodes marked as master.

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisTooManyMasters" %}}

{{% comment %}}

```yaml
alert: RedisTooManyMasters
expr: count(redis_instance_info{role="master"}) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis too many masters (instance {{ $labels.instance }})
    description: |-
        Redis cluster has too many nodes marked as master.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisTooManyMasters.md

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
