---
title: PostgresqlTableNotAutoAnalyzed
description: Troubleshooting for alert PostgresqlTableNotAutoAnalyzed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlTableNotAutoAnalyzed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Table {{ $labels.relname }} has not been auto analyzed for 10 days

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PostgresqlTableNotAutoAnalyzed
expr: (pg_stat_user_tables_last_autoanalyze > 0) and (time() - pg_stat_user_tables_last_autoanalyze) > 24 * 60 * 60 * 10
for: 0m
labels:
    severity: warning
annotations:
    summary: Postgresql table not auto analyzed (instance {{ $labels.instance }})
    description: |-
        Table {{ $labels.relname }} has not been auto analyzed for 10 days
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/PostgresqlTableNotAutoAnalyzed

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
