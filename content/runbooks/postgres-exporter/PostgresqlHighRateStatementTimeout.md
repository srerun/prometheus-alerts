---
title: PostgresqlHighRateStatementTimeout
description: Troubleshooting for alert PostgresqlHighRateStatementTimeout
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlHighRateStatementTimeout

Postgres transactions showing high rate of statement timeouts

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlHighRateStatementTimeout" %}}

{{% comment %}}

```yaml
alert: PostgresqlHighRateStatementTimeout
expr: rate(postgresql_errors_total{type="statement_timeout"}[1m]) > 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql high rate statement timeout (instance {{ $labels.instance }})
    description: |-
        Postgres transactions showing high rate of statement timeouts
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlHighRateStatementTimeout.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `PostgresqlHighRateStatementTimeout`:

## Meaning

The `PostgresqlHighRateStatementTimeout` alert is triggered when the rate of statement timeouts in a PostgreSQL instance exceeds 3 per minute. This alert indicates that there is a high rate of statements timing out in the database, which can lead to performance issues and potentially cause errors or crashes.

## Impact

The impact of this alert can be significant, as statement timeouts can:

* Cause transactions to fail, leading to data inconsistencies and errors
* Increase the load on the database, leading to performance degradation
* Potentially lead to cascading failures and crashes of dependent systems

## Diagnosis

To diagnose the cause of the `PostgresqlHighRateStatementTimeout` alert, follow these steps:

1. **Check the PostgreSQL logs** for errors and warnings related to statement timeouts.
2. **Analyze the query patterns** to identify which queries are causing the timeouts.
3. **Check the system resources** (CPU, memory, disk space) to ensure they are not overloaded.
4. **Verify the database configuration** to ensure that the timeout settings are correctly configured.
5. **Check for any recent changes** to the database, such as new applications or queries that may be causing the timeouts.

## Mitigation

To mitigate the `PostgresqlHighRateStatementTimeout` alert, follow these steps:

1. **Identify and optimize** the queries causing the timeouts.
2. **Adjust the timeout settings** to a more reasonable value, if necessary.
3. **Increase system resources** (CPU, memory, disk space) to handle the load.
4. **Implement query queuing** to manage the concurrent query load.
5. **Consider upgrading** the PostgreSQL instance to a more performant version.
6. **Monitor the database** closely to ensure that the issue is resolved and does not recur.