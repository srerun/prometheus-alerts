---
title: PostgresqlHighRateDeadlock
description: Troubleshooting for alert PostgresqlHighRateDeadlock
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlHighRateDeadlock

Postgres detected deadlocks

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlHighRateDeadlock" %}}

{{% comment %}}

```yaml
alert: PostgresqlHighRateDeadlock
expr: increase(postgresql_errors_total{type="deadlock_detected"}[1m]) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql high rate deadlock (instance {{ $labels.instance }})
    description: |-
        Postgres detected deadlocks
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlHighRateDeadlock.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule `PostgresqlHighRateDeadlock`:

## Meaning

A high rate of deadlocks has been detected in the PostgreSQL database. Deadlocks occur when two or more transactions are blocked, each waiting for the other to release a resource. This can lead to performance issues, errors, and even crashes.

## Impact

* Performance degradation: Deadlocks can cause transactions to stall, leading to increased response times and decreased throughput.
* Errors and failures: Deadlocks can cause transactions to fail, resulting in errors and potential data loss.
* System instability: In severe cases, deadlocks can cause the database to become unresponsive or crash.

## Diagnosis

To diagnose the root cause of the deadlocks, follow these steps:

1. Check the PostgreSQL error logs for deadlock messages.
2. Investigate the transactions involved in the deadlock using the `pg_locks` and `pg_stat_activity` views.
3. Analyze the database schema and transaction patterns to identify potential bottlenecks and conflicts.
4. Review the database configuration and tuning parameters, such as `max_locks_per_transaction` and `deadlock_timeout`.

## Mitigation

To mitigate the high rate of deadlocks, follow these steps:

1. **Immediate mitigation**: Restart the PostgreSQL database or roll back transactions to release any held locks.
2. **Short-term mitigation**: Adjust database configuration and tuning parameters, such as increasing `max_locks_per_transaction` or decreasing `deadlock_timeout`, to reduce the likelihood of deadlocks.
3. **Long-term mitigation**: Implement changes to the database schema and transaction patterns to reduce contention and bottlenecks. This may involve rewriting queries, adding indexes, or modifying the database design.
4. **Monitoring and prevention**: Implement proactive monitoring and alerting to detect potential deadlocks earlier, and take preventative measures such as restructuring transactions or adding retry mechanisms.

Remember to investigate and address the root cause of the deadlocks to prevent future occurrences.