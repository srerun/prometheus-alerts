---
title: RabbitmqTooManyReadyMessages
description: Troubleshooting for alert RabbitmqTooManyReadyMessages
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyReadyMessages

## Meaning
[//]: # "Short paragraph that explains what the alert means"
RabbitMQ too many ready messages on {{ $labels.instace }}

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqTooManyReadyMessages" %}}

<!-- Rule when generated

```yaml
alert: RabbitmqTooManyReadyMessages
expr: sum(rabbitmq_queue_messages_ready) BY (queue) > 1000
for: 1m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many ready messages (instance {{ $labels.instance }})
    description: |-
        RabbitMQ too many ready messages on {{ $labels.instace }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqTooManyReadyMessages.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
