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


Here is a runbook for the Prometheus alert rule "PostgresqlBloatIndexHigh (>80%)" :

## Meaning

The PostgresqlBloatIndexHigh (>80%) alert is triggered when the percentage of bloat in a specific PostgreSQL index is higher than 80%. Index bloat occurs when the index grows excessively large due to continued inserts, updates, and deletes, leading to decreased query performance and increased storage usage.

## Impact

A bloated index can significantly impact database performance, leading to:

* Slower query execution times
* Increased storage usage
* Decreased overall database responsiveness
* Potential crashes or downtime due to the index becoming too large to handle

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected index using the `idxname` label in the alert.
2. Check the `pg_bloat_btree_real_size` metric to determine the actual size of the index.
3. Verify that the bloat percentage is indeed higher than 80%.
4. Review the database logs to identify any recent changes or anomalies that may have contributed to the bloat.

## Mitigation

To mitigate the issue, follow these steps:

1. Execute the `REINDEX INDEX CONCURRENTLY {{ $labels.idxname }};` command to rebuild the affected index. This command will lock the table for a short period, so it's recommended to run it during a maintenance window.
2. Monitor the `pg_bloat_btree_bloat_pct` metric to ensure the bloat percentage returns to a normal level (less than 80%).
3. Consider implementing regular maintenance tasks, such as running `VACUUM` and `REINDEX` commands, to prevent index bloat in the future.
4. Review database configuration and tuning parameters to optimize performance and prevent future bloat issues.