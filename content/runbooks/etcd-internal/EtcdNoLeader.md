---
title: EtcdNoLeader
description: Troubleshooting for alert EtcdNoLeader
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdNoLeader

Etcd cluster have no leader

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdNoLeader" %}}

{{% comment %}}

```yaml
alert: EtcdNoLeader
expr: etcd_server_has_leader == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Etcd no Leader (instance {{ $labels.instance }})
    description: |-
        Etcd cluster have no leader
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdNoLeader.md

```

{{% /comment %}}

</details>


Here is the runbook for the EtcdNoLeader alert:

## Meaning

The EtcdNoLeader alert is triggered when the etcd cluster does not have a leader node. etcd is a distributed key-value store that is used to manage and store data in a distributed system. The leader node is responsible for managing the cluster and ensuring that data is replicated correctly. If the cluster does not have a leader, it can lead to data inconsistencies and errors.

## Impact

The impact of this alert is critical, as it can cause the following issues:

* Data inconsistencies: Without a leader, the etcd cluster may not be able to replicate data correctly, leading to inconsistencies across the nodes.
* Cluster instability: The lack of a leader can cause the cluster to become unstable, leading to errors and failures.
* Service disruptions: Depending on the services that rely on etcd, this alert can cause service disruptions and errors.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the etcd cluster status using the etcdctl command-line tool.
2. Verify that all etcd nodes are running and healthy.
3. Check the etcd logs for any error messages related to leader election or cluster instability.
4. Verify that the network connectivity between the etcd nodes is working correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the etcd configuration and ensure that it is correctly configured for leader election.
2. Verify that all etcd nodes are running with the correct configuration and that there are no network connectivity issues.
3. If the issue persists, try to manually elect a new leader using the etcdctl tool.
4. If the issue still persists, consider restarting the etcd nodes or seeking assistance from a qualified administrator.

Remember to refer to the etcd documentation and the official runbook for more detailed information on troubleshooting and mitigating this issue.