---
title: MysqlHighThreadsRunning
description: Troubleshooting for alert MysqlHighThreadsRunning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlHighThreadsRunning

More than 60% of MySQL connections are in running state on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlHighThreadsRunning" %}}

{{% comment %}}

```yaml
alert: MysqlHighThreadsRunning
expr: max_over_time(mysql_global_status_threads_running[1m]) / mysql_global_variables_max_connections * 100 > 60
for: 2m
labels:
    severity: warning
annotations:
    summary: MySQL high threads running (instance {{ $labels.instance }})
    description: |-
        More than 60% of MySQL connections are in running state on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlHighThreadsRunning.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the MysqlHighThreadsRunning alert:

## Meaning

The MysqlHighThreadsRunning alert is triggered when the percentage of running threads in MySQL exceeds 60% of the maximum allowed connections for a period of 2 minutes. This indicates that the MySQL instance is experiencing high load and may be at risk of performance degradation or even crashing.

## Impact

A high number of running threads in MySQL can lead to:

* Performance degradation: High thread count can cause increased latency and slower query execution times.
* Resource exhaustion: Running threads consume system resources, which can lead to exhaustion of CPU, memory, or disk resources, ultimately causing the MySQL instance to crash.
* Connection timeouts: If the thread count continues to rise, new connections may timeout, leading to errors and impacting application performance.

## Diagnosis

To diagnose the root cause of the MysqlHighThreadsRunning alert, follow these steps:

1. Check the MySQL error logs for any errors or warnings that may indicate the cause of the high thread count.
2. Verify that the MySQL configuration is optimal for the current workload. Check the values of `max_connections`, `thread_cache_size`, and `sort_buffer_size`.
3. Analyze the MySQL slow query log to identify any slow or long-running queries that may be contributing to the high thread count.
4. Check the system resource utilization (CPU, memory, disk) to ensure that the system is not experiencing resource constraints.
5. Verify that the MySQL instance is properly tuned for the underlying hardware and storage configuration.

## Mitigation

To mitigate the MysqlHighThreadsRunning alert, follow these steps:

1. **Improve MySQL configuration**: Adjust the MySQL configuration to optimize thread handling. Consider increasing `thread_cache_size` or `max_connections` if the system resources allow it.
2. **Optimize slow queries**: Identify and optimize slow queries to reduce their execution time and minimize their impact on the thread count.
3. **Scale up instance resources**: If the system resource utilization is high, consider scaling up the instance resources (CPU, memory, disk) to handle the increased load.
4. **Implement connection pooling**: Implement connection pooling to reduce the number of new connections and minimize the impact of thread creation.
5. **Monitor and analyze**: Continuously monitor the MySQL instance and analyze the metrics to identify trends and patterns that may indicate potential issues before they become critical.