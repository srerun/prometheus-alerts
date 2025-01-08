---
title: RabbitmqTooManyConsumers
description: Troubleshooting for alert RabbitmqTooManyConsumers
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyConsumers

Queue should have only 1 consumer

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqTooManyConsumers" %}}

{{% comment %}}

```yaml
alert: RabbitmqTooManyConsumers
expr: rabbitmq_queue_consumers{queue="my-queue"} > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ too many consumers (instance {{ $labels.instance }})
    description: |-
        Queue should have only 1 consumer
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqTooManyConsumers.md

```

{{% /comment %}}

</details>


Here is a runbook for the RabbitMQTooManyConsumers alert:

## Meaning

The RabbitMQTooManyConsumers alert is triggered when the number of consumers connected to the "my-queue" queue exceeds 1. This indicates a potential issue with the RabbitMQ configuration or application design, as queue consumers should be limited to a single instance to prevent message duplication or other issues.

## Impact

The impact of this alert can be significant, as it may lead to:

* Message duplication or loss
* Increased resource utilization on the RabbitMQ server
* Application performance degradation
* Potential data inconsistencies or corruption

## Diagnosis

To diagnose the root cause of the alert, follow these steps:

1. Check the RabbitMQ management UI to verify the number of consumers connected to the "my-queue" queue.
2. Investigate the application logs to identify which consumers are connected to the queue and why they are not properly configured.
3. Verify the RabbitMQ configuration to ensure that the queue is properly defined and configured for a single consumer.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and terminate any unnecessary consumer connections to the "my-queue" queue using the RabbitMQ management UI or the `rabbitmqctl` command-line tool.
2. Review and update the application configuration to ensure that only a single consumer is connected to the queue.
3. Verify that the RabbitMQ configuration is correct and that the queue is properly defined for a single consumer.
4. Monitor the queue and consumer connections to ensure that the issue does not reoccur.

By following these steps, you should be able to resolve the issue and prevent future occurrences of the RabbitMQTooManyConsumers alert.