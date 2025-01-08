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


Here is a runbook for the KafkaTopicOffsetDecreased alert:

## Meaning

The KafkaTopicOffsetDecreased alert is triggered when the current offset of a Kafka topic partition decreases. This can indicate a problem with the Kafka consumers or the Kafka broker itself. A decreasing offset can lead to data loss or duplication, depending on the configuration of the consumers.

## Impact

The impact of a decreasing Kafka topic offset can be significant, leading to:

* Data loss or duplication
* Inconsistent data processing
* Consumer lag or failure
* Potential disruption to downstream applications

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Kafka console consumer or Kafka Tool to verify the current offset and verify that it is decreasing.
2. Investigate the Kafka consumer logs to identify any errors or issues that may be causing the offset to decrease.
3. Check the Kafka broker logs to identify any issues with the broker that may be causing the offset to decrease.
4. Verify that the Kafka consumer configuration is correct and that the consumer is properly configured to commit offsets.
5. Check for any network connectivity issues between the Kafka consumer and the Kafka broker.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and resolve the root cause of the offset decrease, such as fixing consumer configuration issues or resolving network connectivity problems.
2. Use the Kafka console consumer or Kafka Tool to manually adjust the offset to the correct position.
3. If the issue is caused by a consumer failure, restart the consumer and verify that it is properly consuming from the correct offset.
4. Consider increasing the `offsets.retention.minutes` configuration in the Kafka broker to ensure that offsets are retained for a longer period, allowing for easier recovery in case of an offset decrease.
5. Consider implementing offset monitoring and alerting to quickly detect and respond to offset decreases.