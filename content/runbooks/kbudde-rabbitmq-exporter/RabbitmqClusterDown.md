---
title: RabbitmqClusterDown
description: Troubleshooting for alert RabbitmqClusterDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqClusterDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Less than 3 nodes running in RabbitMQ cluster

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqClusterDown" %}}

<!-- Rule when generated

```yaml
alert: RabbitmqClusterDown
expr: sum(rabbitmq_running) < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ cluster down (instance {{ $labels.instance }})
    description: |-
        Less than 3 nodes running in RabbitMQ cluster
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqClusterDown.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
