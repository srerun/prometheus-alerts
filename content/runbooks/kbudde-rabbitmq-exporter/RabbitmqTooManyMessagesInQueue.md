---
title: RabbitmqTooManyMessagesInQueue
description: Troubleshooting for alert RabbitmqTooManyMessagesInQueue
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyMessagesInQueue

Queue is filling up (> 1000 msgs)

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqTooManyMessagesInQueue" %}}

{{% comment %}}

```yaml
alert: RabbitmqTooManyMessagesInQueue
expr: rabbitmq_queue_messages_ready{queue="my-queue"} > 1000
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many messages in queue (instance {{ $labels.instance }})
    description: |-
        Queue is filling up (> 1000 msgs)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqTooManyMessagesInQueue.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
