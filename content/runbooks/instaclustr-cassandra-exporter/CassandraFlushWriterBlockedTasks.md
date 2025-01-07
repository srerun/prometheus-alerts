---
title: CassandraFlushWriterBlockedTasks
description: Troubleshooting for alert CassandraFlushWriterBlockedTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraFlushWriterBlockedTasks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Cassandra flush writer tasks are blocked - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CassandraFlushWriterBlockedTasks
expr: cassandra_thread_pool_blocked_tasks{pool="MemtableFlushWriter"} > 15
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra flush writer blocked tasks (instance {{ $labels.instance }})
    description: |-
        Some Cassandra flush writer tasks are blocked - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CassandraFlushWriterBlockedTasks

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
