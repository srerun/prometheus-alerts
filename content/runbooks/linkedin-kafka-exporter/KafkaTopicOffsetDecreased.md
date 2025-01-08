---
title: KafkaTopicOffsetDecreased
description: Troubleshooting for alert KafkaTopicOffsetDecreased
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KafkaTopicOffsetDecreased

Kafka topic offset has decreased

<details>
  <summary>Alert Rule</summary>

{{% rule "kafka/linkedin-kafka-exporter.yml" "KafkaTopicOffsetDecreased" %}}

{{% comment %}}

```yaml
alert: KafkaTopicOffsetDecreased
expr: delta(kafka_burrow_partition_current_offset[1m]) < 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Kafka topic offset decreased (instance {{ $labels.instance }})
    description: |-
        Kafka topic offset has decreased
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/linkedin-kafka-exporter/KafkaTopicOffsetDecreased.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
