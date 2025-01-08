---
title: NatsHighNumberOfConnections
description: Troubleshooting for alert NatsHighNumberOfConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighNumberOfConnections

NATS server has more than 1000 active connections

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighNumberOfConnections" %}}

{{% comment %}}

```yaml
alert: NatsHighNumberOfConnections
expr: gnatsd_connz_num_connections > 1000
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high number of connections (instance {{ $labels.instance }})
    description: |-
        NATS server has more than 1000 active connections
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighNumberOfConnections.md

```

{{% /comment %}}

</details>


## Meaning

The NatsHighNumberOfConnections alert is triggered when the number of active connections to a NATS server exceeds 1000. This alert indicates that the NATS server is experiencing a high load, which can lead to performance issues, latency, and even crashes.

## Impact

* Performance degradation: High number of connections can cause the NATS server to slow down, leading to increased latency and decreased throughput.
* Resource exhaustion: A large number of connections can consume system resources such as memory, CPU, and network bandwidth, leading to resource exhaustion.
* Increased error rates: High connection counts can lead to increased error rates, including connection timeouts, failures, and message losses.
* Potential crashes: If left unchecked, a high number of connections can cause the NATS server to crash, leading to service disruptions and downtime.

## Diagnosis

To diagnose the root cause of the high number of connections, follow these steps:

1. Check the NATS server logs for any errors or warnings related to connection handling.
2. Verify that the NATS server is properly configured to handle a high volume of connections.
3. Investigate recent changes to the system, such as new applications or services that may be contributing to the increased connection count.
4. Use Prometheus metrics to analyze the connection pattern and identify any anomalies.
5. Verify that the NATS server has sufficient system resources (CPU, memory, and network bandwidth) to handle the increased load.

## Mitigation

To mitigate the high number of connections, follow these steps:

1. Increase the NATS server's system resources (CPU, memory, and network bandwidth) to handle the increased load.
2. Implement connection limits or throttling to prevent excessive connections.
3. Optimize NATS server configuration to improve performance and reduce latency.
4. Identify and troubleshoot any application or service that is causing the high connection count.
5. Consider load balancing or clustering NATS servers to distribute the load and improve overall system resilience.
6. Monitor the NATS server's performance and connection count regularly to prevent similar issues in the future.