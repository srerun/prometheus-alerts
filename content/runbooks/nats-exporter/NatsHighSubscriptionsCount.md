---
title: NatsHighSubscriptionsCount
description: Troubleshooting for alert NatsHighSubscriptionsCount
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighSubscriptionsCount

High number of NATS subscriptions ({{ $value }}) for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighSubscriptionsCount" %}}

{{% comment %}}

```yaml
alert: NatsHighSubscriptionsCount
expr: gnatsd_connz_subscriptions > 50
for: 3m
labels:
    severity: warning
annotations:
    summary: Nats high subscriptions count (instance {{ $labels.instance }})
    description: |-
        High number of NATS subscriptions ({{ $value }}) for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighSubscriptionsCount.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `NatsHighSubscriptionsCount`:

## Meaning

The `NatsHighSubscriptionsCount` alert is triggered when the number of NATS subscriptions exceeds 50 for an instance. This alert indicates that there is a high number of subscriptions to NATS subjects, which can cause performance issues and increased memory usage.

## Impact

A high number of NATS subscriptions can lead to:

* Increased memory usage, potentially causing performance issues or even crashes
* Slower message processing times, affecting the overall system performance
* Potential bottlenecks in the NATS cluster, leading to message loss or delays

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NATS cluster performance metrics, such as CPU usage, memory usage, and message throughput.
2. Investigate the subscription patterns to identify which subjects are causing the high subscription count.
3. Check the application logs for any errors or warnings related to NATS connections or message processing.
4. Verify that the NATS configuration is correct and optimal for the current workload.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the root cause of the high subscription count and take corrective action:
	* If it's due to a misconfigured application, adjust the configuration to reduce the number of subscriptions.
	* If it's due to a sudden increase in workload, consider scaling up the NATS cluster or optimizing the application to handle the load.
2. Implement subscription filtering or subject pruning to reduce the number of unnecessary subscriptions.
3. Consider implementing a subscription limit or quota to prevent excessive subscriptions.
4. Monitor the NATS cluster performance and adjust the configuration as needed to ensure optimal performance.

Additional resources:

* [NATS documentation: Subscription Management](https://docs.nats.io/running-a-nats-service/subscription-management)
* [NATS best practices: Subscription Patterns](https://docs.nats.io/running-a-nats-service/best-practices/subscription-patterns)