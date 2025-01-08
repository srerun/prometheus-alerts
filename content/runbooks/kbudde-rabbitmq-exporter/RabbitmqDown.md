---
title: RabbitmqDown
description: Troubleshooting for alert RabbitmqDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
RabbitMQ node down

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqDown" %}}

{{% comment %}}

```yaml
alert: RabbitmqDown
expr: rabbitmq_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ down (instance {{ $labels.instance }})
    description: |-
        RabbitMQ node down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqDown.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
