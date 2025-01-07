---
title: RabbitmqDeadLetterQueueFillingUp
description: Troubleshooting for alert RabbitmqDeadLetterQueueFillingUp
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqDeadLetterQueueFillingUp

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Dead letter queue is filling up (> 10 msgs)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: RabbitmqDeadLetterQueueFillingUp
expr: rabbitmq_queue_messages{queue="my-dead-letter-queue"} > 10
for: 1m
labels:
    severity: warning
annotations:
    summary: RabbitMQ dead letter queue filling up (instance {{ $labels.instance }})
    description: |-
        Dead letter queue is filling up (> 10 msgs)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RabbitmqDeadLetterQueueFillingUp

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
