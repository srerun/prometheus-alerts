---
title: RabbitmqTooManyConsumers
description: Troubleshooting for alert RabbitmqTooManyConsumers
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyConsumers

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Queue should have only 1 consumer

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RabbitmqTooManyConsumers

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
