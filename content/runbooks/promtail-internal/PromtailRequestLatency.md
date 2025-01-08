---
title: PromtailRequestLatency
description: Troubleshooting for alert PromtailRequestLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PromtailRequestLatency

The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}s 99th percentile latency.

<details>
  <summary>Alert Rule</summary>

{{% rule "promtail/promtail-internal.yml" "PromtailRequestLatency" %}}

{{% comment %}}

```yaml
alert: PromtailRequestLatency
expr: histogram_quantile(0.99, sum(rate(promtail_request_duration_seconds_bucket[5m])) by (le)) > 1
for: 5m
labels:
    severity: critical
annotations:
    summary: Promtail request latency (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}s 99th percentile latency.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/promtail-internal/PromtailRequestLatency.md

```

{{% /comment %}}

</details>



## Meaning
The `PromtailRequestLatency` alert is triggered when the 99th percentile latency of requests to Promtail exceeds 1 second over a 5-minute period. This indicates that a subset of requests are taking significantly longer than expected to process, which could lead to degraded performance or operational issues.

## Impact
High latency in Promtail can result in delayed ingestion and processing of logs. This may cause:
- Delayed visibility into system and application logs.
- Potential loss of critical debugging or monitoring information.
- Increased resource utilization, leading to cascading failures in downstream systems.

## Diagnosis
1. **Confirm the Alert:**
   - Review the alert details and verify the affected instance(s), job, and route.
   - Check the `VALUE` provided in the alert description to understand the severity.

2. **Examine Metrics:**
   - Query the Prometheus expression:
     ```
     histogram_quantile(0.99, sum(rate(promtail_request_duration_seconds_bucket[5m])) by (le))
     ```
     - Identify which instances or routes are contributing to high latency.

3. **Review Logs:**
   - Access the logs for the Promtail instances reporting high latency.
   - Look for error messages, timeouts, or resource exhaustion indicators.

4. **Check Resource Utilization:**
   - Inspect CPU, memory, and disk usage for the affected Promtail instances.
   - Ensure sufficient resources are allocated to handle the current log ingestion volume.

5. **Network Issues:**
   - Investigate potential network latency or connectivity issues between Promtail and its upstream or downstream systems (e.g., Loki).

6. **Configuration Changes:**
   - Review recent configuration or deployment changes that might impact Promtail’s performance.

## Mitigation
1. **Scale Resources:**
   - Increase CPU, memory, or disk resources for the affected Promtail instances.
   - Add additional Promtail replicas to distribute the load.

2. **Optimize Configuration:**
   - Adjust scrape or ingestion settings to balance performance and resource usage.
   - Ensure rate-limiting and backoff mechanisms are correctly configured.

3. **Resolve Bottlenecks:**
   - Address issues in upstream systems sending logs to Promtail (e.g., excessive log volume or high-frequency writes).
   - Investigate and resolve any bottlenecks in downstream systems (e.g., Loki performance).

4. **Restart or Redeploy:**
   - Restart affected Promtail instances to clear transient issues.
   - Roll back recent changes if the issue correlates with a new deployment.

5. **Long-Term Improvements:**
   - Implement monitoring dashboards to track Promtail performance metrics.
   - Perform a root cause analysis (RCA) to address systemic issues and prevent recurrence.

## Additional Information

- **Prometheus Query:**
  ```
  histogram_quantile(0.99, sum(rate(promtail_request_duration_seconds_bucket[5m])) by (le)) > 1
  ```

