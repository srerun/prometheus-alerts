---
title: MongodbReplicationStatus8
description: Troubleshooting for alert MongodbReplicationStatus8
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus8

MongoDB Replication set member as seen from another member of the set, is unreachable

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/dcu-mongodb-exporter.yml" "MongodbReplicationStatus8" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationStatus8
expr: mongodb_replset_member_state == 8
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 8 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member as seen from another member of the set, is unreachable
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/dcu-mongodb-exporter/MongodbReplicationStatus8.md

```

{{% /comment %}}

</details>


Here is a runbook for the MongodbReplicationStatus8 alert rule:

## Meaning

The MongodbReplicationStatus8 alert rule is triggered when a MongoDB replication set member is unreachable from another member of the set. This means that the MongoDB instance is not responding or is not reachable, which can cause data inconsistencies and availability issues.

## Impact

The impact of this alert is critical, as it can result in:

* Data loss or inconsistencies
* Unavailability of the MongoDB instance
* Impact on dependent applications and services
* Potential downtime and revenue loss

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the MongoDB instance status using the MongoDB export metrics.
2. Verify the network connectivity between the MongoDB instances.
3. Check the MongoDB logs for any error messages related to replication.
4. Verify that the MongoDB configuration is correct and up-to-date.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the MongoDB instance if it is not responding.
2. Check and resolve any network connectivity issues.
3. Investigate and resolve any underlying issues causing the replication failure.
4. Verify that the MongoDB replication set is properly configured and functional.
5. Consider performing a failover to a secondary node if the primary node is unreachable.

Note: Refer to the MongoDB documentation and the MongoDB exporter documentation for more detailed troubleshooting steps and configuration recommendations.