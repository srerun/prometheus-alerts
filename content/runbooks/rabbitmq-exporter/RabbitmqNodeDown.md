---
title: RabbitmqNodeDown
description: Troubleshooting for alert RabbitmqNodeDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqNodeDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Less than 3 nodes running in RabbitMQ cluster

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqNodeDown" %}}

{{% comment %}}

```yaml
alert: RabbitmqNodeDown
expr: sum(rabbitmq_build_info) < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ node down (instance {{ $labels.instance }})
    description: |-
        Less than 3 nodes running in RabbitMQ cluster
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqNodeDown.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
