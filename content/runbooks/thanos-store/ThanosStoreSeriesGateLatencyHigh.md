---
title: ThanosStoreSeriesGateLatencyHigh
description: Troubleshooting for alert ThanosStoreSeriesGateLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosStoreSeriesGateLatencyHigh

Thanos Store {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for store series gate requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-store.yml" "ThanosStoreSeriesGateLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosStoreSeriesGateLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(thanos_bucket_store_series_gate_duration_seconds_bucket{job=~".*thanos-store.*"}[5m]))) > 2 and sum by (job) (rate(thanos_bucket_store_series_gate_duration_seconds_count{job=~".*thanos-store.*"}[5m])) > 0)
for: 10m
labels:
    severity: warning
annotations:
    summary: Thanos Store Series Gate Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Store {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for store series gate requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-store/ThanosStoreSeriesGateLatencyHigh.md

```

{{% /comment %}}

</details>


Here is a runbook for the ThanosStoreSeriesGateLatencyHigh alert rule:

## Meaning

The ThanosStoreSeriesGateLatencyHigh alert is triggered when the 99th percentile latency of store series gate requests in Thanos Store exceeds 2 seconds and there are more than 0 requests in the last 5 minutes. This indicates that the Thanos Store is experiencing high latency when handling store series gate requests, which may impact the overall performance and reliability of the system.

## Impact

High latency in store series gate requests can cause:

* Increased response times for queries and writes to the Thanos Store
* Reduced throughput and performance of dependent systems
* Potential data inconsistency and errors due to timeouts or retries
* Increased load on the Thanos Store and its dependencies

## Diagnosis

To diagnose the root cause of the high latency, follow these steps:

1. Check the Thanos Store log for any errors or warnings related to store series gate requests
2. Verify that the Thanos Store is properly configured and has sufficient resources (e.g., CPU, memory, disk space)
3. Investigate any recent changes or updates to the Thanos Store or its dependencies
4. Check the network connectivity and latency between the Thanos Store and its clients
5. Analyze the performance metrics of the Thanos Store and its dependencies to identify any bottlenecks

## Mitigation

To mitigate the high latency, follow these steps:

1. Increase the resources (e.g., CPU, memory, disk space) allocated to the Thanos Store
2. Optimize the configuration of the Thanos Store and its dependencies for better performance
3. Implement caching or other performance optimization techniques for store series gate requests
4. Reduce the load on the Thanos Store by scaling out or distributing the traffic
5. Consider implementing a queue or buffer to handle store series gate requests to reduce latency and improve overall system reliability