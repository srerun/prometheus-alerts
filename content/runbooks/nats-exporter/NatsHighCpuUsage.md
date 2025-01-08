---
title: NatsHighCpuUsage
description: Troubleshooting for alert NatsHighCpuUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighCpuUsage

NATS server is using more than 80% CPU for the last 5 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighCpuUsage" %}}

{{% comment %}}

```yaml
alert: NatsHighCpuUsage
expr: rate(gnatsd_varz_cpu[5m]) > 0.8
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high CPU usage (instance {{ $labels.instance }})
    description: |-
        NATS server is using more than 80% CPU for the last 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighCpuUsage.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule "NatsHighCpuUsage":

## Meaning

This alert is triggered when the NATS server's CPU usage exceeds 80% for a period of 5 minutes. This indicates that the server is experiencing high load, which can lead to performance degradation, slower message processing, and even crashes.

## Impact

High CPU usage on the NATS server can have several impacts on the system:

* Slower message processing and increased latency
* Decreased throughput and system performance
* Increased risk of server crashes and downtime
* Potential data loss or corruption

## Diagnosis

To diagnose the root cause of high CPU usage on the NATS server, follow these steps:

1. Check the NATS server logs for any errors or warnings that may indicate the cause of high CPU usage.
2. Verify that the server is not experiencing any memory leaks or issues.
3. Check the system metrics for any signs of resource starvation (e.g., low memory, high disk usage).
4. Investigate any recent changes to the NATS configuration, code, or dependencies that may be contributing to the high CPU usage.
5. Review the NATS message queue statistics to identify any bottlenecks or issues.

## Mitigation

To mitigate high CPU usage on the NATS server, follow these steps:

1. Identify and resolve any underlying issues causing high CPU usage (e.g., fix bugs, optimize code).
2. Scale up the NATS server instance to increase available resources (e.g., CPU, memory).
3. Implement load balancing or clustering to distribute the load across multiple NATS servers.
4. Optimize NATS configuration settings to reduce CPU usage (e.g., adjust batch sizes, timeouts).
5. Consider upgrading the NATS server to a newer version that may have performance improvements.

Remember to investigate and resolve the underlying cause of high CPU usage to prevent recurring issues.