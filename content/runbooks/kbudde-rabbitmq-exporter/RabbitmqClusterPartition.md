---
title: RabbitmqClusterPartition
description: Troubleshooting for alert RabbitmqClusterPartition
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqClusterPartition

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cluster partition

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqClusterPartition" %}}

<!-- Rule when generated

```yaml
alert: RabbitmqClusterPartition
expr: rabbitmq_partitions > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ cluster partition (instance {{ $labels.instance }})
    description: |-
        Cluster partition
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqClusterPartition.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
