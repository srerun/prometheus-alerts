---
title: PostgresqlTooManyLocksAcquired
description: Troubleshooting for alert PostgresqlTooManyLocksAcquired
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlTooManyLocksAcquired

Too many locks acquired on the database. If this alert happens frequently, we may need to increase the postgres setting max_locks_per_transaction.

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlTooManyLocksAcquired" %}}

{{% comment %}}

```yaml
alert: PostgresqlTooManyLocksAcquired
expr: ((sum (pg_locks_count)) / (pg_settings_max_locks_per_transaction * pg_settings_max_connections)) > 0.20
for: 2m
labels:
    severity: critical
annotations:
    summary: Postgresql too many locks acquired (instance {{ $labels.instance }})
    description: |-
        Too many locks acquired on the database. If this alert happens frequently, we may need to increase the postgres setting max_locks_per_transaction.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlTooManyLocksAcquired.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule `PostgresqlTooManyLocksAcquired`:

## Meaning

This alert is triggered when the number of locks acquired by PostgreSQL exceeds 20% of the maximum allowed locks per transaction, multiplied by the maximum number of connections. This can indicate a performance issue or a potential deadlock situation in the database.

## Impact

The impact of this alert can be significant, as it can lead to:

* Performance degradation: Excessive locks can cause queries to wait, leading to slow response times and decreased throughput.
* Deadlocks: In extreme cases, too many locks can cause deadlocks, resulting in failed transactions and data inconsistencies.
* Increased resource utilization: The database may consume more resources (e.g., CPU, memory) to manage the locks, leading to increased resource utilization.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the PostgreSQL logs for any error messages related to locks or deadlocks.
2. Run the following query to identify the top locking queries: `SELECT * FROM pg_stat_activity WHERE wait_event_type = 'Activity' AND wait_event = 'Lock';`
3. Check the current value of `max_locks_per_transaction` and `max_connections` using the following queries: `SHOW max_locks_per_transaction;` and `SHOW max_connections;`
4. Review the database connection pool configuration to ensure it is properly sized for the workload.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the value of `max_locks_per_transaction` to allow for more locks per transaction. However, be cautious when increasing this value, as it can lead to increased resource utilization.
2. Adjust the database connection pool configuration to reduce the number of connections and alleviate the locking pressure.
3. Optimize the database schema and queries to reduce the need for locks.
4. Consider implementing lock timeouts or deadlock detection mechanisms to prevent prolonged lock waits.

Remember to monitor the database performance and adjust the configuration settings as needed to prevent similar issues in the future.