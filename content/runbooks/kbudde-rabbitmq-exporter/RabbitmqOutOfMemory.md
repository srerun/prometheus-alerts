---
title: RabbitmqOutOfMemory
description: Troubleshooting for alert RabbitmqOutOfMemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqOutOfMemory

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Memory available for RabbmitMQ is low (< 10%)

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqOutOfMemory" %}}

<!-- Rule when generated

```yaml
alert: RabbitmqOutOfMemory
expr: rabbitmq_node_mem_used / rabbitmq_node_mem_limit * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ out of memory (instance {{ $labels.instance }})
    description: |-
        Memory available for RabbmitMQ is low (< 10%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqOutOfMemory.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
