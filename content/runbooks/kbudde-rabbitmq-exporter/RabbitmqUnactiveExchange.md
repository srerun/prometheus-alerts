---
title: RabbitmqUnactiveExchange
description: Troubleshooting for alert RabbitmqUnactiveExchange
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqUnactiveExchange

Exchange receive less than 5 msgs per second

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqUnactiveExchange" %}}

{{% comment %}}

```yaml
alert: RabbitmqUnactiveExchange
expr: rate(rabbitmq_exchange_messages_published_in_total{exchange="my-exchange"}[1m]) < 5
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ unactive exchange (instance {{ $labels.instance }})
    description: |-
        Exchange receive less than 5 msgs per second
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqUnactiveExchange.md

```

{{% /comment %}}

</details>


Here is the runbook for the RabbitMQ Unactive Exchange alert:

## Meaning

The RabbitmqUnactiveExchange alert is triggered when the rate of messages published to a specific exchange ("my-exchange") in RabbitMQ falls below 5 messages per second. This suggests that the exchange is not actively receiving messages, which may indicate a problem with the message producer or the exchange configuration.

## Impact

This alert has a significant impact on the application or service that relies on the messages published to "my-exchange". If the exchange is not receiving messages, it may lead to:

* Delayed or lost messages
* Application downtime or errors
* Data inconsistencies
* Business process disruptions

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ dashboard or the RabbitMQ exporter metrics to verify that the exchange is not receiving messages.
2. Investigate the message producer application or service to ensure it is functioning correctly.
3. Review the exchange configuration to ensure it is properly set up and configured.
4. Check the RabbitMQ server logs for any errors or warnings related to the exchange.
5. Verify that the network connection between the message producer and RabbitMQ is stable and working correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the message producer application or service if it's not functioning correctly.
2. Verify and update the exchange configuration if it's not properly set up.
3. Investigate and resolve any network connectivity issues between the message producer and RabbitMQ.
4. Increase the message publishing rate or adjust the alert threshold if the current rate is too low.
5. Consider implementing additional monitoring and alerting to detect similar issues in other exchanges or queues.

By following these steps, you should be able to identify and resolve the underlying cause of the RabbitmqUnactiveExchange alert and restore normal operation of the affected application or service.