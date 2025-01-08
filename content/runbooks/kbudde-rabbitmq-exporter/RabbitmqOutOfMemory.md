---
title: RabbitmqOutOfMemory
description: Troubleshooting for alert RabbitmqOutOfMemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqOutOfMemory

Memory available for RabbmitMQ is low (< 10%)

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqOutOfMemory" %}}

{{% comment %}}

```yaml
alert: RabbitmqOutOfMemory
expr: rabbitmq_node_mem_used / rabbitmq_node_mem_limit * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ out of memory (instance {{ $labels.instance }})
    description: |-
        Memory available for RabbmitMQ is low (< 10%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqOutOfMemory.md

```

{{% /comment %}}

</details>


Here is the runbook for the RabbitMQ Out of Memory alert:

## Meaning

The RabbitMQ Out of Memory alert is triggered when the memory usage of a RabbitMQ node exceeds 90% of its available memory. This indicates that the RabbitMQ instance is at risk of running out of memory, which can lead to performance issues, crashes, and even data loss.

## Impact

If left unaddressed, high memory usage can cause:

* RabbitMQ performance degradation
* Increased likelihood of crashes and restarts
* Message loss and duplication
* Queues and exchanges becoming unavailable
* Impact on dependent applications and services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ web interface or the Prometheus dashboard to identify the node(s) experiencing high memory usage.
2. Investigate recent changes to the RabbitMQ configuration, such as queue or exchange settings, that may be contributing to the increased memory usage.
3. Verify that the RabbitMQ instance has sufficient available memory and that the system is not experiencing resource constraints.
4. Review the RabbitMQ logs for errors or warnings related to memory usage.
5. Check for any abnormal message throughput or queue growth.

## Mitigation

To mitigate the issue, follow these steps:

1. **Reduce message throughput**: Temporarily reduce the number of messages being published to RabbitMQ to alleviate memory pressure.
2. **Increase available memory**: Consider increasing the available memory for the RabbitMQ instance, either by adjusting the node's resource allocation or by adding more nodes to the cluster.
3. **Optimize RabbitMQ configuration**: Review and optimize RabbitMQ configuration settings, such as queue and exchange settings, to reduce memory usage.
4. **Purge unnecessary data**: Remove unnecessary messages, queues, or exchanges to free up memory.
5. **Restart RabbitMQ**: If necessary, restart the RabbitMQ instance to clear out any cached data and reset memory usage.

Remember to monitor the situation closely and adjust the mitigation steps as needed to prevent further memory-related issues.