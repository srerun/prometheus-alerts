---
title: RabbitmqTooManyMessagesInQueue
description: Troubleshooting for alert RabbitmqTooManyMessagesInQueue
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyMessagesInQueue

Queue is filling up (> 1000 msgs)

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqTooManyMessagesInQueue" %}}

{{% comment %}}

```yaml
alert: RabbitmqTooManyMessagesInQueue
expr: rabbitmq_queue_messages_ready{queue="my-queue"} > 1000
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many messages in queue (instance {{ $labels.instance }})
    description: |-
        Queue is filling up (> 1000 msgs)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqTooManyMessagesInQueue.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `RabbitmqTooManyMessagesInQueue`:

## Meaning

This alert is triggered when the number of messages in a specific RabbitMQ queue (`my-queue`) exceeds 1000 for more than 2 minutes. This indicates that the queue is filling up and may lead to performance issues or even message loss.

## Impact

The impact of this alert is high, as it can lead to:

* Performance degradation of the RabbitMQ cluster
* Message loss or delay
* Increased latency for producers and consumers of the queue
* Potential cascading failures in dependent systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ queue metrics to identify the root cause of the message accumulation:
	* Are there any unexpected spikes in message production?
	* Are consumers experiencing issues or are they slower than usual?
	* Are there any issues with the RabbitMQ cluster itself (e.g., node failures, network connectivity problems)?
2. Investigate the message contents and types to identify any patterns or anomalies:
	* Are there any specific message types or patterns that are contributing to the queue buildup?
	* Are there any issues with message processing or handling?
3. Review system logs and monitoring data to identify any correlated issues:
	* Are there any errors or warnings related to RabbitMQ or dependent systems?
	* Are there any signs of resource constraints (e.g., CPU, memory, disk space)?

## Mitigation

To mitigate the issue, follow these steps:

1. **Alert relevant teams**: Notify the teams responsible for RabbitMQ and dependent systems to investigate and take corrective action.
2. **Reduce message production**: If possible, reduce the rate of message production or pause message production temporarily to allow the queue to drain.
3. **Increase consumer capacity**: Scale up or optimize consumer instances to handle the message backlog.
4. **Purge messages**: If necessary, manually purge messages from the queue to prevent further accumulation.
5. **Investigate and resolve underlying issues**: Identify and address any underlying issues causing the message accumulation, such as issues with message processing or handling.

Remember to update this runbook with specific procedures and contacts relevant to your organization.