---
title: SqlServerDeadlock
description: Troubleshooting for alert SqlServerDeadlock
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SqlServerDeadlock

## Meaning
[//]: # "Short paragraph that explains what the alert means"
SQL Server is having some deadlock.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: SqlServerDeadlock
expr: increase(mssql_deadlocks[1m]) > 5
for: 0m
labels:
    severity: warning
annotations:
    summary: SQL Server deadlock (instance {{ $labels.instance }})
    description: |-
        SQL Server is having some deadlock.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/SqlServerDeadlock

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
