---
title: PulsarTopicLargeBacklogStorageSize
description: Troubleshooting for alert PulsarTopicLargeBacklogStorageSize
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarTopicLargeBacklogStorageSize

The topic backlog storage size is over 5 GB

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarTopicLargeBacklogStorageSize" %}}

{{% comment %}}

```yaml
alert: PulsarTopicLargeBacklogStorageSize
expr: sum(pulsar_storage_size > 5*1024*1024*1024) by (topic)
for: 1h
labels:
    severity: warning
annotations:
    summary: Pulsar topic large backlog storage size (instance {{ $labels.instance }})
    description: |-
        The topic backlog storage size is over 5 GB
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarTopicLargeBacklogStorageSize.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `PulsarTopicLargeBacklogStorageSize`:

## Meaning

This alert is triggered when the storage size of a Pulsar topic's backlog exceeds 5 GB. This indicates that the topic is experiencing a backlog of unprocessed messages, which can lead to performance issues, increased latency, and potentially even data loss.

## Impact

The impact of this alert is moderate to high, as a large backlog of unprocessed messages can:

* Cause performance issues and increased latency for producers and consumers
* Lead to data loss if the backlog is not addressed in a timely manner
* Potentially cause other downstream systems to fail or become unstable

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Identify the affected Pulsar topic using the `topic` label in the alert.
2. Check the Pulsar topic's configuration to ensure that it is properly sized for the expected message volume.
3. Investigate the producer and consumer configurations to ensure that they are functioning correctly.
4. Check the Pulsar cluster's overall performance and resource utilization to ensure that it is not experiencing any bottlenecks.

## Mitigation

To mitigate this alert, follow these steps:

1. Increase the storage capacity of the affected Pulsar topic to accommodate the backlog of messages.
2. Identify and address any issues with the producer or consumer configurations that may be contributing to the backlog.
3. Consider implementing a message retention policy to automatically remove older messages from the topic.
4. Monitor the Pulsar topic's storage size and adjust the retention policy as needed to prevent future backlogs.

Additional resources:

* [Pulsar Topic Configuration](https://pulsar.apache.org/docs/en/configure-topic/)
* [Pulsar Performance Tuning](https://pulsar.apache.org/docs/en/performance-tuning/)
* [Pulsar Message Retention](https://pulsar.apache.org/docs/en/message-retention/)