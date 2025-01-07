---
title: ElasticsearchHealthyDataNodes
description: Troubleshooting for alert ElasticsearchHealthyDataNodes
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchHealthyDataNodes

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Missing data node in Elasticsearch cluster

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ElasticsearchHealthyDataNodes
expr: elasticsearch_cluster_health_number_of_data_nodes < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch Healthy Data Nodes (instance {{ $labels.instance }})
    description: |-
        Missing data node in Elasticsearch cluster
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ElasticsearchHealthyDataNodes

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
