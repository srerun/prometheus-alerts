---
title: MongodbReplicationLag
description: Troubleshooting for alert MongodbReplicationLag
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationLag

Mongodb replication lag is more than 10s

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbReplicationLag" %}}

{{% comment %}}

```yaml
alert: MongodbReplicationLag
expr: (mongodb_rs_members_optimeDate{member_state="PRIMARY"} - on (set) group_right mongodb_rs_members_optimeDate{member_state="SECONDARY"}) / 1000 > 10
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication lag (instance {{ $labels.instance }})
    description: |-
        Mongodb replication lag is more than 10s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbReplicationLag.md

```

{{% /comment %}}

</details>


Here is a runbook for the MongodbReplicationLag alert rule:

## Meaning

The MongodbReplicationLag alert is triggered when the replication lag in a MongoDB replica set exceeds 10 seconds. This means that the secondary members of the replica set are not keeping up with the primary member, and data written to the primary is taking too long to be replicated to the secondaries.

## Impact

If left unaddressed, this alert can lead to:

* Data inconsistencies between primary and secondary members
* Increased risk of data loss in the event of a primary failure
* Decreased performance and availability of the MongoDB cluster
* Potential for application errors and downtime

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the MongoDB logs for any errors or warnings related to replication
2. Verify that the replica set is properly configured and all members are up and running
3. Check the network connectivity and latency between primary and secondary members
4. Investigate any recent changes to the MongoDB configuration or application workload
5. Use the `mongodb_rs_members_optimeDate` metric to identify the specific replica set and members experiencing the lag

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and address any underlying network or infrastructure issues causing the lag
2. Optimize MongoDB configuration for improved replication performance (e.g. adjusting the `syncTimeout` value)
3. Increase the resources (e.g. CPU, RAM) of the secondary members to improve their ability to keep up with the primary
4. Consider adding additional secondary members to the replica set to distribute the load and improve replication performance
5. Perform a manual replication sync to catch up the secondary members, if necessary

Remember to investigate and address the root cause of the issue to prevent future occurrences of this alert.