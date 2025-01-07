---
title: ElasticsearchPendingTasks
description: Troubleshooting for alert ElasticsearchPendingTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ElasticsearchPendingTasks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Elasticsearch has pending tasks. Cluster works slowly.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ElasticsearchPendingTasks
expr: elasticsearch_cluster_health_number_of_pending_tasks > 0
for: 15m
labels:
    severity: warning
annotations:
    summary: Elasticsearch pending tasks (instance {{ $labels.instance }})
    description: |-
        Elasticsearch has pending tasks. Cluster works slowly.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ElasticsearchPendingTasks

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
