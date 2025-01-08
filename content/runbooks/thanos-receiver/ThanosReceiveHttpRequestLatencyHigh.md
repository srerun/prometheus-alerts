---
title: ThanosReceiveHttpRequestLatencyHigh
description: Troubleshooting for alert ThanosReceiveHttpRequestLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHttpRequestLatencyHigh

Thanos Receive {{$labels.job}} has a 99th percentile latency of {{ $value }} seconds for requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveHttpRequestLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveHttpRequestLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(http_request_duration_seconds_bucket{job=~".*thanos-receive.*", handler="receive"}[5m]))) > 10 and sum by (job) (rate(http_request_duration_seconds_count{job=~".*thanos-receive.*", handler="receive"}[5m])) > 0)
for: 10m
labels:
    severity: critical
annotations:
    summary: Thanos Receive Http Request Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} has a 99th percentile latency of {{ $value }} seconds for requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveHttpRequestLatencyHigh.md

```

{{% /comment %}}

</details>


Here is a runbook for the ThanosReceiveHttpRequestLatencyHigh alert rule:

## Meaning

The ThanosReceiveHttpRequestLatencyHigh alert rule is triggered when the 99th percentile latency of HTTP requests to Thanos Receive exceeds 10 seconds in a 5-minute window. This indicates that Thanos Receive is experiencing high latency, which can impact the performance and reliability of the system.

## Impact

High latency in Thanos Receive can lead to:

* Slow data ingestion and processing
* Increased memory usage and potential OOM errors
* Cascade failures in downstream systems
* Deterioration of overall system performance and reliability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Receive logs for errors or unusual patterns.
2. Verify that the system has sufficient resources (CPU, memory, and disk space).
3. Investigate recent changes to the system, such as software updates or configuration changes.
4. Check the network connectivity and latency between Thanos Receive and its dependencies.
5. Analyze the request volume and distribution to identify potential hotspots or bottlenecks.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the resources (CPU, memory, and disk space) allocated to Thanos Receive.
2. Optimize the system configuration to improve performance (e.g., adjust buffer sizes, tuning parameters).
3. Implement load balancing or sharding to distribute the request load.
4. Investigate and address any underlying issues causing high latency (e.g., network congestion, disk I/O bottlenecks).
5. Consider implementing caching or other optimization techniques to reduce the load on Thanos Receive.

Remember to investigate and address the root cause of the issue to prevent similar occurrences in the future.