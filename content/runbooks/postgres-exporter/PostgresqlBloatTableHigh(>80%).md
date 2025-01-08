---
title: PostgresqlBloatTableHigh(>80%)
description: Troubleshooting for alert PostgresqlBloatTableHigh(>80%)
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlBloatTableHigh(>80%)

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The table {{ $labels.relname }} is bloated. You should execute `VACUUM {{ $labels.relname }};`

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlBloatTableHigh(>80%)" %}}

{{% comment %}}

```yaml
alert: PostgresqlBloatTableHigh(>80%)
expr: pg_bloat_table_bloat_pct > 80 and on (relname) (pg_bloat_table_real_size > 200000000)
for: 1h
labels:
    severity: warning
annotations:
    summary: Postgresql bloat table high (> 80%) (instance {{ $labels.instance }})
    description: |-
        The table {{ $labels.relname }} is bloated. You should execute `VACUUM {{ $labels.relname }};`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlBloatTableHigh(>80%).md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
