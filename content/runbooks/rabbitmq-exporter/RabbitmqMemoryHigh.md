---
title: RabbitmqMemoryHigh
description: Troubleshooting for alert RabbitmqMemoryHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqMemoryHigh

A node use more than 90% of allocated RAM

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqMemoryHigh" %}}

{{% comment %}}

```yaml
alert: RabbitmqMemoryHigh
expr: rabbitmq_process_resident_memory_bytes / rabbitmq_resident_memory_limit_bytes * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ memory high (instance {{ $labels.instance }})
    description: |-
        A node use more than 90% of allocated RAM
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqMemoryHigh.md

```

{{% /comment %}}

</details>


## Meaning

The RabbitmqMemoryHigh alert is triggered when the resident memory usage of a RabbitMQ node exceeds 90% of the allocated RAM. This alert indicates that the RabbitMQ instance is consuming a large amount of memory, which can lead to performance issues, slow message processing, and even node crashes.

## Impact

* Performance degradation: High memory usage can slow down message processing, leading to increased latency and decreased throughput.
* Node instability: Prolonged high memory usage can cause the RabbitMQ node to crash, resulting in message loss and downtime.
* Cascading failures: If multiple nodes experience high memory usage, it can lead to a cascading failure of the entire RabbitMQ cluster.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ node's memory usage: Verify the current memory usage using the Prometheus metric `rabbitmq_process_resident_memory_bytes`.
2. Identify the root cause: Investigate the cause of high memory usage, such as:
	* Increased message volume
	* Large message sizes
	* Inefficient message processing
	* Resource-intensive plugins or extensions
3. Review RabbitMQ configuration: Check the RabbitMQ configuration to ensure that it is optimized for the current workload.
4. Monitor node performance: Observe the node's performance metrics, such as CPU usage, to determine if there are any underlying performance issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Reduce message volume: Implement measures to reduce the message volume, such as:
	* Load balancing
	* Message filtering
	* Queue length reduction
2. Optimize message processing: Improve message processing efficiency by:
	* Optimizing plugin configurations
	* Implementing message batching
	* Reducing message sizes
3. Increase allocated RAM: Consider increasing the allocated RAM for the RabbitMQ node to ensure it has sufficient resources.
4. Restart the RabbitMQ node: If the issue persists, restart the RabbitMQ node to clear out any temporary memory allocations.
5. Implement permanent fixes: Implement permanent fixes, such as optimization of RabbitMQ configuration, plugin updates, or architecture changes, to prevent similar issues in the future.