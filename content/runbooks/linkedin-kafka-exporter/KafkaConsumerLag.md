---
title: KafkaConsumerLag
description: Troubleshooting for alert KafkaConsumerLag
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KafkaConsumerLag

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Kafka consumer has a 30 minutes and increasing lag

<details>
  <summary>Alert Rule</summary>

{{% rule "kafka/linkedin-kafka-exporter.yml" "KafkaConsumerLag" %}}

<!-- Rule when generated

```yaml
alert: KafkaConsumerLag
expr: kafka_burrow_topic_partition_offset - on(partition, cluster, topic) group_right() kafka_burrow_partition_current_offset >= (kafka_burrow_topic_partition_offset offset 15m - on(partition, cluster, topic) group_right() kafka_burrow_partition_current_offset offset 15m) AND kafka_burrow_topic_partition_offset - on(partition, cluster, topic) group_right() kafka_burrow_partition_current_offset > 0
for: 15m
labels:
    severity: warning
annotations:
    summary: Kafka consumer lag (instance {{ $labels.instance }})
    description: |-
        Kafka consumer has a 30 minutes and increasing lag
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/linkedin-kafka-exporter/KafkaConsumerLag.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
