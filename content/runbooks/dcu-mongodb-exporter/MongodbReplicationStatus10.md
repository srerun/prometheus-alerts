---
title: MongodbReplicationStatus10
description: Troubleshooting for alert MongodbReplicationStatus10
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus10

MongoDB Replication set member was once in a replica set but was subsequently removed

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationStatus10" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationStatus10
expr: mongodb_replset_member_state == 10
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 10 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member was once in a replica set but was subsequently removed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationStatus10.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `MongodbReplicationStatus10`:

## Meaning

This alert is triggered when a MongoDB replica set member has a state of 10, indicating that it was once part of a replica set but has been subsequently removed. This can occur due to various reasons such as node failure, network issues, or intentional removal of the node from the replica set.

## Impact

A MongoDB replica set member with a state of 10 can lead to data inconsistencies and availability issues. Data written to the removed node may not be replicated to other nodes, resulting in data loss or inconsistencies. Additionally, the removal of a node can affect the overall performance and redundancy of the replica set.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the MongoDB replica set configuration to ensure that the removed node was intentionally removed or if it was due to a failure.
2. Review the MongoDB logs to identify the reason for the node's removal from the replica set.
3. Verify the replication status of the remaining nodes in the replica set to ensure they are in a healthy state.
4. Check the network connectivity between the nodes to ensure there are no issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Restore the removed node to the replica set, if possible, or add a new node to replace it.
2. Resync the data on the restored or new node with the rest of the replica set.
3. Verify the replication status of the entire replica set to ensure all nodes are in a healthy state.
4. Implement measures to prevent similar issues in the future, such as node monitoring, automatic failover, and data backups.

Note: This runbook provides general guidance and may need to be customized based on the specific MongoDB deployment and environment.