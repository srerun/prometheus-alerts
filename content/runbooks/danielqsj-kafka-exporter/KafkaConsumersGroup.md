---
title: KafkaConsumersGroup
description: Troubleshooting for alert KafkaConsumersGroup
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KafkaConsumersGroup

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Kafka consumers group

<details>
  <summary>Alert Rule</summary>

{{% rule "kafka/danielqsj-kafka-exporter.yml" "KafkaConsumersGroup" %}}

{{% comment %}}

```yaml
alert: KafkaConsumersGroup
expr: sum(kafka_consumergroup_lag) by (consumergroup) > 50
for: 1m
labels:
    severity: critical
annotations:
    summary: Kafka consumers group (instance {{ $labels.instance }})
    description: |-
        Kafka consumers group
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/danielqsj-kafka-exporter/KafkaConsumersGroup.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
