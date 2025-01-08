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

PostgreSQL dead tuples is too large

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlTooManyDeadTuples" %}}

{{% comment %}}

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

{{% /comment %}}

</details>


Here is the runbook for the PostgresqlTooManyDeadTuples alert:

## Meaning
The PostgresqlTooManyDeadTuples alert is triggered when the number of dead tuples in a PostgreSQL database exceeds 10,000 or 10% of the total number of live and dead tuples. Dead tuples are rows in a table that are no longer visible to the user due to updates or deletes. A high number of dead tuples can lead to performance issues and increased disk usage.

## Impact
If left unaddressed, a high number of dead tuples can cause:

* Performance degradation: Excessive dead tuples can lead to slower query performance, as the database needs to skip over them to retrieve live data.
* Increased disk usage: Dead tuples occupy space on disk, which can lead to storage issues if not cleaned up regularly.
* Increased I/O operations: Vacuuming dead tuples requires additional I/O operations, which can impact system performance.

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the PostgreSQL logs for any errors related to dead tuples or vacuuming.
2. Run the following query to identify the tables with the most dead tuples:
```
SELECT schemaname, relname, n_dead_tup
FROM pg_stat_user_tables
ORDER BY n_dead_tup DESC;
```
3. Verify that the `autovacuum` process is running and configured correctly. Check the `autovacuum` log files for any errors or issues.
4. Review the database configuration and check if the `vacuum_threshold` and `vacuum_scale_factor` settings are optimal for the workload.

## Mitigation
To mitigate the issue, follow these steps:

1. Run a manual `VACUUM` command on the affected tables to reclaim space and remove dead tuples:
```
VACUUM (FULL) table_name;
```
2. Adjust the `autovacuum` settings to run more frequently or with a higher threshold to prevent dead tuples from building up in the future.
3. Consider increasing the disk space available to the database to prevent storage issues.
4. Monitor the database performance and dead tuple count to ensure the issue is resolved and does not recur.

Remember to follow best practices for vacuuming and maintaining your PostgreSQL database to prevent similar issues in the future.