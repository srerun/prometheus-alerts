---
title: RabbitmqDeadLetterQueueFillingUp
description: Troubleshooting for alert RabbitmqDeadLetterQueueFillingUp
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqDeadLetterQueueFillingUp

Dead letter queue is filling up (> 10 msgs)

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqDeadLetterQueueFillingUp" %}}

{{% comment %}}

```yaml
alert: RabbitmqDeadLetterQueueFillingUp
expr: rabbitmq_queue_messages{queue="my-dead-letter-queue"} > 10
for: 1m
labels:
    severity: warning
annotations:
    summary: RabbitMQ dead letter queue filling up (instance {{ $labels.instance }})
    description: |-
        Dead letter queue is filling up (> 10 msgs)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqDeadLetterQueueFillingUp.md

```

{{% /comment %}}

</details>


Here is a runbook for the RabbitMQ Dead Letter Queue Filling Up alert:

## Meaning

The RabbitMQ Dead Letter Queue Filling Up alert is triggered when the number of messages in the dead letter queue exceeds 10. This queue is a special queue in RabbitMQ where messages are sent when they cannot be delivered to their intended queue. A growing dead letter queue can indicate issues with message processing, network connectivity, or configuration problems.

## Impact

If left unchecked, a growing dead letter queue can lead to:

* Increased memory usage and potential crashes of RabbitMQ nodes
* Delayed or lost messages
* Increased latency for message processing
* Difficulty in troubleshooting and debugging issues due to the accumulating messages in the dead letter queue

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ logs for errors or warnings related to message processing or network connectivity.
2. Verify that the dead letter queue is not being drained or consumed by a consumer.
3. Check the RabbitMQ node metrics for any signs of memory pressure or high CPU usage.
4. Investigate recent changes to the RabbitMQ configuration, message producers, or consumers.
5. Use the RabbitMQ management UI or CLI tools to inspect the dead letter queue and its contents.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and fix the root cause of messages being sent to the dead letter queue (e.g., fix message serialization issues, network connectivity problems, or configuration errors).
2. Configure a dead letter queue consumer or drain to process and remove messages from the queue.
3. Monitor the RabbitMQ node metrics and adjust resource allocation as needed to prevent memory pressure or high CPU usage.
4. Consider implementing retries or circuit breakers in message producers to prevent overloading the dead letter queue.
5. Verify that the dead letter queue is being properly monitored and alerted on to prevent future occurrences.