---
title: RabbitmqMemoryHigh
description: Troubleshooting for alert RabbitmqMemoryHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqMemoryHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A node use more than 90% of allocated RAM

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqMemoryHigh" %}}

{{% comment %}}

```yaml
alert: RabbitmqMemoryHigh
expr: rabbitmq_process_resident_memory_bytes / rabbitmq_resident_memory_limit_bytes * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ memory high (instance {{ $labels.instance }})
    description: |-
        A node use more than 90% of allocated RAM
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqMemoryHigh.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
