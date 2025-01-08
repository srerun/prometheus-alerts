---
title: MongodbReplicationStatus9
description: Troubleshooting for alert MongodbReplicationStatus9
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus9

MongoDB Replication set member is actively performing a rollback. Data is not available for reads

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationStatus9" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationStatus9
expr: mongodb_replset_member_state == 9
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 9 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member is actively performing a rollback. Data is not available for reads
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationStatus9.md

```

{{% /comment %}}

</details>


## Meaning

The MongodbReplicationStatus9 alert is triggered when a MongoDB replica set member enters a rollback state, indicated by a replication status of 9. This means that the node is actively reverting to a previous state, and data is not available for reads until the rollback is complete.

## Impact

The impact of this alert is high, as it indicates that data is temporarily unavailable for reads, which can affect the performance and availability of dependent applications. Additionally, rollback operations can lead to data inconsistencies and potential data loss if not addressed promptly.

## Diagnosis

To diagnose the issue, check the following:

* Verify the MongoDB replica set member's status using the MongoDB shell or a monitoring tool like MongoDB Atlas or MongoDB Compass.
* Check the MongoDB logs for any error messages related to the rollback operation.
* Investigate the cause of the rollback, such as network issues, disk space problems, or software errors.
* Check the replication lag and oplog size to ensure they are within acceptable limits.

## Mitigation

To mitigate the issue, follow these steps:

* Immediately investigate the cause of the rollback and address the underlying issue.
* If the rollback is due to a network issue, ensure that the network connection is stable and reliable.
* If the rollback is due to disk space issues, free up disk space or consider increasing storage capacity.
* If the rollback is due to software errors, apply any necessary patches or updates.
* Once the underlying issue is resolved, allow the MongoDB replica set member to complete the rollback operation.
* Monitor the replication status and oplog size to ensure they return to a healthy state.
* Perform a manual resync of the replica set member to ensure data consistency.
* Consider implementing additional monitoring and automation to prevent or detect similar issues in the future.