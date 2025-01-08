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

Kafka consumer has a 30 minutes and increasing lag

<details>
  <summary>Alert Rule</summary>

{{% rule "kafka/linkedin-kafka-exporter.yml" "KafkaConsumerLag" %}}

{{% comment %}}

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

{{% /comment %}}

</details>


Here is a sample runbook for the KafkaConsumerLag alert:

## Meaning

The KafkaConsumerLag alert is triggered when a Kafka consumer's lag (the difference between the latest offset in a partition and the consumer's current offset) has been increasing for at least 30 minutes and is currently greater than 0. This means that the consumer is not keeping up with the production of new messages in the partition, which can lead to data loss, increased latency, and other issues.

## Impact

* Data loss: If the consumer lag continues to increase, it may lead to missed messages, which can result in data loss or inconsistencies.
* Increased latency: As the consumer lags behind, it may take longer for messages to be processed, leading to increased latency and slower system response times.
* System performance: A lagging consumer can put additional pressure on the Kafka cluster, leading to decreased performance and potential issues with other consumers or producers.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Kafka cluster's overall health and performance using metrics such as Kafka's request latency, throughput, and error rates.
2. Verify that the consumer is properly configured and running correctly.
3. Check the consumer's logs for any errors or warnings that may indicate the cause of the lag.
4. Use tools like `kafka-consumer-groups` or `kafka-console-consumer` to inspect the consumer group's offset and lag.
5. Investigate any recent changes to the Kafka cluster, consumer configuration, or application code that may be contributing to the lag.

## Mitigation

To mitigate the KafkaConsumerLag issue, follow these steps:

1. **Increase consumer instances or partitions**: If the consumer is under-provisioned, adding more instances or partitions can help it catch up with the incoming messages.
2. **Optimize consumer configuration**: Review the consumer configuration and optimize settings such as `fetch.min.bytes`, `fetch.max.bytes`, and `max.partition.fetch.bytes` to improve its performance.
3. **Investigate and resolve underlying issues**: Address any underlying issues that may be causing the lag, such as network connectivity problems, Kafka broker issues, or application code bugs.
4. **Consider rebalancing the consumer group**: If the lag is due to an uneven distribution of partitions among consumers, consider rebalancing the consumer group to redistribute the load.
5. **Monitor the consumer's progress**: Continue to monitor the consumer's lag and adjust the mitigation strategy as needed to ensure the consumer catches up with the incoming messages.