---
title: RedisClusterFlapping
description: Troubleshooting for alert RedisClusterFlapping
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisClusterFlapping

Changes have been detected in Redis replica connection. This can occur when replica nodes lose connection to the master and reconnect (a.k.a flapping).

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisClusterFlapping" %}}

{{% comment %}}

```yaml
alert: RedisClusterFlapping
expr: changes(redis_connected_slaves[1m]) > 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Redis cluster flapping (instance {{ $labels.instance }})
    description: |-
        Changes have been detected in Redis replica connection. This can occur when replica nodes lose connection to the master and reconnect (a.k.a flapping).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisClusterFlapping.md

```

{{% /comment %}}

</details>


Here is a runbook for the RedisClusterFlapping alert:

## Meaning
The RedisClusterFlapping alert is triggered when there are frequent changes in the number of connected Redis replicas within a short period of time (1 minute). This indicates that the Replica nodes are constantly losing and re-establishing connection to the Master node, causing instability in the Redis cluster.

## Impact
The flapping of Redis replicas can lead to:

* Data inconsistencies: As replicas disconnect and reconnect, they may not always be in sync with the Master node, leading to data inconsistencies.
* Performance issues: The constant reconnecting and syncing of replicas can cause performance issues, such as increased latency and decreased throughput.
* Increased load on Master node: The Master node may experience increased load as it tries to keep the replicas in sync, leading to performance degradation.

## Diagnosis
To diagnose the issue, follow these steps:

* Check the Redis cluster logs for errors or warnings related to replica connections.
* Verify that the network connectivity between the Master node and Replica nodes is stable.
* Check the system resources (CPU, memory, disk space) of the Master node and Replica nodes to ensure they are not overloaded.
* Review the Redis configuration to ensure that it is correctly set up and that the replica nodes are properly configured.

## Mitigation
To mitigate the issue, follow these steps:

* Investigate and resolve any underlying network issues that may be causing the replicas to lose connection to the Master node.
* Check and adjust the Redis configuration to ensure that the replica nodes are properly configured and that the Master node is not overloaded.
* Consider increasing the `redis-connected-slaves` metric threshold to reduce the sensitivity of the alert.
* Implement measures to reduce the load on the Master node, such as load balancing or sharding.
* Consider upgrading the Redis version or using a more stable Redis cluster setup.