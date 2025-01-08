---
title: NatsSlowConsumers
description: Troubleshooting for alert NatsSlowConsumers
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsSlowConsumers

There are slow consumers in NATS for {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsSlowConsumers" %}}

{{% comment %}}

```yaml
alert: NatsSlowConsumers
expr: gnatsd_varz_slow_consumers > 0
for: 3m
labels:
    severity: critical
annotations:
    summary: Nats slow consumers (instance {{ $labels.instance }})
    description: |-
        There are slow consumers in NATS for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsSlowConsumers.md

```

{{% /comment %}}

</details>


Here is the runbook for the NatsSlowConsumers alert:

## Meaning

This alert is triggered when the NATS cluster has slow consumers, which can impact the performance and reliability of the messaging system. Slow consumers can cause messages to accumulate in the NATS cluster, leading to increased latency and potentially causing instability in the system.

## Impact

The impact of slow consumers in NATS can be significant, leading to:

* Increased latency for message processing
* Accumulation of messages in the NATS cluster, potentially causing instability
* Potential loss of messages if the cluster becomes overwhelmed
* Impact on dependent systems that rely on NATS for messaging

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NATS cluster performance metrics, such as message latency and queue depth, to identify the extent of the issue.
2. Investigate the specific consumers that are slow, using the `gnatsd_varz_slow_consumers` metric and other relevant logs and metrics.
3. Determine the root cause of the slow consumers, such as high CPU usage, network issues, or application-level problems.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and restart or replace the slow consumers to restore normal performance.
2. Implement measures to prevent slow consumers in the future, such as:
	* Optimizing consumer application performance
	* Implementing load balancing or scaling for consumers
	* Improving network connectivity and reducing latency
	* Implementing timeouts and retries for consumer connections
3. Consider implementing additional monitoring and alerting for NATS cluster performance and consumer health to detect issues early.
4. Review and adjust the NATS cluster configuration to ensure it is optimized for performance and can handle the message volume.