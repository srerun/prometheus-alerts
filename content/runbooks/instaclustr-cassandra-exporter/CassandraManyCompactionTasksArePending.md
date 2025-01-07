---
title: CassandraManyCompactionTasksArePending
description: Troubleshooting for alert CassandraManyCompactionTasksArePending
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraManyCompactionTasksArePending

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Many Cassandra compaction tasks are pending - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CassandraManyCompactionTasksArePending
expr: cassandra_table_estimated_pending_compactions > 100
for: 0m
labels:
    severity: warning
annotations:
    summary: Cassandra many compaction tasks are pending (instance {{ $labels.instance }})
    description: |-
        Many Cassandra compaction tasks are pending - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CassandraManyCompactionTasksArePending

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
