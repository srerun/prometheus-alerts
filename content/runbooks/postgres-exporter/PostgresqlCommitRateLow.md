---
title: PostgresqlCommitRateLow
description: Troubleshooting for alert PostgresqlCommitRateLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlCommitRateLow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Postgresql seems to be processing very few transactions

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlCommitRateLow" %}}

<!-- Rule when generated

```yaml
alert: PostgresqlCommitRateLow
expr: rate(pg_stat_database_xact_commit[1m]) < 10
for: 2m
labels:
    severity: critical
annotations:
    summary: Postgresql commit rate low (instance {{ $labels.instance }})
    description: |-
        Postgresql seems to be processing very few transactions
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlCommitRateLow.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
