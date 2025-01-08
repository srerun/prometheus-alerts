---
title: PostgresqlHighRateStatementTimeout
description: Troubleshooting for alert PostgresqlHighRateStatementTimeout
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlHighRateStatementTimeout

Postgres transactions showing high rate of statement timeouts

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlHighRateStatementTimeout" %}}

{{% comment %}}

```yaml
alert: PostgresqlHighRateStatementTimeout
expr: rate(postgresql_errors_total{type="statement_timeout"}[1m]) > 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql high rate statement timeout (instance {{ $labels.instance }})
    description: |-
        Postgres transactions showing high rate of statement timeouts
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlHighRateStatementTimeout.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
