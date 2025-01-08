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
The total connections of a node is too high

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqTooManyConnections" %}}

{{% comment %}}

```yaml
alert: RabbitmqTooManyConnections
expr: rabbitmq_connections > 1000
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many connections (instance {{ $labels.instance }})
    description: |-
        The total connections of a node is too high
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqTooManyConnections.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
