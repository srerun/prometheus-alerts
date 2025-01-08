---
title: MysqlTooManyConnections(>80%)
description: Troubleshooting for alert MysqlTooManyConnections(>80%)
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlTooManyConnections(>80%)

More than 80% of MySQL connections are in use on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlTooManyConnections(>80%)" %}}

{{% comment %}}

```yaml
alert: MysqlTooManyConnections(>80%)
expr: max_over_time(mysql_global_status_threads_connected[1m]) / mysql_global_variables_max_connections * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: MySQL too many connections (> 80%) (instance {{ $labels.instance }})
    description: |-
        More than 80% of MySQL connections are in use on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlTooManyConnections(>80%).md

```

{{% /comment %}}

</details>


## Meaning

The MysqlTooManyConnections alert is triggered when the percentage of used MySQL connections exceeds 80% of the total available connections. This alert indicates that the MySQL instance is experiencing high connection usage, which can lead to performance issues, increased latency, and even connectivity errors.

## Impact

The impact of this alert can be significant, as it can lead to:

* Slow query performance and increased latency
* Errors and timeouts when establishing new connections
* Increased memory usage and potential crashes
* Degraded overall system performance and availability

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the MySQL instance's current connection count and compare it to the maximum allowed connections.
2. Evaluate the connection usage pattern over time to identify any trends or spikes.
3. Investigate recent changes to the application or system that may be contributing to the increased connection usage.
4. Review MySQL logs for any errors or warnings related to connection issues.
5. Verify that the MySQL configuration is optimal for the current workloads and traffic.

## Mitigation

To mitigate this alert, follow these steps:

1. **Increase the max_connections variable**: Consider increasing the max_connections variable to accommodate the increased connection demand. However, be cautious of potential performance implications and ensure the instance can handle the increased load.
2. **Optimize MySQL configuration**: Review and optimize the MySQL configuration to ensure it is tuned for the current workloads and traffic.
3. **Implement connection pooling**: Consider implementing connection pooling to reduce the number of connections established and improve overall performance.
4. **Reduce application concurrency**: If possible, reduce the concurrency of the application to decrease the number of connections being established.
5. **Monitor and analyze connection usage**: Continuously monitor and analyze connection usage to identify trends and patterns, and make adjustments as needed.

By following these steps, you can mitigate the impact of high connection usage and ensure the stability and performance of your MySQL instance.