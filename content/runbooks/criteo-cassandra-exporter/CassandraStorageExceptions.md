---
title: CassandraStorageExceptions
description: Troubleshooting for alert CassandraStorageExceptions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraStorageExceptions

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Something is going wrong with cassandra storage

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraStorageExceptions" %}}

<!-- Rule when generated

```yaml
alert: CassandraStorageExceptions
expr: changes(cassandra_stats{name="org:apache:cassandra:metrics:storage:exceptions:count"}[1m]) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra storage exceptions (instance {{ $labels.instance }})
    description: |-
        Something is going wrong with cassandra storage
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraStorageExceptions.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
