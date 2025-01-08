---
title: PostgresqlHighRollbackRate
description: Troubleshooting for alert PostgresqlHighRollbackRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlHighRollbackRate

Ratio of transactions being aborted compared to committed is > 2 %

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlHighRollbackRate" %}}

{{% comment %}}

```yaml
alert: PostgresqlHighRollbackRate
expr: sum by (namespace,datname) ((rate(pg_stat_database_xact_rollback{datname!~"template.*|postgres",datid!="0"}[3m])) / ((rate(pg_stat_database_xact_rollback{datname!~"template.*|postgres",datid!="0"}[3m])) + (rate(pg_stat_database_xact_commit{datname!~"template.*|postgres",datid!="0"}[3m])))) > 0.02
for: 0m
labels:
    severity: warning
annotations:
    summary: Postgresql high rollback rate (instance {{ $labels.instance }})
    description: |-
        Ratio of transactions being aborted compared to committed is > 2 %
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlHighRollbackRate.md

```

{{% /comment %}}

</details>


Here is a runbook for the PostgresqlHighRollbackRate alert rule:

## Meaning

The PostgresqlHighRollbackRate alert rule is triggered when the ratio of rolled back transactions to total transactions (rolled back + committed) in a PostgreSQL database exceeds 2% over a 3-minute window. This indicates that there is a high number of transactions being aborted, which can lead to performance issues, increased latency, and potential data inconsistencies.

## Impact

A high rollback rate can have significant impacts on the performance and reliability of your PostgreSQL database:

* Increased latency: Frequent rollbacks can lead to increased latency and slower query performance.
* Reduced throughput: High rollback rates can reduce the overall throughput of your database, leading to decreased performance and responsiveness.
* Data inconsistencies: In extreme cases, high rollback rates can lead to data inconsistencies and potential data loss.

## Diagnosis

To diagnose the cause of the high rollback rate, follow these steps:

1. Check the PostgreSQL logs for errors and warnings related to transaction rollbacks.
2. Investigate recent changes to the database schema or configuration that may be contributing to the high rollback rate.
3. Use the `pg_stat_activity` view to identify long-running transactions or transactions that are frequently being rolled back.
4. Verify that the database is properly configured for concurrency and that the correct isolation level is being used.
5. Check for any signs of resource contention or excessive locking.

## Mitigation

To mitigate the high rollback rate, follow these steps:

1. Identify and address any underlying issues causing transactions to be rolled back, such as database configuration issues or schema problems.
2. Optimize database performance by adjusting configuration settings, such as increasing the `max_connections` or `shared_buffers` settings.
3. Implement retry logic or exponential backoff in applications to reduce the load on the database and minimize the impact of rollbacks.
4. Consider implementing connection pooling or load balancing to distribute the load across multiple database instances.
5. Monitor the database performance and transaction metrics closely to ensure that the high rollback rate is not indicative of a larger issue.