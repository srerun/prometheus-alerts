---
title: RabbitmqTooManyUnackMessages
description: Troubleshooting for alert RabbitmqTooManyUnackMessages
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyUnackMessages

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many unacknowledged messages

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqTooManyUnackMessages" %}}

<!-- Rule when generated

```yaml
alert: RabbitmqTooManyUnackMessages
expr: sum(rabbitmq_queue_messages_unacked) BY (queue) > 1000
for: 1m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many unack messages (instance {{ $labels.instance }})
    description: |-
        Too many unacknowledged messages
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqTooManyUnackMessages.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
