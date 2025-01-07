---
title: CassandraCompactionExecutorBlockedTasks
description: Troubleshooting for alert CassandraCompactionExecutorBlockedTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraCompactionExecutorBlockedTasks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Cassandra compaction executor tasks are blocked - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CassandraCompactionExecutorBlockedTasks
expr: cassandra_thread_pool_blocked_tasks{pool="CompactionExecutor"} > 15
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra compaction executor blocked tasks (instance {{ $labels.instance }})
    description: |-
        Some Cassandra compaction executor tasks are blocked - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CassandraCompactionExecutorBlockedTasks

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
