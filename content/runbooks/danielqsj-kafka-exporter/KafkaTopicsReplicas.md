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


Here is the runbook for the KafkaTopicsReplicas alert:

## Meaning

The KafkaTopicsReplicas alert is triggered when the number of in-sync replicas for a Kafka topic is less than 3. This means that the topic is not properly replicated, which can lead to data loss or inconsistencies in the event of a failure.

## Impact

The impact of this alert is high, as it can result in:

* Data loss or inconsistencies in the event of a failure
* Reduced reliability and availability of the Kafka cluster
* Potential disruption to applications that rely on the affected topic

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the KafkaTopicPartitionInSyncReplica metric in Prometheus to identify the affected topic and the number of in-sync replicas.
2. Verify that the Kafka broker configuration is correct and that the expected number of replicas is set.
3. Check the Kafka broker logs for any errors or issues related to replication.
4. Verify that the network connectivity between the Kafka brokers is stable and that there are no issues with the underlying infrastructure.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the affected topic and the underlying reason for the lack of in-sync replicas.
2. Increase the number of replicas for the affected topic by updating the Kafka topic configuration.
3. Verify that the replication factor is set correctly and that the Kafka brokers are properly configured.
4. Monitor the KafkaTopicPartitionInSyncReplica metric to ensure that the issue is resolved and the topic is properly replicated.
5. Consider implementing automated topic creation and replication management to prevent similar issues in the future.

Note: This runbook assumes that the Kafka cluster is managed by the operations team and that they have the necessary knowledge and access to perform the mitigation steps. If this is not the case, additional guidance and support may be necessary.