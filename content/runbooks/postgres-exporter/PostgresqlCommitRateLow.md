---
title: PostgresqlCommitRateLow
description: Troubleshooting for alert PostgresqlCommitRateLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlCommitRateLow

Postgresql seems to be processing very few transactions

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlCommitRateLow" %}}

{{% comment %}}

```yaml
alert: PostgresqlCommitRateLow
expr: rate(pg_stat_database_xact_commit[1m]) < 10
for: 2m
labels:
    severity: critical
annotations:
    summary: Postgresql commit rate low (instance {{ $labels.instance }})
    description: |-
        Postgresql seems to be processing very few transactions
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlCommitRateLow.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `PostgresqlCommitRateLow`:

## Meaning

The `PostgresqlCommitRateLow` alert is triggered when the rate of committed transactions in a PostgreSQL database falls below 10 per minute. This could indicate a problem with the database, the application, or the infrastructure that prevents transactions from being committed successfully.

## Impact

A low commit rate can have significant consequences on the application and the business. Some possible impacts include:

* Delayed or lost transactions, leading to data inconsistencies and errors
* Reduced system performance, leading to user frustration and decreased productivity
* Increased latency, leading to timeouts and failures
* Decreased system availability, leading to revenue loss and reputational damage

## Diagnosis

To diagnose the root cause of the low commit rate, follow these steps:

1. **Check the PostgreSQL logs**: Review the PostgreSQL logs to identify any errors or warnings related to transaction processing.
2. **Verify database connections**: Check the number of active connections to the database and verify that they are within expected limits.
3. **Investigate transaction latency**: Analyze the latency of transactions using tools like `pg_stat_activity` or external monitoring tools.
4. **Check disk space and I/O**: Verify that the disk space and I/O resources are sufficient to handle the transaction load.
5. **Review application logs**: Check the application logs to identify any errors or issues that may be related to transaction processing.

## Mitigation

To mitigate the effects of a low commit rate, follow these steps:

1. **Investigate and resolve underlying issues**: Address the root cause of the low commit rate, as identified during diagnosis.
2. **Increase database resources**: Temporarily increase the database resources (e.g., connections, CPU, memory) to handle the transaction load.
3. **Optimize database configuration**: Review and optimize database configuration settings, such as `max_connections` and `shared_buffers`, to improve performance.
4. **Implement transaction retry mechanisms**: Implement retry mechanisms in the application to handle temporary failures and reduce the impact of lost transactions.
5. **Notify stakeholders**: Inform stakeholders of the issue and provide regular updates on the mitigation efforts and expected resolution time.