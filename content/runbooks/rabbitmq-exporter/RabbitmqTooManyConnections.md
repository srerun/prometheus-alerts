---
title: RabbitmqTooManyConnections
description: Troubleshooting for alert RabbitmqTooManyConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqTooManyConnections

The total connections of a node is too high

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqTooManyConnections" %}}

{{% comment %}}

```yaml
alert: RabbitmqTooManyConnections
expr: rabbitmq_connections > 1000
for: 2m
labels:
    severity: warning
annotations:
    summary: RabbitMQ too many connections (instance {{ $labels.instance }})
    description: |-
        The total connections of a node is too high
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqTooManyConnections.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule "RabbitmqTooManyConnections":

## Meaning

This alert is triggered when the number of connections to a RabbitMQ node exceeds 1000. This can indicate a potential performance issue or resource overload on the node.

## Impact

A high number of connections to a RabbitMQ node can lead to:

* Decreased performance and throughput
* Increased memory usage and resource consumption
* Potential node crashes or instability
* Impacted message processing and delivery times

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

1. Check the RabbitMQ node's connection statistics to identify the source of the high connection count.
2. Review the node's configuration to ensure it is properly sized and configured for the current workload.
3. Verify that there are no issues with the application(s) producing messages to the RabbitMQ node.
4. Check for any network connectivity issues or firewall rules that may be preventing connections from being closed.
5. Review RabbitMQ logs for any error messages related to connection handling.

## Mitigation

To mitigate this issue, follow these steps:

1. Reduce the number of connections to the RabbitMQ node by:
	* Adjusting application configuration to reduce message production rates
	* Implementing connection pooling or queuing mechanisms
	* Increasing the node's resources (e.g., RAM, CPU)
2. Optimize RabbitMQ node configuration for better performance and resource utilization.
3. Monitor RabbitMQ node statistics and adjust as necessary to prevent future connection count issues.
4. Consider implementing RabbitMQ clustering or load balancing to distribute connections across multiple nodes.
5. Verify that firewall rules are configured to allow connections to be closed properly.