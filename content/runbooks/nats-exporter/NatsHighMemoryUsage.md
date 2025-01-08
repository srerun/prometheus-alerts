---
title: NatsHighMemoryUsage
description: Troubleshooting for alert NatsHighMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighMemoryUsage

NATS server memory usage is above 200MB for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighMemoryUsage" %}}

{{% comment %}}

```yaml
alert: NatsHighMemoryUsage
expr: gnatsd_varz_mem > 200 * 1024 * 1024
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high memory usage (instance {{ $labels.instance }})
    description: |-
        NATS server memory usage is above 200MB for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighMemoryUsage.md

```

{{% /comment %}}

</details>


Here is a runbook for the NatsHighMemoryUsage alert rule:

## Meaning

The NatsHighMemoryUsage alert is triggered when the memory usage of a NATS server exceeds 200MB. This alert indicates that the NATS server is consuming an abnormal amount of memory, which can lead to performance issues, slow responses, and potentially even crashes.

## Impact

High memory usage on a NATS server can have several impacts on the system:

* Decreased performance: Excessive memory consumption can slow down the NATS server, leading to slower message processing times and delayed responses.
* Increased latency: High memory usage can cause the NATS server to spend more time garbage collecting, leading to increased latency and potentially even message loss.
* Increased risk of crashes: If memory usage continues to grow unchecked, the NATS server may eventually crash or become unresponsive, leading to downtime and data loss.

## Diagnosis

To diagnose the root cause of high memory usage on a NATS server, follow these steps:

1. Check the NATS server logs for any error messages or signs of memory-related issues.
2. Investigate recent changes to the NATS server configuration, such as changes to the message queue size or clustering settings.
3. Review the NATS server metrics, such as the message queue size, connection count, and subscriber count, to identify potential bottlenecks.
4. Use debugging tools, such as `nats-top` or `nats-queue-metrics`, to gain more insight into the NATS server's memory usage and message flow.

## Mitigation

To mitigate high memory usage on a NATS server, follow these steps:

1. Reduce the message queue size: Decrease the maximum number of messages allowed in the queue to prevent memory usage from growing further.
2. Implement message queuing limits: Set limits on the number of messages that can be queued per subject to prevent excessive memory consumption.
3. Optimize clustering settings: Adjust clustering settings, such as the number of nodes or the clustering algorithm, to reduce memory usage and improve performance.
4. Restart the NATS server: If all else fails, restart the NATS server to clear out any existing memory usage and restart with a clean slate.

Remember to monitor the NATS server's memory usage and adjust these settings as needed to ensure optimal performance and prevent further alerts.