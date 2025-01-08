---
title: RabbitmqNodeNotDistributed
description: Troubleshooting for alert RabbitmqNodeNotDistributed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqNodeNotDistributed

Distribution link state is not 'up'

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqNodeNotDistributed" %}}

{{% comment %}}

```yaml
alert: RabbitmqNodeNotDistributed
expr: erlang_vm_dist_node_state < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ node not distributed (instance {{ $labels.instance }})
    description: |-
        Distribution link state is not 'up'
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqNodeNotDistributed.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the RabbitmqNodeNotDistributed alert:

## Meaning

The RabbitmqNodeNotDistributed alert is triggered when a RabbitMQ node's distribution link state is not in the "up" state. This means that the node is not properly connected to other nodes in the cluster, which can lead to inconsistencies in message queues, loss of data, and other issues.

## Impact

If left unaddressed, this issue can cause significant disruptions to message processing and queuing, leading to:

* Loss of data or messages
* Inconsistent queue states across nodes
* Increased latency or timeouts in message processing
* Potential cascade failures in dependent systems

## Diagnosis

To diagnose this issue, follow these steps:

1. Check the RabbitMQ node's logs for errors or warnings related to distribution link state.
2. Use the RabbitMQ management UI or command-line tools to check the node's current distribution link state.
3. Verify that the node is properly configured to join the cluster and that there are no network connectivity issues between nodes.
4. Check for any recent changes or updates to the RabbitMQ configuration or network infrastructure.

## Mitigation

To mitigate this issue, follow these steps:

1. Check the RabbitMQ node's configuration to ensure it is properly configured to join the cluster.
2. Restart the RabbitMQ node to re-establish the distribution link.
3. Verify that the distribution link state is "up" after the restart.
4. If the issue persists, check the network infrastructure and ensure that there are no connectivity issues between nodes.
5. Consider increasing the logging verbosity or enabling debug logging to gather more information about the issue.
6. If the issue cannot be resolved, consider reaching out to RabbitMQ support or consulting the RabbitMQ documentation for further guidance.