---
title: NatsHighJetstreamMemoryUsage
description: Troubleshooting for alert NatsHighJetstreamMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighJetstreamMemoryUsage

JetStream memory usage is over 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighJetstreamMemoryUsage" %}}

{{% comment %}}

```yaml
alert: NatsHighJetstreamMemoryUsage
expr: gnatsd_varz_jetstream_stats_memory / gnatsd_varz_jetstream_config_max_memory > 0.8
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high JetStream memory usage (instance {{ $labels.instance }})
    description: |-
        JetStream memory usage is over 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighJetstreamMemoryUsage.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The `NatsHighJetstreamMemoryUsage` alert is triggered when the memory usage of JetStream in NATS exceeds 80% of the configured maximum memory. This indicates that JetStream is consuming a significant amount of memory, which can lead to performance issues, increased latency, and potentially even crashes.

## Impact

If this alert is not addressed, the high memory usage can cause:

* Performance degradation: JetStream may become unresponsive or slow, leading to delayed or lost messages.
* Increased latency: Message processing times may increase, affecting the overall system performance.
* Crashes: In extreme cases, NATS may crash or become unstable, causing data loss or system downtime.
* Cascading failures: High memory usage can also impact other components in the system, leading to a broader outage.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NATS logs for any errors or warnings related to JetStream memory usage.
2. Verify the JetStream configuration to ensure it is properly set up and tuned for the current workload.
3. Investigate recent changes to the system, such as increased message volumes or new message types, that may be contributing to the high memory usage.
4. Use tools like `gnatsd_varz` or `nats-top` to monitor JetStream memory usage and identifies areas for optimization.

## Mitigation

To mitigate the issue, follow these steps:

1. **Reduce message payload size**: If possible, reduce the size of message payloads to decrease memory usage.
2. **Optimize JetStream configuration**: Review and adjust JetStream configuration settings, such as `max_memory` or `queue_size`, to ensure they are properly tuned for the current workload.
3. **Scale up or out**: Consider increasing the resources (e.g., RAM, CPU) of the NATS instances or adding more instances to distribute the workload and reduce memory pressure.
4. **Implement message deduplication**: If duplicate messages are contributing to high memory usage, consider implementing message deduplication mechanisms to reduce the load on JetStream.

Remember to continuously monitor JetStream memory usage and adjust the mitigation strategies as needed to prevent similar issues in the future.