---
title: RedisReplicationBroken
description: Troubleshooting for alert RedisReplicationBroken
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisReplicationBroken

Redis instance lost a slave

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisReplicationBroken" %}}

{{% comment %}}

```yaml
alert: RedisReplicationBroken
expr: delta(redis_connected_slaves[1m]) < 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis replication broken (instance {{ $labels.instance }})
    description: |-
        Redis instance lost a slave
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisReplicationBroken.md

```

{{% /comment %}}

</details>


## Meaning

The `RedisReplicationBroken` alert is triggered when the number of connected Redis slaves decreases over a 1-minute period. This indicates that one or more Redis slaves have disconnected from the master instance, which can lead to data inconsistencies and loss.

## Impact

The impact of this alert is critical, as it can result in:

* Data loss or inconsistencies between Redis instances
* Increased latency or errors in applications relying on Redis
* Potential for data inconsistencies to spread to other nodes in the cluster
* Downtime or reduced performance of dependent services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Redis instance's log files for errors or warnings related to replication or connections.
2. Verify the network connectivity between the Redis master and slave instances.
3. Check the Redis configuration to ensure that replication is properly configured.
4. Review the `redis_connected_slaves` metric to identify the specific slave instance that disconnected.
5. Investigate any recent changes or updates to the Redis cluster or underlying infrastructure.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve the underlying cause of the slave disconnection (e.g., network issue, configuration error, etc.).
2. Re-establish connectivity between the Redis master and slave instances.
3. Verify that replication is working correctly and data is being synced between instances.
4. Consider increasing the number of Redis slaves to improve redundancy and fault tolerance.
5. Review and adjust the Redis cluster's configuration to prevent similar issues in the future.