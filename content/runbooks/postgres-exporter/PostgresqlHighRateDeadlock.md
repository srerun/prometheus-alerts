---
title: PostgresqlHighRateDeadlock
description: Troubleshooting for alert PostgresqlHighRateDeadlock
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlHighRateDeadlock

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Postgres detected deadlocks

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PostgresqlHighRateDeadlock
expr: increase(postgresql_errors_total{type="deadlock_detected"}[1m]) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql high rate deadlock (instance {{ $labels.instance }})
    description: |-
        Postgres detected deadlocks
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PostgresqlHighRateDeadlock

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
