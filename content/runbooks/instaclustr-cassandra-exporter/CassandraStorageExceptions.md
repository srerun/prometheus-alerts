---
title: CassandraStorageExceptions
description: Troubleshooting for alert CassandraStorageExceptions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraStorageExceptions

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Something is going wrong with cassandra storage - {{ $labels.cassandra_cluster }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CassandraStorageExceptions
expr: changes(cassandra_storage_exceptions_total[1m]) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra storage exceptions (instance {{ $labels.instance }})
    description: |-
        Something is going wrong with cassandra storage - {{ $labels.cassandra_cluster }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CassandraStorageExceptions

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
