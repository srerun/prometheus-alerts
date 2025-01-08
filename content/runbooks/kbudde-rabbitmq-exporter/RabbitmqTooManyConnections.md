---
title: RabbitmqTooManyConnections
description: Troubleshooting for alert RabbitmqTooManyConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
RabbitMQ instance has too many connections (> 1000)

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqTooManyConnections" %}}

<!-- Rule when generated

```yaml
alert: RabbitmqTooManyConnections
expr: rabbitmq_connectionsTotal > 1000
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many connections (instance {{ $labels.instance }})
    description: |-
        RabbitMQ instance has too many connections (> 1000)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqTooManyConnections.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
