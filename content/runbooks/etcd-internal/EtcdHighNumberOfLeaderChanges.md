---
title: EtcdHighNumberOfLeaderChanges
description: Troubleshooting for alert EtcdHighNumberOfLeaderChanges
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdHighNumberOfLeaderChanges

Etcd leader changed more than 2 times during 10 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdHighNumberOfLeaderChanges" %}}

{{% comment %}}

```yaml
alert: EtcdHighNumberOfLeaderChanges
expr: increase(etcd_server_leader_changes_seen_total[10m]) > 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Etcd high number of leader changes (instance {{ $labels.instance }})
    description: |-
        Etcd leader changed more than 2 times during 10 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdHighNumberOfLeaderChanges.md

```

{{% /comment %}}

</details>


Here is the runbook for the `EtcdHighNumberOfLeaderChanges` alert:

## Meaning

The `EtcdHighNumberOfLeaderChanges` alert is triggered when the number of etcd leader changes seen in the last 10 minutes exceeds 2. This alert indicates that the etcd cluster is experiencing instability, which can lead to issues with cluster availability and data consistency.

## Impact

A high number of etcd leader changes can have several negative impacts on the cluster:

* **Cluster instability**: Frequent leader changes can cause the cluster to become unstable, leading to issues with node communication and data replication.
* **Data inconsistencies**: With multiple leaders, there is a risk of data inconsistencies between nodes, which can lead to errors and data loss.
* **Performance degradation**: The constant switching of leaders can cause performance issues, such as slow query responses and increased latency.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check etcd cluster health**: Verify the overall health of the etcd cluster using the `etcdctl` command-line tool or the etcd dashboard.
2. **Investigate node logs**: Review the logs of the etcd nodes to identify any errors or issues that may be contributing to the leader changes.
3. **Verify network connectivity**: Ensure that the network connectivity between etcd nodes is stable and not experiencing any issues.
4. **Check for resource constraints**: Verify that the etcd nodes have sufficient resources (e.g., CPU, memory, disk space) to operate effectively.

## Mitigation

To mitigate the issue, follow these steps:

1. **Identify and address underlying causes**: Address any underlying issues identified during diagnosis, such as node failures, network connectivity issues, or resource constraints.
2. **Adjust etcd configuration**: Consider adjusting etcd configuration parameters, such as the election timeout or heartbeat interval, to improve cluster stability.
3. **Implement etcd cluster monitoring**: Implement monitoring tools to track etcd cluster health and performance, and set up alerts for anomalies.
4. **Plan for etcd maintenance**: Schedule regular maintenance for etcd nodes, such as backups and upgrades, to ensure the cluster remains healthy and stable.