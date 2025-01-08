---
title: SidekiqQueueSize
description: Troubleshooting for alert SidekiqQueueSize
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SidekiqQueueSize

Sidekiq queue {{ $labels.name }} is growing

<details>
  <summary>Alert Rule</summary>

{{% rule "sidekiq/strech-sidekiq-exporter.yml" "SidekiqQueueSize" %}}

{{% comment %}}

```yaml
alert: SidekiqQueueSize
expr: sidekiq_queue_size > 100
for: 1m
labels:
    severity: warning
annotations:
    summary: Sidekiq queue size (instance {{ $labels.instance }})
    description: |-
        Sidekiq queue {{ $labels.name }} is growing
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/strech-sidekiq-exporter/SidekiqQueueSize.md

```

{{% /comment %}}

</details>


Here is a runbook for the SidekiqQueueSize alert rule:

## Meaning

The SidekiqQueueSize alert is triggered when the size of a Sidekiq queue exceeds 100 jobs. This alert indicates that the queue is growing and may lead to performance issues or errors in the application.

## Impact

If this alert is not addressed, it can lead to:

* Delayed or lost jobs in the queue
* Increased memory usage and potential crashes
* Slower response times and decreased system performance
* Potential errors or failures in dependent systems or services

## Diagnosis

To diagnose the root cause of the growing Sidekiq queue, follow these steps:

1. Check the Sidekiq queue metrics to identify the specific queue(s) experiencing growth.
2. Investigate recent changes to the application code or configuration that may be causing the queue growth.
3. Review system logs for errors or exceptions related to Sidekiq or the affected queue.
4. Check for any stuck or failed jobs in the queue that may be contributing to the growth.

## Mitigation

To mitigate the growing Sidekiq queue, follow these steps:

1. Identify and address the root cause of the queue growth, based on the diagnosis steps above.
2. Consider increasing the Sidekiq worker count or adjusting the queue settings to process jobs more efficiently.
3. Implement retries or circuit breakers to handle failed or stuck jobs.
4. Monitor the queue size and adjust the alert threshold or notification settings as needed.

Remember to follow best practices for Sidekiq queue management and tuning to prevent similar issues in the future.