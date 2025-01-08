---
title: RedisTooManyConnections
description: Troubleshooting for alert RedisTooManyConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisTooManyConnections

Redis is running out of connections (> 90% used)

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisTooManyConnections" %}}

{{% comment %}}

```yaml
alert: RedisTooManyConnections
expr: redis_connected_clients / redis_config_maxclients * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Redis too many connections (instance {{ $labels.instance }})
    description: |-
        Redis is running out of connections (> 90% used)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisTooManyConnections.md

```

{{% /comment %}}

</details>


Here is a runbook for the RedisTooManyConnections alert rule:

## Meaning

The RedisTooManyConnections alert is triggered when the number of connected clients to Redis exceeds 90% of the maximum allowed connections. This alert indicates that Redis is running out of connections, which can lead to performance issues and potentially even crashes.

## Impact

The impact of this alert can be significant, as it can cause:

* Slow performance and increased latency for applications relying on Redis
* Increased risk of Redis crashes or failures
* Potential data loss or corruption
* Disruption to dependent services and applications

## Diagnosis

To diagnose the root cause of the RedisTooManyConnections alert, follow these steps:

1. Check the Redis configuration file to verify that the `maxclients` setting is correctly set.
2. Investigate the number of connected clients and identify any patterns or spikes in connection usage.
3. Review the application logs to identify any issues or errors related to Redis connections.
4. Check the Redis metrics to identify any performance bottlenecks or issues.

## Mitigation

To mitigate the RedisTooManyConnections alert, follow these steps:

1. Increase the `maxclients` setting in the Redis configuration file to allow for more connections.
2. Optimize application behavior to reduce the number of connections required.
3. Implement connection pooling or other connection management strategies to reduce the load on Redis.
4. Consider horizontal scaling or sharding Redis to distribute the load and increase capacity.
5. Implement monitoring and alerting to detect potential issues before they become critical.

Additional resources:

* Refer to the Redis documentation for configuration and tuning guidelines.
* Consult with application developers and Redis administrators to identify and address root causes.
* Review Redis best practices for connection management and performance optimization.