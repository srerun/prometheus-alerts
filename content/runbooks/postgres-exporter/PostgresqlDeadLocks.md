---
title: PostgresqlDeadLocks
description: Troubleshooting for alert PostgresqlDeadLocks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlDeadLocks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
PostgreSQL has dead-locks

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PostgresqlDeadLocks
expr: increase(pg_stat_database_deadlocks{datname!~"template.*|postgres"}[1m]) > 5
for: 0m
labels:
    severity: warning
annotations:
    summary: Postgresql dead locks (instance {{ $labels.instance }})
    description: |-
        PostgreSQL has dead-locks
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PostgresqlDeadLocks

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
