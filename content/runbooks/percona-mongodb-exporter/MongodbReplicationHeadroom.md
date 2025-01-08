---
title: MongodbReplicationHeadroom
description: Troubleshooting for alert MongodbReplicationHeadroom
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationHeadroom

MongoDB replication headroom is <= 0

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbReplicationHeadroom" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationHeadroom
expr: sum(avg(mongodb_mongod_replset_oplog_head_timestamp - mongodb_mongod_replset_oplog_tail_timestamp)) - sum(avg(mongodb_rs_members_optimeDate{member_state="PRIMARY"} - on (set) group_right mongodb_rs_members_optimeDate{member_state="SECONDARY"})) <= 0
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication headroom (instance {{ $labels.instance }})
    description: |-
        MongoDB replication headroom is <= 0
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbReplicationHeadroom.md

```

{{% /comment %}}

</details>


Here is a runbook for the MongodbReplicationHeadroom alert rule:

## Meaning

The MongodbReplicationHeadroom alert is triggered when the replication headroom in a MongoDB replica set falls to zero or below. Replication headroom refers to the buffer of unread operations in the oplog that allows a secondary node to catch up with the primary node in case of a failover. A headroom of zero or below indicates that the secondary nodes are not able to keep up with the primary node, which can lead to data loss or inconsistencies in case of a failover.

## Impact

The impact of this alert is high, as it can lead to data inconsistencies or loss in case of a failover. If the secondary nodes are not able to keep up with the primary node, it can cause the following issues:

* Data loss or inconsistencies in case of a failover
* Increased downtime and reduced availability of the MongoDB cluster
* Difficulty in recovering from a failover or rolling back to a previous state

## Diagnosis

To diagnose the cause of the low replication headroom, follow these steps:

1. Check the MongoDB logs for any signs of replication issues or errors.
2. Verify that the network connectivity between the nodes is stable and there are no issues with the oplog.
3. Check the resource utilization of the nodes, including CPU, memory, and disk usage, to identify any bottlenecks.
4. Review the MongoDB configuration to ensure that it is optimized for replication.
5. Check the oplog size and adjust it if necessary to ensure that it is large enough to handle the write workload.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and address any underlying issues with the replication process, such as network connectivity or oplog errors.
2. Adjust the MongoDB configuration to optimize replication, such as increasing the oplog size or adjusting the replication timeout.
3. Consider adding more nodes to the replica set to increase the replication capacity.
4. Implement a more robust monitoring and alerting system to detect replication issues earlier.
5. Test the failover process to ensure that it is working correctly and that the secondary nodes can catch up with the primary node quickly.

Remember to consult the official MongoDB documentation and the Percona MongoDB Exporter documentation for more information on troubleshooting and optimizing MongoDB replication.