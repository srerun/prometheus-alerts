---
title: RedisRejectedConnections
description: Troubleshooting for alert RedisRejectedConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RedisRejectedConnections

Some connections to Redis has been rejected

<details>
  <summary>Alert Rule</summary>

{{% rule "redis/oliver006-redis-exporter.yml" "RedisRejectedConnections" %}}

{{% comment %}}

```yaml
alert: RedisRejectedConnections
expr: increase(redis_rejected_connections_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Redis rejected connections (instance {{ $labels.instance }})
    description: |-
        Some connections to Redis has been rejected
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/oliver006-redis-exporter/RedisRejectedConnections.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the RedisRejectedConnections alert:

## Meaning

The RedisRejectedConnections alert indicates that Redis has rejected one or more incoming connections within the last minute. This could be due to various reasons such as Redis reaching its maximum connection limit, network issues, or configuration problems.

## Impact

If left unaddressed, rejected connections to Redis can lead to:

* Increased latency and timeouts for applications relying on Redis
* Data loss or inconsistencies due to failed writes or reads
* Cascading failures in dependent services or applications
* Decreased overall system performance and reliability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Redis server logs for error messages related to connection rejections
2. Verify the current connection count and maximum allowed connections using the `redis_info` metric
3. Investigate network connectivity issues between the Redis instance and connecting clients
4. Review recent changes to Redis configuration or deployment

## Mitigation

To mitigate the issue, follow these steps:

1. **Temporary Fix**: Increase the maximum allowed connections limit on the Redis instance to accommodate the current connection demand
2. **Root Cause Analysis**: Identify and address the underlying cause of the connection rejections (e.g., network issues, configuration problems)
3. **Monitoring and Alerting**: Implement additional monitoring and alerting to detect connection rejections and capacity issues proactively
4. **Long-term Solution**: Consider scaling Redis instances, distributing load across multiple instances, or optimizing Redis configuration for better performance and reliability.

Remember to update the alert annotations with the root cause and mitigation steps taken to resolve the issue.