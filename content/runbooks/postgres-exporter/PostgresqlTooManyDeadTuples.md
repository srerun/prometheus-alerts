---
title: PostgresqlTooManyDeadTuples
description: Troubleshooting for alert PostgresqlTooManyDeadTuples
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlTooManyDeadTuples

## Meaning
[//]: # "Short paragraph that explains what the alert means"
PostgreSQL dead tuples is too large

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlTooManyDeadTuples" %}}

<!-- Rule when generated

```yaml
alert: PostgresqlTooManyDeadTuples
expr: ((pg_stat_user_tables_n_dead_tup > 10000) / (pg_stat_user_tables_n_live_tup + pg_stat_user_tables_n_dead_tup)) >= 0.1
for: 2m
labels:
    severity: warning
annotations:
    summary: Postgresql too many dead tuples (instance {{ $labels.instance }})
    description: |-
        PostgreSQL dead tuples is too large
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlTooManyDeadTuples.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
