---
title: EtcdInsufficientMembers
description: Troubleshooting for alert EtcdInsufficientMembers
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# EtcdInsufficientMembers

Etcd cluster should have an odd number of members

<details>
  <summary>Alert Rule</summary>

{{% rule "etcd/etcd-internal.yml" "EtcdInsufficientMembers" %}}

{{% comment %}}

```yaml
alert: EtcdInsufficientMembers
expr: count(etcd_server_id) % 2 == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Etcd insufficient Members (instance {{ $labels.instance }})
    description: |-
        Etcd cluster should have an odd number of members
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/etcd-internal/EtcdInsufficientMembers.md

```

{{% /comment %}}

</details>


Here is a runbook for the EtcdInsufficientMembers alert:

## Meaning

The EtcdInsufficientMembers alert is triggered when the number of etcd members in the cluster is even. Etcd, a distributed key-value store, requires an odd number of members to maintain a quorum and ensure the cluster's availability. With an even number of members, the cluster is at risk of becoming unavailable or inconsistent in case of a node failure.

## Impact

If the EtcdInsufficientMembers alert is not addressed, it can lead to:

* Cluster unavailability: With an even number of members, etcd may not be able to achieve a quorum, leading to cluster downtime and data inconsistency.
* Data loss: In the event of a node failure, the cluster may not be able to recover, resulting in data loss or corruption.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the etcd cluster membership: Run the command `etcdctl member list` to verify the current number of members in the cluster.
2. Identify the missing member: Check the etcd server logs to determine which member is missing or not participating in the cluster.
3. Verify etcd configuration: Review the etcd configuration files to ensure that the member count is correctly set and matches the expected number of members.

## Mitigation

To mitigate the issue, follow these steps:

1. Add or replace the missing member: Bring up a new etcd member or replace the failed node to restore the odd number of members in the cluster.
2. Verify etcd cluster health: Run the command `etcdctl cluster` to verify that the cluster is healthy and all members are participating.
3. Monitor etcd metrics: Keep a close eye on etcd metrics, such as `etcd_server_id`, to ensure that the cluster remains healthy and available.

Remember to update the etcd configuration files and deployment scripts to prevent similar issues in the future.