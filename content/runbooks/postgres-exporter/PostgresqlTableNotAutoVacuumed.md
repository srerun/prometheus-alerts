---
title: PostgresqlTableNotAutoVacuumed
description: Troubleshooting for alert PostgresqlTableNotAutoVacuumed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlTableNotAutoVacuumed

Table {{ $labels.relname }} has not been auto vacuumed for 10 days

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlTableNotAutoVacuumed" %}}

{{% comment %}}

```yaml
alert: PostgresqlTableNotAutoVacuumed
expr: (pg_stat_user_tables_last_autovacuum > 0) and (time() - pg_stat_user_tables_last_autovacuum) > 60 * 60 * 24 * 10
for: 0m
labels:
    severity: warning
annotations:
    summary: Postgresql table not auto vacuumed (instance {{ $labels.instance }})
    description: |-
        Table {{ $labels.relname }} has not been auto vacuumed for 10 days
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlTableNotAutoVacuumed.md

```

{{% /comment %}}

</details>


## Meaning
This alert triggers when a PostgreSQL table has not been auto-vacuumed for over 10 days. Auto-vacuum is a maintenance task that helps optimize table performance by reclaiming storage and updating table statistics. A failure to perform auto-vacuum can lead to bloated tables, degraded query performance, and outdated query plans.

## Impact
1. **Performance Degradation**: Queries involving the affected table may experience slower execution times due to bloated table size and inefficient index usage.
2. **Increased Storage Usage**: Lack of vacuuming can result in table and index bloat, consuming excessive disk space.
3. **Stale Query Plans**: Outdated statistics can lead to suboptimal execution plans, further impacting query performance.

## Diagnosis
1. **Identify Affected Table**:
   - Check the alert annotations to find the specific table name: `{{ $labels.relname }}`.
   - Identify the instance where the alert originated: `{{ $labels.instance }}`.

2. **Query PostgreSQL Statistics**:
   - Connect to the PostgreSQL instance.
   - Run the following query to confirm the last auto-vacuum time:
     ```sql
     SELECT relname, last_autovacuum
     FROM pg_stat_user_tables
     WHERE relname = '<TABLE_NAME>';
     ```
   - Replace `<TABLE_NAME>` with the affected table name.

3. **Check Auto-Vacuum Settings**:
   - Inspect the `autovacuum_enabled` parameter for the table:
     ```sql
     SELECT relname, reloptions
     FROM pg_class
     WHERE relname = '<TABLE_NAME>';
     ```
   - Confirm global auto-vacuum settings using:
     ```sql
     SHOW autovacuum;
     ```

4. **Inspect Workload and Activity**:
   - Review recent activity on the table to determine if heavy updates or inserts are causing bloat.
   - Check if vacuum operations are being blocked by long-running transactions:
     ```sql
     SELECT pid, state, query, age(clock_timestamp(), xact_start) AS transaction_age
     FROM pg_stat_activity
     WHERE state = 'active';
     ```

## Mitigation
1. **Manually Vacuum the Table**:
   - Run the following command to vacuum the table:
     ```sql
     VACUUM ANALYZE <TABLE_NAME>;
     ```
   - For larger tables, consider using the `VACUUM FULL` command to reclaim space, but note it requires an exclusive lock.

2. **Adjust Auto-Vacuum Settings**:
   - Increase the frequency or lower the threshold for auto-vacuum by modifying the table-specific parameters:
     ```sql
     ALTER TABLE <TABLE_NAME> SET (autovacuum_vacuum_threshold = 50, autovacuum_vacuum_scale_factor = 0.02);
     ```
   - Update global settings in the `postgresql.conf` file or via `ALTER SYSTEM`:
     ```sql
     ALTER SYSTEM SET autovacuum_vacuum_cost_limit = 200;
     ```
     Reload the configuration:
     ```bash
     SELECT pg_reload_conf();
     ```

3. **Resolve Blocking Issues**:
   - Identify and terminate blocking transactions, if necessary:
     ```sql
     SELECT pg_terminate_backend(<PID>);
     ```
     Replace `<PID>` with the process ID of the blocking transaction.

4. **Monitor and Validate**:
   - Verify the vacuum operation completed successfully by checking `pg_stat_user_tables`.
   - Confirm the alert has resolved.

5. **Long-Term Remediation**:
   - Monitor vacuum logs to identify recurring issues.
   - Adjust workload or table design to reduce bloat and improve vacuum efficiency.

---
For further details, refer to the [official documentation](https://www.postgresql.org/docs/current/routine-vacuuming.html).

