---
title: RabbitmqTooManyUnackMessages
description: Troubleshooting for alert RabbitmqTooManyUnackMessages
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyUnackMessages

Too many unacknowledged messages

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqTooManyUnackMessages" %}}

{{% comment %}}

```yaml
alert: RabbitmqTooManyUnackMessages
expr: sum(rabbitmq_queue_messages_unacked) BY (queue) > 1000
for: 1m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many unack messages (instance {{ $labels.instance }})
    description: |-
        Too many unacknowledged messages
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqTooManyUnackMessages.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `RabbitmqTooManyUnackMessages`:

## Meaning
This alert is triggered when the number of unacknowledged messages in a RabbitMQ queue exceeds 1000. This indicates that messages are not being processed or acknowledged in a timely manner, which can lead to message loss, queue buildup, and performance issues.

## Impact
The impact of this alert can be significant, as unacknowledged messages can lead to:

* Message loss: If messages are not acknowledged, they may be lost if the RabbitMQ node restarts or crashes.
* Queue buildup: Unacknowledged messages can cause queues to build up, leading to increased memory usage, slower performance, and potentially even queue crashes.
* Performance issues: Unacknowledged messages can cause RabbitMQ to slow down, leading to delays in message processing and potential timeouts.

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the RabbitMQ queue metrics to identify the specific queue(s) with the high number of unacknowledged messages.
2. Investigate the application or service that is consuming from the affected queue(s) to determine if there are any issues or slowdowns.
3. Check the RabbitMQ node logs for any errors or exceptions related to message processing or acknowledgment.
4. Verify that the RabbitMQ node is properly configured and that the Erlang processes are running correctly.

## Mitigation
To mitigate the issue, follow these steps:

1. Investigate and resolve any issues with the application or service that is consuming from the affected queue(s).
2. Consider increasing the number of workers or consumers to process the backlog of messages.
3. Monitor the RabbitMQ queue metrics to ensure that the number of unacknowledged messages is decreasing.
4. Consider implementing a dead-letter queue to handle messages that cannot be processed or acknowledged.
5. If necessary, restart the RabbitMQ node or Erlang processes to recover from any potential issues.

Note: This runbook is just a starting point, and the specific steps and procedures may vary depending on your environment and setup.