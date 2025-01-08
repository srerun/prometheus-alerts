---
title: RabbitmqTooManyReadyMessages
description: Troubleshooting for alert RabbitmqTooManyReadyMessages
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyReadyMessages

RabbitMQ too many ready messages on {{ $labels.instace }}

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqTooManyReadyMessages" %}}

{{% comment %}}

```yaml
alert: RabbitmqTooManyReadyMessages
expr: sum(rabbitmq_queue_messages_ready) BY (queue) > 1000
for: 1m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many ready messages (instance {{ $labels.instance }})
    description: |-
        RabbitMQ too many ready messages on {{ $labels.instace }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqTooManyReadyMessages.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `RabbitmqTooManyReadyMessages`:

## Meaning

The `RabbitmqTooManyReadyMessages` alert is triggered when the number of ready messages in a RabbitMQ queue exceeds 1000. This indicates that the queue is experiencing a high volume of messages that are ready to be consumed, but have not been processed yet.

## Impact

A high number of ready messages in a RabbitMQ queue can have significant impacts on the performance and reliability of the system. Some potential consequences include:

* Increased memory usage on the RabbitMQ node, leading to potential crashes or slow performance
* Delays in message processing, causing latency and affecting the overall throughput of the system
* Increased risk of message loss or corruption, leading to data integrity issues

## Diagnosis

To diagnose the root cause of the `RabbitmqTooManyReadyMessages` alert, follow these steps:

1. Check the RabbitMQ queue metrics to identify the specific queue experiencing the high volume of ready messages.
2. Investigate the message production rate and consumption rate to determine if there is an imbalance.
3. Review the application logs to identify any errors or issues that may be contributing to the buildup of ready messages.
4. Check the RabbitMQ node resource utilization (e.g., CPU, memory, disk) to ensure it is not experiencing any resource constraints.

## Mitigation

To mitigate the `RabbitmqTooManyReadyMessages` alert, take the following steps:

1. Investigate and resolve any underlying application errors or issues that may be causing the buildup of ready messages.
2. Adjust the message production rate or add additional consumer nodes to balance the message consumption rate.
3. Implement message routing or filtering to reduce the volume of messages in the affected queue.
4. Consider increasing the RabbitMQ node resources (e.g., adding more nodes, increasing memory) to handle the increased message volume.
5. Monitor the queue metrics closely to ensure the issue is resolved and the queue is returning to a normal state.