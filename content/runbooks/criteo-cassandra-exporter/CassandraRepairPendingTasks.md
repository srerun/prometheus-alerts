---
title: CassandraRepairPendingTasks
description: Troubleshooting for alert CassandraRepairPendingTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraRepairPendingTasks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Cassandra repair tasks are pending

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraRepairPendingTasks" %}}

<!-- Rule when generated

```yaml
alert: CassandraRepairPendingTasks
expr: cassandra_stats{name="org:apache:cassandra:metrics:threadpools:internal:antientropystage:pendingtasks:value"} > 2
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra repair pending tasks (instance {{ $labels.instance }})
    description: |-
        Some Cassandra repair tasks are pending
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraRepairPendingTasks.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
