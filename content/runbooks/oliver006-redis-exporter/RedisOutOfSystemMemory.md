---
title: RedisOutOfSystemMemory
description: Troubleshooting for alert RedisOutOfSystemMemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisOutOfSystemMemory

Redis is running out of system memory (> 90%)

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisOutOfSystemMemory" %}}

{{% comment %}}

```yaml
alert: RedisOutOfSystemMemory
expr: redis_memory_used_bytes / redis_total_system_memory_bytes * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Redis out of system memory (instance {{ $labels.instance }})
    description: |-
        Redis is running out of system memory (> 90%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisOutOfSystemMemory.md

```

{{% /comment %}}

</details>


Here is a runbook for the RedisOutOfSystemMemory alert:

## Meaning

The RedisOutOfSystemMemory alert is triggered when the used memory of a Redis instance exceeds 90% of the total system memory. This indicates that Redis is running low on available memory, which can lead to performance issues, increased latency, and even crashes.

## Impact

If left unaddressed, this alert can have a significant impact on the system's performance and availability. Some potential consequences include:

* Slow response times and increased latency for Redis queries
* Increased risk of Redis crashes or failures, leading to data loss or corruption
* Cascading failures in dependent systems or applications
* Decreased system reliability and availability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Redis instance's memory usage using the `redis_info` command or a monitoring tool like Prometheus.
2. Verify that the system has sufficient memory available and that there are no other resource-intensive processes running on the same host.
3. Review the Redis configuration to ensure that the `maxmemory` setting is set appropriately for the instance.
4. Check for any memory leaks or inefficient usage in Redis or dependent applications.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately add more memory to the system or terminate any unnecessary processes to free up memory.
2. Investigate and address any memory leaks or inefficient usage in Redis or dependent applications.
3. Adjust the `maxmemory` setting in the Redis configuration to prevent memory exhaustion.
4. Consider implementing memory-efficient data structures and algorithms in Redis or dependent applications.
5. Monitor the Redis instance's memory usage closely and adjust the alert threshold as necessary to prevent false positives.

Remember to update the alert rule and runbook accordingly to reflect the changes made to the system.