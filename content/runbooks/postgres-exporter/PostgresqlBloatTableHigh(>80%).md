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


Here is a runbook for the Prometheus alert rule:

## Meaning

This alert rule is triggered when a PostgreSQL table is experiencing high bloat, meaning that the table has grown excessively large and is no longer optimal for querying. This can cause performance issues and slow down queries.

## Impact

* Slow query performance
* Increased disk usage
* Potential crashes or timeouts
* Decreased overall system performance

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected table: Check the `relname` label in the alert to determine which table is experiencing bloat.
2. Check table size: Use the `pg_bloat_table_real_size` metric to see how large the table has grown.
3. Run `VACUUM` command: Execute the `VACUUM` command on the affected table to remove any dead tuples and reclaim space.
4. Investigate query patterns: Review query logs to identify if there are any inefficient queries that may be contributing to the bloat.

## Mitigation

To mitigate the issue, follow these steps:

1. Run `VACUUM` command: Execute the `VACUUM` command on the affected table to remove any dead tuples and reclaim space.
2. Reindex the table: Run the `REINDEX` command on the affected table to rebuild the index and improve query performance.
3. Optimize queries: Review and optimize any inefficient queries that may be contributing to the bloat.
4. Schedule regular maintenance: Set up regular maintenance tasks to run `VACUUM` and `REINDEX` commands to prevent bloat from occurring in the future.

Remember to monitor the table's size and performance after taking these mitigation steps to ensure the issue is resolved.