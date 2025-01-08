---
title: CassandraCompactionExecutorBlockedTasks
description: Troubleshooting for alert CassandraCompactionExecutorBlockedTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraCompactionExecutorBlockedTasks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Cassandra compaction executor tasks are blocked

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraCompactionExecutorBlockedTasks" %}}

<!-- Rule when generated

```yaml
alert: CassandraCompactionExecutorBlockedTasks
expr: cassandra_stats{name="org:apache:cassandra:metrics:threadpools:internal:compactionexecutor:currentlyblockedtasks:count"} > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra compaction executor blocked tasks (instance {{ $labels.instance }})
    description: |-
        Some Cassandra compaction executor tasks are blocked
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraCompactionExecutorBlockedTasks.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
