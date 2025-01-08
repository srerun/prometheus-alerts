---
title: MongodbReplicaMemberUnhealthy
description: Troubleshooting for alert MongodbReplicaMemberUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicaMemberUnhealthy

MongoDB replica member is not healthy

<details>
  <summary>Alert Rule</summary>

{{% rule "mongodb/percona-mongodb-exporter.yml" "MongodbReplicaMemberUnhealthy" %}}

{{% comment %}}

```yaml
alert: MongodbReplicaMemberUnhealthy
expr: mongodb_rs_members_health == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Mongodb replica member unhealthy (instance {{ $labels.instance }})
    description: |-
        MongoDB replica member is not healthy
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbReplicaMemberUnhealthy.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `MongodbReplicaMemberUnhealthy`:

## Meaning

The `MongodbReplicaMemberUnhealthy` alert is triggered when a MongoDB replica member is not healthy. This means that the replica member is not able to communicate with the primary node or other replica members, which can lead to data inconsistencies and impact the overall performance of the MongoDB cluster.

## Impact

The impact of an unhealthy MongoDB replica member can be significant, as it can lead to:

* Data inconsistencies and potential data loss
* Reduced performance and availability of the MongoDB cluster
* Increased latency and errors for applications relying on the MongoDB cluster
* Potential for secondary effects on dependent systems and services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the MongoDB replica member's logs for errors or warnings related to connectivity or replication.
2. Verify the replica member's status using the `mongo` shell or the Percona MongoDB Exporter.
3. Check the network connectivity between the replica member and the primary node.
4. Validate the replica member's configuration and ensure it matches the expected configuration.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the MongoDB replica member service to attempt to re-establish connectivity.
2. Check and repair any disk errors or corruption on the replica member.
3. Verify the replica member's configuration and update it if necessary.
4. If the issue persists, consider re-syncing the replica member with the primary node.
5. If the issue is related to network connectivity, investigate and resolve any underlying network issues.

Remember to refer to the [Percona MongoDB Exporter documentation](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/percona-mongodb-exporter/MongodbReplicaMemberUnhealthy.md) for additional guidance and best practices for troubleshooting and resolving MongoDB replica member health issues.