---
title: SidekiqSchedulingLatencyTooHigh
description: Troubleshooting for alert SidekiqSchedulingLatencyTooHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SidekiqSchedulingLatencyTooHigh

Sidekiq jobs are taking more than 1min to be picked up. Users may be seeing delays in background processing.

<details>
  <summary>Alert Rule</summary>

{{% rule "sidekiq/strech-sidekiq-exporter.yml" "SidekiqSchedulingLatencyTooHigh" %}}

{{% comment %}}

```yaml
alert: SidekiqSchedulingLatencyTooHigh
expr: max(sidekiq_queue_latency) > 60
for: 0m
labels:
    severity: critical
annotations:
    summary: Sidekiq scheduling latency too high (instance {{ $labels.instance }})
    description: |-
        Sidekiq jobs are taking more than 1min to be picked up. Users may be seeing delays in background processing.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/strech-sidekiq-exporter/SidekiqSchedulingLatencyTooHigh.md

```

{{% /comment %}}

</details>


## Meaning

The SidekiqSchedulingLatencyTooHigh alert is triggered when the maximum sidekiq queue latency exceeds 60 seconds. This indicates that Sidekiq jobs are taking more than 1 minute to be picked up, which can result in delays in background processing.

## Impact

* Users may experience delays in background processing, leading to a degraded user experience.
* Critical business processes may be affected, causing revenue loss or other operational issues.
* The high latency can also lead to job processing failures, causing data inconsistencies and further cascading failures.

## Diagnosis

* Check the Sidekiq queue latency metrics in Prometheus to identify the specific queue(s) experiencing high latency.
* Investigate the root cause of the high latency, such as:
	+ High CPU usage or memory pressure on the Sidekiq node.
	+ Network connectivity issues or high latency between nodes.
	+ Too many pending jobs in the queue, leading to congestion.
	+ Misconfigured Sidekiq settings or worker pool size.
* Review the Sidekiq logs for any errors or exceptions that may indicate the cause of the high latency.

## Mitigation

* Immediately investigate and address the root cause of the high latency to minimize the impact on users and business processes.
* Consider increasing the Sidekiq worker pool size to process jobs more efficiently.
* Optimize Sidekiq settings, such as the concurrency or timeout values, to improve job processing performance.
* Implement load balancing or queue sharding to distribute the job processing load and reduce latency.
* Consider implementing a circuit breaker or other resilience mechanisms to prevent cascading failures.
* Monitor the Sidekiq queue latency metrics closely to ensure the mitigation steps are effective and make adjustments as needed.