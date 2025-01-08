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


Here is a runbook for the Prometheus alert rule:

## Meaning

This alert is triggered when the lag of a Kafka consumers group exceeds 50 messages for more than 1 minute. The lag is calculated as the sum of the kafka_consumergroup_lag metric by consumergroup. This indicates that the consumers in the group are not keeping up with the producers, which can lead to message loss or latency.

## Impact

If left unchecked, this issue can lead to:

* Message loss or duplication
* Increased latency in processing messages
* Unprocessable messages accumulating in the Kafka topic
* Downstream systems experiencing errors or delays

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Kafka consumers group metrics to identify the specific consumergroup experiencing the lag.
2. Investigate the consumer instance(s) in the group to determine if they are experiencing high CPU usage, memory issues, or network connectivity problems.
3. Verify that the consumer configuration is correct, including topics, partitions, and offset management.
4. Review the Kafka broker logs for any errors or issues that may be contributing to the lag.

## Mitigation

To mitigate the issue, follow these steps:

1. Check and adjust the consumer instance resources (e.g., increase CPU or memory) to ensure they can keep up with the message volume.
2. Verify that the consumer configuration is optimal, including tuning parameters such as batch size, fetch size, and maxlag.
3. Implement backpressure mechanisms, such as pause and resume, to control the message flow and prevent overload.
4. Consider rebalancing the Kafka topic partitions to distribute the load more evenly across consumers.
5. If necessary, adjust the alert threshold or time window to better suit the specific use case and environment.

Remember to monitor the situation and adjust the mitigation strategies as needed to prevent further occurrences of this alert.