---
title: ElasticsearchInitializingShardsTooLong
description: Troubleshooting for alert ElasticsearchInitializingShardsTooLong
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchInitializingShardsTooLong

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Elasticsearch has been initializing shards for 15 min

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ElasticsearchInitializingShardsTooLong
expr: elasticsearch_cluster_health_initializing_shards > 0
for: 15m
labels:
    severity: warning
annotations:
    summary: Elasticsearch initializing shards too long (instance {{ $labels.instance }})
    description: |-
        Elasticsearch has been initializing shards for 15 min
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ElasticsearchInitializingShardsTooLong

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
