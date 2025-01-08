---
title: PostgresqlTooManyLocksAcquired
description: Troubleshooting for alert PostgresqlTooManyLocksAcquired
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlTooManyLocksAcquired

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many locks acquired on the database. If this alert happens frequently, we may need to increase the postgres setting max_locks_per_transaction.

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlTooManyLocksAcquired" %}}

<!-- Rule when generated

```yaml
alert: PostgresqlTooManyLocksAcquired
expr: ((sum (pg_locks_count)) / (pg_settings_max_locks_per_transaction * pg_settings_max_connections)) > 0.20
for: 2m
labels:
    severity: critical
annotations:
    summary: Postgresql too many locks acquired (instance {{ $labels.instance }})
    description: |-
        Too many locks acquired on the database. If this alert happens frequently, we may need to increase the postgres setting max_locks_per_transaction.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlTooManyLocksAcquired.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
