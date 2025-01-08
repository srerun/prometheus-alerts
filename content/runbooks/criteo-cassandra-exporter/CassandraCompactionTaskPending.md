---
title: CassandraCompactionTaskPending
description: Troubleshooting for alert CassandraCompactionTaskPending
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraCompactionTaskPending

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Many Cassandra compaction tasks are pending. You might need to increase I/O capacity by adding nodes to the cluster.

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraCompactionTaskPending" %}}

{{% comment %}}

```yaml
alert: CassandraCompactionTaskPending
expr: avg_over_time(cassandra_stats{name="org:apache:cassandra:metrics:compaction:pendingtasks:value"}[1m]) > 100
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra compaction task pending (instance {{ $labels.instance }})
    description: |-
        Many Cassandra compaction tasks are pending. You might need to increase I/O capacity by adding nodes to the cluster.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraCompactionTaskPending.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
