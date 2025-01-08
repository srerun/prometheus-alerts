---
title: CassandraCommitlogPendingTasks
description: Troubleshooting for alert CassandraCommitlogPendingTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraCommitlogPendingTasks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Unexpected number of Cassandra commitlog pending tasks

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraCommitlogPendingTasks" %}}

<!-- Rule when generated

```yaml
alert: CassandraCommitlogPendingTasks
expr: cassandra_stats{name="org:apache:cassandra:metrics:commitlog:pendingtasks:value"} > 15
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra commitlog pending tasks (instance {{ $labels.instance }})
    description: |-
        Unexpected number of Cassandra commitlog pending tasks
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraCommitlogPendingTasks.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
