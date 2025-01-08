---
title: KafkaTopicsReplicas
description: Troubleshooting for alert KafkaTopicsReplicas
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KafkaTopicsReplicas

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Kafka topic in-sync partition

<details>
  <summary>Alert Rule</summary>

{{% rule "kafka/danielqsj-kafka-exporter.yml" "KafkaTopicsReplicas" %}}

{{% comment %}}

```yaml
alert: KafkaTopicsReplicas
expr: sum(kafka_topic_partition_in_sync_replica) by (topic) < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Kafka topics replicas (instance {{ $labels.instance }})
    description: |-
        Kafka topic in-sync partition
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/danielqsj-kafka-exporter/KafkaTopicsReplicas.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
