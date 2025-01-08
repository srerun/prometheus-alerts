---
title: ClickhouseHighTcpConnections
description: Troubleshooting for alert ClickhouseHighTcpConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseHighTcpConnections

High number of TCP connections, indicating heavy client or inter-cluster communication.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseHighTcpConnections" %}}

{{% comment %}}

```yaml
alert: ClickhouseHighTcpConnections
expr: ClickHouseMetrics_TCPConnection > 400
for: 5m
labels:
    severity: warning
annotations:
    summary: ClickHouse High TCP Connections (instance {{ $labels.instance }})
    description: |-
        High number of TCP connections, indicating heavy client or inter-cluster communication.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseHighTcpConnections.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning 

This alert is triggered when the number of TCP connections to a ClickHouse instance exceeds 400 for more than 5 minutes. This could indicate heavy client or inter-cluster communication, which may lead to performance issues or increased latency.

## Impact 

A high number of TCP connections can cause:

* Increased memory usage and CPU load on the ClickHouse instance
* Slower query performance and increased latency
* Potential connection timeouts and errors
* Increased load on the network, potentially affecting other services

## Diagnosis 

To diagnose the issue, follow these steps:

1. Check the ClickHouse instance's metrics to identify the source of the high TCP connection count.
2. Verify that the instance is not experiencing any CPU or memory resource constraints.
3. Review the query logs to identify any long-running or resource-intensive queries that may be contributing to the high connection count.
4. Check for any misconfigured or malfunctioning clients or applications that may be causing the high connection count.

## Mitigation 

To mitigate the issue, follow these steps:

1. Identify and terminate any rogue or idle client connections to the ClickHouse instance.
2. Optimize resource-intensive queries to reduce their impact on the instance.
3. Consider increasing the instance's resource allocation (e.g., CPU, memory) to handle the increased load.
4. Implement connection pooling or other optimization techniques to reduce the number of TCP connections.
5. Review and adjust the ClickHouse configuration to optimize performance and reduce the likelihood of high TCP connection counts in the future.