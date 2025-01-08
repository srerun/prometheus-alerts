---
title: RedisDisconnectedSlaves
description: Troubleshooting for alert RedisDisconnectedSlaves
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisDisconnectedSlaves

Redis not replicating for all slaves. Consider reviewing the redis replication status.

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisDisconnectedSlaves" %}}

{{% comment %}}

```yaml
alert: RedisDisconnectedSlaves
expr: count without (instance, job) (redis_connected_slaves) - sum without (instance, job) (redis_connected_slaves) - 1 > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis disconnected slaves (instance {{ $labels.instance }})
    description: |-
        Redis not replicating for all slaves. Consider reviewing the redis replication status.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisDisconnectedSlaves.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `RedisDisconnectedSlaves`:

## Meaning

The `RedisDisconnectedSlaves` alert indicates that one or more Redis slaves are disconnected from their master. This means that data replication is not occurring, and the slaves are no longer receiving updates from the master. This can lead to data inconsistencies and potential data loss if not addressed promptly.

## Impact

The impact of this alert is critical, as it can result in:

* Data inconsistencies between the master and slaves
* Potential data loss if the master fails and the slaves are not updated
* Inconsistent application behavior, as the slaves may not reflect the latest data changes
* Increased risk of data corruption or loss in the event of a failover

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Redis replication status using the `redis-cli` command or a Redis GUI tool
2. Verify that the Redis master is properly configured and running
3. Check the network connectivity between the Redis master and slaves
4. Review the Redis logs for any error messages related to replication
5. Check the `redis_connected_slaves` metric in Prometheus to identify which slaves are disconnected

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately investigate and resolve any network connectivity issues between the Redis master and slaves
2. Verify that the Redis master and slaves are properly configured and running
3. Restart the Redis slaves to re-establish replication
4. Monitor the Redis replication status and metrics to ensure that the issue is resolved
5. Consider setting up additional monitoring and alerting for Redis replication status to prevent similar issues in the future