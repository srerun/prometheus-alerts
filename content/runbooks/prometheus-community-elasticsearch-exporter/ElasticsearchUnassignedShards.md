---
title: ElasticsearchUnassignedShards
description: Troubleshooting for alert ElasticsearchUnassignedShards
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchUnassignedShards

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Elasticsearch has unassigned shards

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ElasticsearchUnassignedShards
expr: elasticsearch_cluster_health_unassigned_shards > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch unassigned shards (instance {{ $labels.instance }})
    description: |-
        Elasticsearch has unassigned shards
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ElasticsearchUnassignedShards

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
