---
title: ThanosQueryInstantLatencyHigh
description: Troubleshooting for alert ThanosQueryInstantLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryInstantLatencyHigh

Thanos Query {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for instant queries.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryInstantLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosQueryInstantLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(http_request_duration_seconds_bucket{job=~".*thanos-query.*", handler="query"}[5m]))) > 40 and sum by (job) (rate(http_request_duration_seconds_bucket{job=~".*thanos-query.*", handler="query"}[5m])) > 0)
for: 10m
labels:
    severity: critical
annotations:
    summary: Thanos Query Instant Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for instant queries.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryInstantLatencyHigh.md

```

{{% /comment %}}

</details>


Here is a runbook for the ThanosQueryInstantLatencyHigh alert:

## Meaning

The ThanosQueryInstantLatencyHigh alert is triggered when the 99th percentile latency of Thanos query instant requests exceeds 40 seconds. This indicates that Thanos query instances are experiencing high latency, which can impact the performance and responsiveness of the system.

## Impact

High latency in Thanos query instant requests can have a significant impact on the system:

* Delayed query responses can lead to slow dashboard refreshes, affecting the user experience.
* Increased latency can cause timeouts, leading to failed queries and errors.
* High latency can also indicate underlying issues with the Thanos query instances, such as resource constraints or network connectivity problems.

## Diagnosis

To diagnose the root cause of the high latency, follow these steps:

1. Check the Thanos query instance logs for errors or warnings related to query execution.
2. Investigate the system resources (CPU, memory, disk) of the Thanos query instances to identify if they are experiencing resource constraints.
3. Verify the network connectivity and latency between the Thanos query instances and the storage backend.
4. Analyze the query patterns and workload to identify if there are any bottlenecks or hotspots.

## Mitigation

To mitigate the high latency issue, follow these steps:

1. **Scale Thanos query instances**: Increase the number of Thanos query instances to distribute the workload and reduce latency.
2. **Optimize query patterns**: Optimize query patterns to reduce the load on the Thanos query instances and improve performance.
3. **Tune system resources**: Adjust system resources (CPU, memory, disk) to ensure that the Thanos query instances have sufficient resources to handle the workload.
4. **Improve network connectivity**: Improve network connectivity and latency between the Thanos query instances and the storage backend.
5. **Implement caching**: Implement caching mechanisms to reduce the load on the Thanos query instances and improve performance.

Note: This is just a sample runbook, and you may need to adapt it to your specific use case and environment.