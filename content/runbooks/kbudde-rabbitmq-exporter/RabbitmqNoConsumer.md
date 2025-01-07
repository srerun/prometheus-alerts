---
title: RabbitmqNoConsumer
description: Troubleshooting for alert RabbitmqNoConsumer
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqNoConsumer

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Queue has no consumer

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RabbitmqNoConsumer

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
