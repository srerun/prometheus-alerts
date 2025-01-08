---
title: RedisOutOfConfiguredMaxmemory
description: Troubleshooting for alert RedisOutOfConfiguredMaxmemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisOutOfConfiguredMaxmemory

Redis is running out of configured maxmemory (> 90%)

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisOutOfConfiguredMaxmemory" %}}

{{% comment %}}

```yaml
alert: RedisOutOfConfiguredMaxmemory
expr: redis_memory_used_bytes / redis_memory_max_bytes * 100 > 90 and on(instance) redis_memory_max_bytes > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Redis out of configured maxmemory (instance {{ $labels.instance }})
    description: |-
        Redis is running out of configured maxmemory (> 90%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisOutOfConfiguredMaxmemory.md

```

{{% /comment %}}

</details>


## Meaning

The RedisOutOfConfiguredMaxmemory alert is triggered when the Redis instance's used memory exceeds 90% of its configured maximum allowed memory. This alert indicates that Redis is running low on memory, which can lead to performance issues, slow queries, and even crashes.

## Impact

If left unaddressed, this issue can cause:

* Slow queries and response times
* Increased latency
* Error rates increase
* Redis instance crashes or becomes unresponsive
* Data loss or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Redis instance's memory usage and max memory settings using the Redis CLI or a monitoring tool like Prometheus.
2. Verify that the Redis instance is not experiencing any unusual load or traffic.
3. Check the system logs for any errors or warnings related to Redis.
4. Review the Redis configuration file to ensure that the max memory setting is correctly configured.
5. Check for any memory leaks or inefficiencies in the application using Redis.

## Mitigation

To mitigate this issue, follow these steps:

1. **Increase the max memory setting**: Adjust the Redis instance's max memory setting to a higher value, if possible.
2. **Optimize Redis configuration**: Review and optimize the Redis configuration to ensure it is correctly set up for the workload.
3. **Reduce memory usage**: Identify and address any memory leaks or inefficiencies in the application using Redis.
4. **Scale Redis instance**: Consider scaling the Redis instance to a larger instance type or adding more nodes to the Redis cluster.
5. **Monitor Redis performance**: Closely monitor Redis performance and adjust settings as needed to prevent future occurrences of this issue.

Remember to follow the recommended practices for Redis performance optimization and monitoring to prevent similar issues in the future.