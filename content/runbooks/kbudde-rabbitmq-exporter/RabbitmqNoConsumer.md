---
title: RabbitmqNoConsumer
description: Troubleshooting for alert RabbitmqNoConsumer
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqNoConsumer

Queue has no consumer

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqNoConsumer" %}}

{{% comment %}}

```yaml
alert: RabbitmqNoConsumer
expr: rabbitmq_queue_consumers == 0
for: 1m
labels:
    severity: critical
annotations:
    summary: RabbitMQ no consumer (instance {{ $labels.instance }})
    description: |-
        Queue has no consumer
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqNoConsumer.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the RabbitmqNoConsumer alert:

## Meaning

The RabbitmqNoConsumer alert is triggered when a RabbitMQ queue has no consumers. This means that messages are not being processed from the queue, which can lead to a backlog of messages and potentially cause issues with application functionality.

## Impact

The impact of this alert can be significant, as it can cause:

* Message processing delays or failures
* Increased queue length, potentially leading to performance issues or even node crashes
* Potential data loss or corruption if messages are not processed correctly
* Impact on dependent applications or services that rely on the RabbitMQ queue

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check RabbitMQ queue status**: Log in to the RabbitMQ management UI and check the status of the affected queue.
2. **Verify consumer configuration**: Check the configuration of the consumer application or service to ensure it is correctly configured to consume from the queue.
3. **Check consumer logs**: Review the consumer application logs to see if there are any error messages related to connecting to the RabbitMQ queue.
4. **Check RabbitMQ node status**: Verify that the RabbitMQ node is running and healthy.

## Mitigation

To mitigate the issue, follow these steps:

1. **Restart the consumer application**: Restart the consumer application or service to re-establish the connection to the RabbitMQ queue.
2. **Check and update consumer configuration**: Verify and update the consumer configuration to ensure it is correctly set up to consume from the queue.
3. **Check RabbitMQ node status**: Verify that the RabbitMQ node is running and healthy, and take corrective action if necessary.
4. **Monitor queue status**: Closely monitor the queue status to ensure the issue is resolved and the queue is processing messages correctly.

Additional resources:

* [RabbitMQ documentation: Queues and Consumers](https://www.rabbitmq.com/tutorials/amqp-concepts.html)
* [RabbitMQ exporter documentation: Queues](https://github.com/kbudde/rabbitmq_exporter#queues)