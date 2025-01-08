---
title: ZookeeperTooManyLeaders
description: Troubleshooting for alert ZookeeperTooManyLeaders
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZookeeperTooManyLeaders

Zookeeper cluster has too many nodes marked as leader

<details>
  <summary>Alert Rule</summary>

{{% rule "zookeeper/dabealu-zookeeper-exporter.yml" "ZookeeperTooManyLeaders" %}}

{{% comment %}}

```yaml
alert: ZookeeperTooManyLeaders
expr: sum(zk_server_leader) > 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Zookeeper Too Many Leaders (instance {{ $labels.instance }})
    description: |-
        Zookeeper cluster has too many nodes marked as leader
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dabealu-zookeeper-exporter/ZookeeperTooManyLeaders.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule "ZookeeperTooManyLeaders":

## Meaning

This alert indicates that the Zookeeper cluster has more than one node marked as the leader. In a healthy Zookeeper cluster, there should only be one leader node. Having multiple leader nodes can cause inconsistencies and errors in the cluster.

## Impact

Having multiple leader nodes in a Zookeeper cluster can lead to:

* Inconsistent data across the cluster
* Errors in client applications that rely on Zookeeper
* Unstable cluster behavior
* Potential data loss or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Zookeeper cluster status using the Zookeeper command-line tool or a monitoring tool like Zookeeper Exporter.
2. Verify that the cluster has more than one node marked as the leader.
3. Check the Zookeeper server logs for any errors or warnings related to leadership elections or node failures.
4. Check the network connectivity between the Zookeeper nodes to ensure that there are no issues with communication.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the cause of the multiple leader nodes, such as a node failure or network partition.
2. If a node failure is suspected, try to restart the failed node or replace it with a new one.
3. If a network partition is suspected, try to resolve the network connectivity issue.
4. If the issue persists, consider restarting the entire Zookeeper cluster to ensure a clean leadership election.
5. Monitor the cluster closely to ensure that the issue is resolved and the cluster is stable.

Remember to always exercise caution when making changes to a Zookeeper cluster, as it can affect the stability of the entire system.