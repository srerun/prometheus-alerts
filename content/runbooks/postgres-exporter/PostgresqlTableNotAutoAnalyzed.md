---
title: PostgresqlTableNotAutoAnalyzed
description: Troubleshooting for alert PostgresqlTableNotAutoAnalyzed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlTableNotAutoAnalyzed

Table {{ $labels.relname }} has not been auto analyzed for 10 days

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlTableNotAutoAnalyzed" %}}

{{% comment %}}

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
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlTableNotAutoAnalyzed.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `PostgresqlTableNotAutoAnalyzed`:

## Meaning

The `PostgresqlTableNotAutoAnalyzed` alert is triggered when a PostgreSQL table has not been auto-analyzed for 10 days. Auto-analysis is a process that updates table statistics, which are used by the query planner to optimize query execution. Without up-to-date statistics, the query planner may make suboptimal decisions, leading to poor query performance.

## Impact

The impact of not auto-analyzing a PostgreSQL table can be significant, leading to:

* Poor query performance
* Increased CPU usage
* Increased I/O usage
* Decreased overall system performance

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected table by checking the `relname` label in the alert.
2. Verify that auto-analysis is enabled for the table by checking the `pg_stat_user_tables` view.
3. Check the PostgreSQL logs for any errors or warnings related to auto-analysis.
4. Verify that the PostgreSQL instance has sufficient resources (CPU, memory, I/O) to perform auto-analysis.

## Mitigation

To mitigate the issue, follow these steps:

1. Run the `ANALYZE` command on the affected table to update its statistics.
2. Verify that auto-analysis is enabled for the table and schedule a regular maintenance window to run `ANALYZE` periodically.
3. Optimize the PostgreSQL instance configuration to ensure sufficient resources for auto-analysis.
4. Monitor the table's statistics and query performance to ensure that the issue is resolved.

Additionally, consider implementing a regular maintenance schedule to run `ANALYZE` on all tables to prevent this issue from happening in the future.