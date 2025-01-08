---
title: PulsarLargeMessagePayload
description: Troubleshooting for alert PulsarLargeMessagePayload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarLargeMessagePayload

Observing large message payload (> 1MB)

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarLargeMessagePayload" %}}

{{% comment %}}

```yaml
alert: PulsarLargeMessagePayload
expr: sum(pulsar_entry_size_overflow > 0) by (topic)
for: 1h
labels:
    severity: warning
annotations:
    summary: Pulsar large message payload (instance {{ $labels.instance }})
    description: |-
        Observing large message payload (> 1MB)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarLargeMessagePayload.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

This alert rule is triggered when the sum of Pulsar entries with a size overflow (larger than 1MB) is greater than 0 for a specific topic over a 1-hour period. This indicates that there are large message payloads being sent to Pulsar, which can cause performance issues and increased storage usage.

## Impact

* Performance degradation: Large message payloads can slow down Pulsar's processing and lead to increased latency.
* Storage usage increase: Large messages can consume a significant amount of storage space, leading to increased storage costs and potentially causing storage capacity issues.
* Potential data loss: If the large messages are not processed correctly, there is a risk of data loss or corruption.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Identify the topic**: Check the `topic` label in the alert to identify which topic is experiencing the large message payloads.
2. **Check message sizes**: Use Pulsar's built-in metrics or a tool like `pulsar-admin` to check the message sizes for the affected topic.
3. **Investigate message content**: Analyze the content of the large messages to determine if they are legitimate or if there is an issue with the message producer.
4. **Check producer configuration**: Verify that the message producer is configured correctly and that there are no issues with the producer's message formatting or serialization.

## Mitigation

To mitigate the issue, follow these steps:

1. **Adjust producer configuration**: If the large messages are due to incorrect producer configuration, adjust the configuration to limit message sizes or implement message compression.
2. **Optimize message serialization**: If the large messages are due to inefficient serialization, optimize the serialization format to reduce message sizes.
3. **Increase storage capacity**: If the large messages are legitimate, consider increasing Pulsar's storage capacity to accommodate the larger message sizes.
4. **Implement message deduplication**: If the large messages are due to duplicate messages, implement message deduplication to reduce the overall message size.