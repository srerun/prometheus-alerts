---
title: ThanosQueryRangeLatencyHigh
description: Troubleshooting for alert ThanosQueryRangeLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryRangeLatencyHigh

Thanos Query {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for range queries.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryRangeLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosQueryRangeLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(http_request_duration_seconds_bucket{job=~".*thanos-query.*", handler="query_range"}[5m]))) > 90 and sum by (job) (rate(http_request_duration_seconds_count{job=~".*thanos-query.*", handler="query_range"}[5m])) > 0)
for: 10m
labels:
    severity: critical
annotations:
    summary: Thanos Query Range Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for range queries.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryRangeLatencyHigh.md

```

{{% /comment %}}

</details>


## Meaning

The `ThanosQueryRangeLatencyHigh` alert is triggered when the 99th percentile latency of range queries in Thanos Query exceeds 90 seconds for a specific job/instance. This alert indicates that the query range requests are experiencing high latency, which may impact the performance and responsiveness of the system.

## Impact

The high latency of range queries can have a significant impact on the overall system performance, leading to:

* Slow query responses, resulting in delayed decision-making and potentially affecting business operations
* Increased latency can cause cascading failures, leading to a broader system outage
* Poor user experience, especially for applications that rely heavily on query range requests

## Diagnosis

To diagnose the root cause of the high latency, follow these steps:

1. Check the Thanos Query logs for any error messages or exceptions related to query range requests
2. Analyze the system metrics, such as CPU utilization, memory usage, and disk I/O, to identify any resource bottlenecks
3. Verify if there are any recent changes or updates to the Thanos Query configuration, deployment, or underlying infrastructure that may be contributing to the high latency
4. Use tools like Prometheus and Grafana to visualize the query latency and identify any trends or patterns

## Mitigation

To mitigate the high latency of range queries, follow these steps:

1. **Investigate and resolve any underlying infrastructure issues**: Identify and address any resource bottlenecks, such as high CPU utilization or disk I/O issues, and ensure that the underlying infrastructure is properly scaled and configured.
2. **Optimize Thanos Query configuration**: Review and optimize the Thanos Query configuration, such as adjusting the query concurrency, caching, and indexing settings to improve performance.
3. **Implement query optimization techniques**: Apply query optimization techniques, such as query rewriting, indexing, and caching, to improve query performance.
4. **Consider horizontal scaling**: Consider horizontally scaling the Thanos Query instances to distribute the load and improve performance.
5. **Monitor and analyze query patterns**: Continuously monitor and analyze query patterns to identify opportunities for optimization and improvement.

Remember to follow the runbook's guidelines and take corrective actions to resolve the issue and prevent future occurrences.