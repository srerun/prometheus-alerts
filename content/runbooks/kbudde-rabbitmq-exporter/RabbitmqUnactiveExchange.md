---
title: RabbitmqUnactiveExchange
description: Troubleshooting for alert RabbitmqUnactiveExchange
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqUnactiveExchange

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Exchange receive less than 5 msgs per second

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/RabbitmqUnactiveExchange

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
