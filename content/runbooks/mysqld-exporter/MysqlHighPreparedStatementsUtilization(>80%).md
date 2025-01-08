---
title: MysqlHighPreparedStatementsUtilization(>80%)
description: Troubleshooting for alert MysqlHighPreparedStatementsUtilization(>80%)
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlHighPreparedStatementsUtilization(>80%)

High utilization of prepared statements (>80%) on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlHighPreparedStatementsUtilization(>80%)" %}}

{{% comment %}}

```yaml
alert: MysqlHighPreparedStatementsUtilization(>80%)
expr: max_over_time(mysql_global_status_prepared_stmt_count[1m]) / mysql_global_variables_max_prepared_stmt_count * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: MySQL high prepared statements utilization (> 80%) (instance {{ $labels.instance }})
    description: |-
        High utilization of prepared statements (>80%) on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlHighPreparedStatementsUtilization(>80%).md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule:

## Meaning

The "MysqlHighPreparedStatementsUtilization" alert is triggered when the utilization of prepared statements in MySQL exceeds 80%. This means that the proportion of prepared statements in use compared to the maximum allowed has surpasssed the threshold, which may indicate a potential performance issue or resource bottleneck.

## Impact

A high prepared statements utilization can lead to:

* Increased memory usage and potential out-of-memory errors
* Decreased query performance and slower response times
* Increased risk of downtime and service unavailability
* Potential security risks if prepared statements are not properly sanitized

## Diagnosis

To diagnose the root cause of the high prepared statements utilization, follow these steps:

1. **Check the MySQL error logs** for any errors related to prepared statements or memory usage.
2. **Verify the current configuration** of the `max_prepared_stmt_count` variable to ensure it is set appropriately for the workload.
3. **Analyze query performance** using tools like the MySQL slow query log or a query analyzer to identify slow or inefficient queries that may be contributing to the high utilization.
4. **Investigate recent changes** to the workload or application code that may be causing an increase in prepared statement usage.

## Mitigation

To mitigate the high prepared statements utilization, follow these steps:

1. **Increase the `max_prepared_stmt_count`** variable to a higher value, but be cautious of potential memory usage increases.
2. **Optimize queries** to reduce the number of prepared statements in use, focusing on slow or inefficient queries identified during diagnosis.
3. **Implement query caching** to reduce the number of times prepared statements need to be executed.
4. **Consider implementing connection pooling** to reduce the number of connections and prepared statements in use.
5. **Monitor the situation** and adjust the mitigation strategies as needed to ensure the utilization remains below the threshold.