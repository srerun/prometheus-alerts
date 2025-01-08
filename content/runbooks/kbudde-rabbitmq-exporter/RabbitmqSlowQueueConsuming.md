---
title: RabbitmqSlowQueueConsuming
description: Troubleshooting for alert RabbitmqSlowQueueConsuming
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqSlowQueueConsuming

Queue messages are consumed slowly (> 60s)

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqSlowQueueConsuming" %}}

{{% comment %}}

```yaml
alert: RabbitmqSlowQueueConsuming
expr: time() - rabbitmq_queue_head_message_timestamp{queue="my-queue"} > 60
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ slow queue consuming (instance {{ $labels.instance }})
    description: |-
        Queue messages are consumed slowly (> 60s)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqSlowQueueConsuming.md

```

{{% /comment %}}

</details>


## Meaning

The RabbitMQ Slow Queue Consuming alert is triggered when the time elapsed since the last message was consumed from the "my-queue" queue exceeds 60 seconds. This alert indicates that the queue is not being consumed efficiently, which can lead to backlog of messages and potential issues with the application that relies on the queue.

## Impact

The impact of this alert includes:

* Delayed processing of messages in the queue, which can cause issues with the application that relies on the queue.
* Increased message backlog, which can lead to increased memory usage and potential errors.
* Potential impact on the overall performance and availability of the system.

## Diagnosis

To diagnose the issue, follow these steps:

* Check the RabbitMQ queue metrics to identify the root cause of the slow consumption.
* Investigate if there are any issues with the consumer application that is processing messages from the queue.
* Verify if there are any errors or issues with the RabbitMQ cluster or nodes.
* Check the system resource utilization (e.g., CPU, memory, disk) to ensure that they are within normal limits.

## Mitigation

To mitigate the issue, follow these steps:

* Investigate and resolve any issues with the consumer application that is processing messages from the queue.
* Check and optimize the RabbitMQ configuration to ensure that it is properly sized and configured for the workload.
* Consider scaling up or adding more nodes to the RabbitMQ cluster to increase processing capacity.
* Implement queue-based flow control or other mechanisms to prevent message backlog and ensure efficient consumption.
* Monitor the queue metrics and system resources to ensure that the issue is resolved and does not recur.