---
title: RabbitmqSlowQueueConsuming
description: Troubleshooting for alert RabbitmqSlowQueueConsuming
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqSlowQueueConsuming

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Queue messages are consumed slowly (> 60s)

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RabbitmqSlowQueueConsuming

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
