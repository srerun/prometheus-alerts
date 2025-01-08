---
title: ClickhouseMemoryUsageCritical
description: Troubleshooting for alert ClickhouseMemoryUsageCritical
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseMemoryUsageCritical

Memory usage is critically high, over 90%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseMemoryUsageCritical" %}}

{{% comment %}}

```yaml
alert: ClickhouseMemoryUsageCritical
expr: ClickHouseAsyncMetrics_CGroupMemoryUsed / ClickHouseAsyncMetrics_CGroupMemoryTotal * 100 > 90
for: 5m
labels:
    severity: critical
annotations:
    summary: ClickHouse Memory Usage Critical (instance {{ $labels.instance }})
    description: |-
        Memory usage is critically high, over 90%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseMemoryUsageCritical.md

```

{{% /comment %}}

</details>


Here is a runbook for the ClickhouseMemoryUsageCritical alert:

## Meaning

The ClickhouseMemoryUsageCritical alert is triggered when the memory usage of a ClickHouse instance exceeds 90% of the total available memory. This alert indicates that the instance is at risk of running out of memory, which can lead to performance issues, slow queries, and even instance crashes.

## Impact

The impact of high memory usage on a ClickHouse instance can be severe. It can cause:

* Slow query performance and increased latency
* Increased risk of instance crashes and downtime
* Potential data loss or corruption
* Degraded performance for dependent services and applications
* Increased resource utilization and costs

## Diagnosis

To diagnose the root cause of the high memory usage, follow these steps:

1. Check the ClickHouse logs for any error messages or warnings related to memory usage.
2. Investigate recent changes to the ClickHouse configuration, such as increased buffer sizes or altered merge settings.
3. Review the query patterns and identify any resource-intensive queries that may be contributing to the high memory usage.
4. Use tools such as `top` or `htop` to monitor the system resource utilization and identify any other resource bottlenecks.
5. Check the ClickHouse metrics, such as `ClickHouseAsyncMetrics_CGroupMemoryUsed` and `ClickHouseAsyncMetrics_CGroupMemoryTotal`, to understand the memory usage trend and identify any anomalies.

## Mitigation

To mitigate the high memory usage, follow these steps:

1. **Adjust ClickHouse configuration**: Review and adjust the ClickHouse configuration to optimize memory usage, such as reducing buffer sizes or altering merge settings.
2. **Optimize queries**: Identify and optimize resource-intensive queries to reduce their memory footprint.
3. **Increase instance resources**: Consider increasing the instance resources, such as adding more memory or CPU, to accommodate the workload.
4. **Implement query queuing**: Implement query queuing to prevent excessive memory usage during peak load periods.
5. **Monitor and alert**: Continuously monitor the ClickHouse memory usage and set up alerts to notify teams of potential issues before they become critical.

Remember to also review the ClickHouse documentation and best practices for optimizing memory usage and performance.