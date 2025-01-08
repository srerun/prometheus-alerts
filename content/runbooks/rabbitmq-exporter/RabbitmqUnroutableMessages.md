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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
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


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
