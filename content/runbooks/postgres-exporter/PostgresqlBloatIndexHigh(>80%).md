---
title: PostgresqlBloatIndexHigh(>80%)
description: Troubleshooting for alert PostgresqlBloatIndexHigh(>80%)
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlBloatIndexHigh(>80%)

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The index {{ $labels.idxname }} is bloated. You should execute `REINDEX INDEX CONCURRENTLY {{ $labels.idxname }};`

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlBloatIndexHigh(>80%)" %}}

{{% comment %}}

```yaml
alert: PostgresqlBloatIndexHigh(>80%)
expr: pg_bloat_btree_bloat_pct > 80 and on (idxname) (pg_bloat_btree_real_size > 100000000)
for: 1h
labels:
    severity: warning
annotations:
    summary: Postgresql bloat index high (> 80%) (instance {{ $labels.instance }})
    description: |-
        The index {{ $labels.idxname }} is bloated. You should execute `REINDEX INDEX CONCURRENTLY {{ $labels.idxname }};`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlBloatIndexHigh(>80%).md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
