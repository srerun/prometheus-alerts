---
title: PulsarHighWriteLatency
description: Troubleshooting for alert PulsarHighWriteLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarHighWriteLatency

Messages cannot be written in a timely fashion

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarHighWriteLatency" %}}

{{% comment %}}

```yaml
alert: PulsarHighWriteLatency
expr: sum(pulsar_storage_write_latency_overflow > 0) by (topic)
for: 1h
labels:
    severity: critical
annotations:
    summary: Pulsar high write latency (instance {{ $labels.instance }})
    description: |-
        Messages cannot be written in a timely fashion
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarHighWriteLatency.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `PulsarHighWriteLatency`:

## Meaning

The `PulsarHighWriteLatency` alert is triggered when there is an excessive number of write latency overflows in Pulsar, indicating that messages are not being written in a timely fashion. This can lead to performance issues, data loss, or even complete system unavailability.

## Impact

The impact of this alert can be severe, as it affects the ability to write data to Pulsar. This can lead to:

* Performance degradation: High write latency can cause slower data processing, leading to delays in downstream applications.
* Data loss: Unwritten messages may be lost, resulting in incomplete or inaccurate data.
* System unavailability: In extreme cases, high write latency can cause Pulsar to become unavailable, leading to complete system downtime.

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the Pulsar cluster configuration to ensure that it is properly sized and configured for the current workload.
2. Investigate disk usage and I/O performance to identify any bottlenecks or issues with storage.
3. Review Pulsar broker logs for errors or exceptions related to write operations.
4. Verify that there are no network connectivity issues between Pulsar brokers and bookies.
5. Check the topic configuration to ensure that the write latency threshold is set correctly.

## Mitigation

To mitigate this alert, follow these steps:

1. Scale up the Pulsar cluster to increase capacity and reduce write latency.
2. Optimize disk usage and I/O performance by increasing storage capacity, improving disk configuration, or using faster storage.
3. Implement load balancing to distribute write traffic across multiple brokers.
4. Implement queuing or buffering mechanisms to handle message bursts and reduce write latency.
5. Consider upgrading Pulsar version to take advantage of performance improvements and bug fixes.

Remember to investigate and address the root cause of the issue, rather than just treating the symptoms. Once the issue is resolved, the alert should clear automatically.