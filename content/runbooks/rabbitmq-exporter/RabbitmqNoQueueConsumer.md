---
title: RabbitmqNoQueueConsumer
description: Troubleshooting for alert RabbitmqNoQueueConsumer
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqNoQueueConsumer

A queue has less than 1 consumer

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqNoQueueConsumer" %}}

{{% comment %}}

```yaml
alert: RabbitmqNoQueueConsumer
expr: rabbitmq_queue_consumers < 1
for: 1m
labels:
    severity: warning
annotations:
    summary: RabbitMQ no queue consumer (instance {{ $labels.instance }})
    description: |-
        A queue has less than 1 consumer
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqNoQueueConsumer.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The RabbitmqNoQueueConsumer alert is triggered when a RabbitMQ queue has less than one consumer. This means that messages in the queue are not being processed, which can lead to a backlog of messages and impact the performance and reliability of the system.

## Impact

The impact of this alert can be significant, as it can cause:

* Messages to accumulate in the queue, leading to a backlog of unprocessed messages
* Delays in processing critical tasks and requests
* Increased latency and decreased system performance
* Potential data loss or inconsistencies if messages are not processed in a timely manner

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

1. Check the RabbitMQ queue metrics to identify the specific queue with no consumers.
2. Investigate the consumer application or service to determine why it is not consuming messages from the queue.
3. Check the consumer logs for errors or exceptions that may be preventing it from consuming messages.
4. Verify that the consumer is properly configured and connected to the RabbitMQ broker.
5. Check the RabbitMQ broker logs for any errors or issues that may be preventing consumers from connecting.

## Mitigation

To mitigate this issue, follow these steps:

1. Identify and restart the consumer application or service to re-establish connectivity to the RabbitMQ queue.
2. Verify that the consumer is properly configured and connected to the RabbitMQ broker.
3. Check the RabbitMQ queue configuration to ensure that it is properly set up and configured.
4. Consider increasing the number of consumers or improving the consumer application's performance to handle the message volume.
5. Monitor the RabbitMQ queue metrics to ensure that the issue is resolved and the queue is being properly consumed.