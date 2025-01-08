---
title: RabbitmqUnroutableMessages
description: Troubleshooting for alert RabbitmqUnroutableMessages
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqUnroutableMessages

A queue has unroutable messages

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqUnroutableMessages" %}}

{{% comment %}}

```yaml
alert: RabbitmqUnroutableMessages
expr: increase(rabbitmq_channel_messages_unroutable_returned_total[1m]) > 0 or increase(rabbitmq_channel_messages_unroutable_dropped_total[1m]) > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ unroutable messages (instance {{ $labels.instance }})
    description: |-
        A queue has unroutable messages
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqUnroutableMessages.md

```

{{% /comment %}}

</details>


Here is a runbook for the RabbitMQ Unroutable Messages alert:

## Meaning

The RabbitMQ Unroutable Messages alert indicates that there are messages in a RabbitMQ queue that cannot be routed to a destination exchange or queue. This can be due to a misconfigured routing key, a non-existent exchange or queue, or a network connectivity issue.

## Impact

Unroutable messages can cause issues with message processing and lead to data loss or duplication. If left unattended, this can result in:

* Messages being dropped or lost
* Applications experiencing errors or timeouts
* Data inconsistencies between systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ management UI for the specific queue experiencing unroutable messages.
2. Verify the routing key configuration for the queue.
3. Check if the destination exchange or queue exists and is properly configured.
4. Review the RabbitMQ logs for any error messages related to message routing.
5. Use the RabbitMQ `rabbitmqctl` command to inspect the queue and exchange configurations.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and fix any misconfigured routing keys or exchange/queue names.
2. Create or update the destination exchange or queue if it does not exist.
3. Verify that network connectivity between RabbitMQ nodes is stable and functioning correctly.
4. Consider implementing a dead-letter queue to handle unroutable messages.
5. Monitor the queue and exchange configurations regularly to prevent similar issues from occurring in the future.

Remember to update the alert annotations to reflect the correct runbook URL and instance labels.