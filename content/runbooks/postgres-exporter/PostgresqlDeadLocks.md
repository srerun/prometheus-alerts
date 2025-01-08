---
title: PostgresqlDeadLocks
description: Troubleshooting for alert PostgresqlDeadLocks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlDeadLocks

PostgreSQL has dead-locks

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlDeadLocks" %}}

{{% comment %}}

```yaml
alert: PostgresqlDeadLocks
expr: increase(pg_stat_database_deadlocks{datname!~"template.*|postgres"}[1m]) > 5
for: 0m
labels:
    severity: warning
annotations:
    summary: Postgresql dead locks (instance {{ $labels.instance }})
    description: |-
        PostgreSQL has dead-locks
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlDeadLocks.md

```

{{% /comment %}}

</details>


## Meaning

The `PostgresqlDeadLocks` alert is triggered when the number of deadlocks in a PostgreSQL database exceeds 5 within a 1-minute window. Deadlocks occur when two or more transactions are blocked, each waiting for the other to release a resource. This can lead to performance degradation, slow query execution, and even crashes.

## Impact

* Performance degradation: Deadlocks can cause queries to take longer to execute, leading to increased latency and decreased throughput.
* Unavailability: In extreme cases, deadlocks can cause PostgreSQL to become unresponsive or even crash, leading to downtime and data loss.
* Data Inconsistency: Deadlocks can also lead to inconsistent data, as transactions may be rolled back or terminated unexpectedly.

## Diagnosis

To diagnose the root cause of the deadlocks:

1. Check the PostgreSQL logs for error messages related to deadlocks.
2. Use the `pg_stat_activity` view to identify the queries and transactions involved in the deadlock.
3. Analyze the database schema and query patterns to identify potential bottlenecks and hotspots.
4. Check for any ongoing maintenance operations, such as vacuuming or indexing, that may be contributing to the deadlocks.

## Mitigation

To mitigate the deadlocks:

1. **Immediate Response**: Cancel any blocked queries and restart the affected transactions.
2. **Short-Term Fix**: Adjust the `max_locks_per_transaction` setting to increase the number of locks available to transactions.
3. **Long-Term Solution**:
	* Implement row-level locking instead of table-level locking.
	* Optimize database schema and query patterns to reduce contention and hotspots.
	* Increase the ` deadlock_timeout` setting to allow transactions to wait longer before aborting.
4. **Monitoring and Prevention**: Continuously monitor PostgreSQL performance and query activity to identify potential deadlock situations before they occur.
5. **Post-Incident Analysis**: Perform a post-mortem analysis to identify the root cause of the deadlock and implement changes to prevent similar situations in the future.