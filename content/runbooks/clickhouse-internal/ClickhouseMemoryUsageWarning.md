---
title: ClickhouseMemoryUsageWarning
description: Troubleshooting for alert ClickhouseMemoryUsageWarning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseMemoryUsageWarning

Memory usage is over 80%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseMemoryUsageWarning" %}}

{{% comment %}}

```yaml
alert: ClickhouseMemoryUsageWarning
expr: ClickHouseAsyncMetrics_CGroupMemoryUsed / ClickHouseAsyncMetrics_CGroupMemoryTotal * 100 > 80
for: 5m
labels:
    severity: warning
annotations:
    summary: ClickHouse Memory Usage Warning (instance {{ $labels.instance }})
    description: |-
        Memory usage is over 80%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseMemoryUsageWarning.md

```

{{% /comment %}}

</details>


Here is a runbook for the ClickhouseMemoryUsageWarning alert:

## Meaning

The ClickhouseMemoryUsageWarning alert is triggered when the memory usage of a ClickHouse instance exceeds 80% of the total available memory. This alert is raised to prevent potential performance issues and out-of-memory errors that can lead to service disruptions.

## Impact

If this alert is not addressed, high memory usage can cause:

* Slow query performance
* Increased latency
* Potential crashes or restarts of the ClickHouse instance
* Data loss or corruption
* Downtime for dependent services and applications

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ClickHouse metrics dashboard for recent changes in memory usage patterns.
2. Verify that the instance is not experiencing high CPU usage or disk I/O latency.
3. Review the query logs to identify any resource-intensive queries or jobs that may be contributing to the high memory usage.
4. Check for any recent changes to the ClickHouse configuration or schema that may be causing increased memory allocation.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and optimize resource-intensive queries or jobs that are contributing to high memory usage.
2. Consider increasing the memory allocation for the ClickHouse instance, if possible.
3. Implement memory-efficient data processing and storage practices, such as using compression and aggregates.
4. Monitor the ClickHouse instance for signs of memory leaks or other anomalies.
5. Consider scaling out the ClickHouse cluster to distribute the workload and reduce memory pressure on individual instances.

Remember to update the ClickHouse configuration and monitoring settings as needed to prevent future occurrences of this alert.