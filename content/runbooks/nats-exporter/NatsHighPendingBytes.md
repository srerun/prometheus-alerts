---
title: NatsHighPendingBytes
description: Troubleshooting for alert NatsHighPendingBytes
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighPendingBytes

High number of NATS pending bytes

## Meaning

This alert triggers when the number of pending bytes for NATS connections exceeds 100,000 bytes for a duration of 3 minutes. Pending bytes represent messages waiting to be delivered to clients, indicating potential performance issues or bottlenecks in message processing.

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighPendingBytes" %}}

{{% comment %}}

```yaml
alert: NatsHighPendingBytes
expr: gnatsd_connz_pending_bytes > 100000
for: 3m
labels:
    severity: warning
annotations:
    summary: Nats high pending bytes (instance {{ $labels.instance }})
    description: |-
        High number of NATS pending bytes ({{ $value }}) for {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighPendingBytes.md

```

{{% /comment %}}

</details>


## Impact

- High pending bytes may cause delays in message delivery, impacting downstream systems or applications relying on NATS for real-time data.
- Prolonged high pending bytes can lead to resource exhaustion in the NATS server, including increased memory usage and potential message loss.
- It may indicate a slow or unresponsive client, insufficient server resources, or network congestion.

## Diagnosis

1. **Check the Alert Details:**
   - Review the alert annotations to identify the affected instance and the current value of pending bytes.
   - Example:
     - Instance: `{{ $labels.instance }}`
     - Pending Bytes: `{{ $value }}`

2. **Inspect Client Connections:**
   - Access the NATS management interface or use monitoring tools to view detailed connection statistics.
   - Identify clients with high pending bytes and verify if they are slow or unresponsive.

3. **Review Resource Utilization:**
   - Check the NATS server’s CPU, memory, and network usage to ensure sufficient resources are available.

4. **Check Application Logs:**
   - Look for errors or warnings in the logs of applications connected to NATS.
   - Identify any issues with message processing or acknowledgment.

5. **Network Analysis:**
   - Analyze network performance between the NATS server and clients to identify potential bottlenecks.

## Mitigation

1. **Address Client Issues:**
   - Restart or fix any unresponsive or slow clients causing the high pending bytes.
   - Optimize client applications to process messages more efficiently.

2. **Scale NATS Infrastructure:**
   - Add more NATS server instances or increase resources (CPU, memory) for existing servers.
   - Consider implementing clustering for better load distribution.

3. **Optimize Message Flow:**
   - Reduce the volume of messages published to NATS if unnecessary or redundant data is being sent.
   - Adjust NATS server and client configuration settings, such as buffer sizes, to better handle the message load.

4. **Network Improvements:**
   - Investigate and resolve any network-related issues, such as high latency or packet loss, between the NATS server and clients.

5. **Clear Message Backlog:**
   - If safe, purge undelivered messages for non-critical topics or queues to reduce pending bytes.

## Preventative Measures

- Implement proactive monitoring for NATS metrics such as pending bytes, message rates, and resource utilization.
- Establish alert thresholds and response plans tailored to your system’s load patterns.
- Regularly test and validate client and server configurations for optimal performance.
- Ensure adequate resource allocation for NATS servers and clients based on anticipated workloads.

