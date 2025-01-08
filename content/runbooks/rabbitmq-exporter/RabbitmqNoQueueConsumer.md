---
title: RabbitmqNoQueueConsumer
description: Troubleshooting for alert RabbitmqNoQueueConsumer
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqNoQueueConsumer

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A queue has less than 1 consumer

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqNoQueueConsumer" %}}

{{% comment %}}

```yaml
alert: RabbitmqNoQueueConsumer
expr: rabbitmq_queue_consumers < 1
for: 1m
labels:
    severity: warning
annotations:
    summary: RabbitMQ no queue consumer (instance {{ $labels.instance }})
    description: |-
        A queue has less than 1 consumer
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqNoQueueConsumer.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
