---
title: MongodbReplicationStatus6
description: Troubleshooting for alert MongodbReplicationStatus6
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus6

MongoDB Replication set member as seen from another member of the set, is not yet known

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationStatus6" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationStatus6
expr: mongodb_replset_member_state == 6
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 6 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member as seen from another member of the set, is not yet known
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationStatus6.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "MongodbReplicationStatus6":

## Meaning

The MongodbReplicationStatus6 alert is triggered when a MongoDB replica set member is reported as having a state of 6, indicating that the member's replication status is unknown. This can occur when a new member is added to the replica set, or when there are issues with communication between members.

## Impact

If not addressed, a MongoDB replica set member with an unknown replication status can lead to data inconsistencies, reduced database performance, and potentially even complete cluster failure. This can have a significant impact on applications relying on the database, causing errors, downtime, and revenue loss.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check MongoDB logs**: Review the MongoDB logs on the affected node to identify any errors or warnings related to replication.
2. **Verify replica set configuration**: Confirm that the replica set is properly configured and that all members are correctly listed.
3. **Check network connectivity**: Verify that there are no network issues preventing communication between replica set members.
4. **Run a MongoDB `replSetGetStatus` command**: Execute the `replSetGetStatus` command on the affected node to gather more information about the replica set's status.

## Mitigation

To mitigate the issue, follow these steps:

1. **Restart the MongoDB service**: Restart the MongoDB service on the affected node to attempt to re-establish replication.
2. **Verify replica set configuration**: Double-check the replica set configuration to ensure it is correct and consistent across all members.
3. **Check for and apply MongoDB updates**: Ensure that all MongoDB nodes are running the latest version and apply any available updates.
4. **Monitor MongoDB performance**: Closely monitor MongoDB performance and replication status to ensure the issue is resolved.

If the issue persists, consider involving a MongoDB administrator or expert for further assistance.