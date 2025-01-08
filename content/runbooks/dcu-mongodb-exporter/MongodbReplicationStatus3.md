---
title: MongodbReplicationStatus3
description: Troubleshooting for alert MongodbReplicationStatus3
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus3

MongoDB Replication set member either perform startup self-checks, or transition from completing a rollback or resync

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationStatus3" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationStatus3
expr: mongodb_replset_member_state == 3
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 3 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member either perform startup self-checks, or transition from completing a rollback or resync
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationStatus3.md

```

{{% /comment %}}

</details>


## Meaning

The MongoDBReplicationStatus3 alert is triggered when a MongoDB replica set member is in a state of replication status 3. This state indicates that the member is either performing startup self-checks or transitioning from completing a rollback or resync. This alert is critical, indicating a potential issue with the MongoDB replication process.

## Impact

The impact of this alert is that the MongoDB replica set member may not be fully operational, which can lead to:

* Data inconsistencies between nodes
* Reduced redundancy and high availability
* Increased risk of data loss or corruption
* Potential performance degradation

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the MongoDB replica set status using the `mongo` shell or a MongoDB monitoring tool.
2. Investigate the replication status of the affected node using the `rs.status()` command.
3. Review the MongoDB logs to identify any error messages or issues related to replication.
4. Verify that the node is properly configured and that there are no network connectivity issues.
5. Check for any recent changes or updates to the MongoDB configuration or deployment.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the MongoDB documentation for troubleshooting replication issues.
2. Restart the affected MongoDB node to allow it to rejoin the replica set.
3. If the issue persists, consider re-configuring the replica set or seeking assistance from a MongoDB expert.
4. Verify that the MongoDB replica set is properly configured and that all nodes are in a healthy state.
5. Consider implementing additional monitoring and alerting to detect replication issues earlier.