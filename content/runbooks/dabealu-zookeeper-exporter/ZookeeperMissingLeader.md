---
title: ZookeeperMissingLeader
description: Troubleshooting for alert ZookeeperMissingLeader
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZookeeperMissingLeader

Zookeeper cluster has no node marked as leader

<details>
  <summary>Alert Rule</summary>

{{% rule "zookeeper/dabealu-zookeeper-exporter.yml" "ZookeeperMissingLeader" %}}

{{% comment %}}

```yaml
alert: ZookeeperMissingLeader
expr: sum(zk_server_leader) == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Zookeeper missing leader (instance {{ $labels.instance }})
    description: |-
        Zookeeper cluster has no node marked as leader
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dabealu-zookeeper-exporter/ZookeeperMissingLeader.md

```

{{% /comment %}}

</details>


## Meaning

The ZookeeperMissingLeader alert is triggered when Prometheus detects that there is no node marked as the leader in the Zookeeper cluster. This is determined by the `zk_server_leader` metric, which should have a non-zero value indicating the presence of a leader node. If the sum of this metric is zero, it means that no node is currently acting as the leader, which can have severe consequences for the cluster's availability and integrity.

## Impact

The absence of a Zookeeper leader node can lead to:

* Cluster instability and potential data loss
* Inability to perform writes or updates to the Zookeeper database
* Failure of dependent systems and applications that rely on Zookeeper for coordination and configuration management
* Increased latency and errors in distributed systems that use Zookeeper

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Zookeeper cluster status using the Zookeeper CLI or a monitoring tool like Apache Zookeeper UI.
2. Verify that all Zookeeper nodes are running and reachable.
3. Check the Zookeeper logs for any error messages related to leader election or node connectivity.
4. Use Prometheus and Grafana to visualize the `zk_server_leader` metric and identify any trends or patterns that may indicate the cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the cause of the leader node failure and resolve it (e.g., restart the node, fix network connectivity issues, etc.).
2. If the leader node is down, promote another node to become the leader using the Zookeeper CLI or API.
3. Verify that the `zk_server_leader` metric returns a non-zero value, indicating the presence of a leader node.
4. Monitor the Zookeeper cluster for any further issues and take corrective action if necessary.
5. Consider implementing measures to prevent future leader node failures, such as deploying a highly available Zookeeper cluster or implementing automated leader election and failover.