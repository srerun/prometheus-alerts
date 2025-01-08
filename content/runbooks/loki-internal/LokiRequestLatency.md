---
title: LokiRequestLatency
description: Troubleshooting for alert LokiRequestLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiRequestLatency

The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}s 99th percentile latency

<details>
  <summary>Alert Rule</summary>

{{% rule "loki/loki-internal.yml" "LokiRequestLatency" %}}

{{% comment %}}

```yaml
alert: LokiRequestLatency
expr: (histogram_quantile(0.99, sum(rate(loki_request_duration_seconds_bucket{route!~"(?i).*tail.*"}[5m])) by (le)))  > 1
for: 5m
labels:
    severity: critical
annotations:
    summary: Loki request latency (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}s 99th percentile latency
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/loki-internal/LokiRequestLatency.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the LokiRequestLatency alert:

## Meaning

The LokiRequestLatency alert is triggered when the 99th percentile request latency for Loki requests exceeds 1 second over a 5-minute window. This alert is critical, indicating a severe performance issue with Loki.

## Impact

* High latency can cause delays in log processing and querying, leading to poor user experience and delayed insights.
* This can also lead to increased memory usage and potential crashes in Loki components.
* Services relying on Loki for logging and monitoring may be impacted, causing cascading failures.

## Diagnosis

To diagnose the root cause of this issue:

1. Check the Loki instance(s) indicated in the alert label for any signs of high load, GC pauses, or resource contention.
2. Investigate recent changes to the Loki configuration, deployed applications, or infrastructure that may be contributing to the increased latency.
3. Review the Loki request logs to identify patterns or anomalies in the requests that may be causing the latency.
4. Verify that the Loki cluster is properly sized and scaled to handle the current load.

## Mitigation

To mitigate this issue:

1. Immediately investigate and address any underlying resource contention or configuration issues in the Loki instance(s).
2. Consider scaling up or horizontally scaling the Loki cluster to handle the increased load.
3. Optimize Loki configuration to improve performance, such as adjusting the ingestion rate, batch size, or query concurrency.
4. Implement caching or other optimization techniques to reduce the load on Loki.
5. Roll back recent changes to the Loki configuration or deployed applications if deemed necessary.
6. Monitor the Loki request latency and adjust the alert threshold accordingly to ensure it is set to a reasonable value for the specific Loki deployment.