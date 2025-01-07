---
title: ElasticsearchClusterRed
description: Troubleshooting for alert ElasticsearchClusterRed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchClusterRed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Elastic Cluster Red status

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ElasticsearchClusterRed
expr: elasticsearch_cluster_health_status{color="red"} == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Elasticsearch Cluster Red (instance {{ $labels.instance }})
    description: |-
        Elastic Cluster Red status
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ElasticsearchClusterRed

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
