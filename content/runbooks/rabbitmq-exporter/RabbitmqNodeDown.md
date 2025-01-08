---
title: RabbitmqNodeDown
description: Troubleshooting for alert RabbitmqNodeDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqNodeDown

Less than 3 nodes running in RabbitMQ cluster

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqNodeDown" %}}

{{% comment %}}

```yaml
alert: RabbitmqNodeDown
expr: sum(rabbitmq_build_info) < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ node down (instance {{ $labels.instance }})
    description: |-
        Less than 3 nodes running in RabbitMQ cluster
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqNodeDown.md

```

{{% /comment %}}

</details>


Here is a runbook for the RabbitmqNodeDown alert rule:

## Meaning

The RabbitmqNodeDown alert is triggered when the number of nodes running in the RabbitMQ cluster falls below 3. This means that one or more nodes in the cluster are not available, which can lead to reduced capacity, increased latency, and potential data loss.

## Impact

The impact of this alert is critical, as it can cause:

* Reduced message processing capacity, leading to increased latency and potential message loss
* Increased risk of data loss and inconsistencies between nodes
* Potential business impact due to reduced system availability and reliability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the RabbitMQ cluster status using the RabbitMQ Management UI or the `rabbitmqctl` command-line tool
2. Verify the number of nodes running in the cluster and identify which nodes are down
3. Check the system logs for errors or warnings related to the down nodes
4. Verify that the RabbitMQ exporter is correctly configured and running

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the down nodes or replace them if they are failed
2. Verify that the nodes are properly configured and connected to the cluster
3. Check for any software or hardware issues that may be preventing the nodes from running
4. Consider adding more nodes to the cluster to increase capacity and redundancy
5. Verify that the RabbitMQ exporter is correctly configured and running to ensure accurate monitoring and alerting.