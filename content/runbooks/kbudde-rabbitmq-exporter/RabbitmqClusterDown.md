---
title: RabbitmqClusterDown
description: Troubleshooting for alert RabbitmqClusterDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqClusterDown

Less than 3 nodes running in RabbitMQ cluster

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/kbudde-rabbitmq-exporter.yml" "RabbitmqClusterDown" %}}

{{% comment %}}

```yaml
alert: RabbitmqClusterDown
expr: sum(rabbitmq_running) < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: RabbitMQ cluster down (instance {{ $labels.instance }})
    description: |-
        Less than 3 nodes running in RabbitMQ cluster
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kbudde-rabbitmq-exporter/RabbitmqClusterDown.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the RabbitmqClusterDown alert rule:

## Meaning

The RabbitmqClusterDown alert is triggered when less than 3 nodes are running in the RabbitMQ cluster, indicating that the cluster is down or not fully functional. This alert is critical as it can cause disruptions to message processing and potentially lead to data loss or inconsistencies.

## Impact

The impact of a RabbitMQ cluster being down can be significant, leading to:

* Message loss or delay
* Inconsistent data state
* Disruptions to dependent applications and services
* Potential data corruption or inconsistencies
* Increased latency or timeouts

## Diagnosis

To diagnose the root cause of the RabbitMQ cluster being down, follow these steps:

1. Check the RabbitMQ node status: Verify the status of each node in the cluster to identify which nodes are not running.
2. Review system logs: Analyze system logs to identify any errors or issues that may be causing nodes to fail or not start.
3. Check network connectivity: Verify that network connectivity between nodes is stable and healthy.
4. Verify configuration: Check the RabbitMQ configuration to ensure it is correct and consistent across all nodes.
5. Run RabbitMQ diagnostics: Run RabbitMQ built-in diagnostics to identify any issues with the cluster.

## Mitigation

To mitigate the RabbitMQ cluster being down, follow these steps:

1. Restart failed nodes: Attempt to restart any failed nodes to bring them back online.
2. Investigate and resolve root cause: Identify and resolve the root cause of the node failure to prevent recurrence.
3. Check disk space and resources: Verify that nodes have sufficient disk space and resources to operate correctly.
4. Implement automated restarts: Configure automated restarts for RabbitMQ nodes to minimize downtime.
5. Consider failover scenarios: Evaluate failover scenarios to ensure that the cluster can recover from node failures.

Remember to update the runbook with specific instructions and procedures relevant to your environment and team.