---
title: PostgresqlHighRollbackRate
description: Troubleshooting for alert PostgresqlHighRollbackRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlHighRollbackRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ratio of transactions being aborted compared to committed is > 2 %

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PostgresqlHighRollbackRate
expr: sum by (namespace,datname) ((rate(pg_stat_database_xact_rollback{datname!~"template.*|postgres",datid!="0"}[3m])) / ((rate(pg_stat_database_xact_rollback{datname!~"template.*|postgres",datid!="0"}[3m])) + (rate(pg_stat_database_xact_commit{datname!~"template.*|postgres",datid!="0"}[3m])))) > 0.02
for: 0m
labels:
    severity: warning
annotations:
    summary: Postgresql high rollback rate (instance {{ $labels.instance }})
    description: |-
        Ratio of transactions being aborted compared to committed is > 2 %
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/PostgresqlHighRollbackRate

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
