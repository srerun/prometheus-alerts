---
title: PulsarTopicVeryLargeBacklogStorageSize
description: Troubleshooting for alert PulsarTopicVeryLargeBacklogStorageSize
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarTopicVeryLargeBacklogStorageSize

The topic backlog storage size is over 20 GB

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarTopicVeryLargeBacklogStorageSize" %}}

{{% comment %}}

```yaml
alert: PulsarTopicVeryLargeBacklogStorageSize
expr: sum(pulsar_storage_size > 20*1024*1024*1024) by (topic)
for: 1h
labels:
    severity: critical
annotations:
    summary: Pulsar topic very large backlog storage size (instance {{ $labels.instance }})
    description: |-
        The topic backlog storage size is over 20 GB
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarTopicVeryLargeBacklogStorageSize.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `PulsarTopicVeryLargeBacklogStorageSize`:

## Meaning

The `PulsarTopicVeryLargeBacklogStorageSize` alert is triggered when the storage size of a Pulsar topic backlog exceeds 20 GB. This can indicate a potential issue with message processing or retention in the Pulsar cluster.

## Impact

The impact of this alert is high, as a large backlog storage size can lead to:

* Increased storage costs
* Performance degradation of the Pulsar cluster
* Potential message loss or duplication
* Increased latency for message processing

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Identify the affected topic(s) by checking the `topic` label in the alert.
2. Check the Pulsar cluster's storage usage and availability.
3. Investigate the message processing rate and latency for the affected topic(s).
4. Verify that there are no issues with the message producer or consumer applications.
5. Check the Pulsar cluster's configuration and retention settings.

## Mitigation

To mitigate this alert, follow these steps:

1. Identify and resolve any issues with message processing or retention in the Pulsar cluster.
2. Consider increasing the storage capacity of the Pulsar cluster.
3. Adjust the retention settings or message TTL (time to live) for the affected topic(s).
4. Implement a more efficient message processing strategy, such as using a larger consumer group or increasing the consumer instances.
5. Monitor the Pulsar cluster's storage usage and message processing rate to ensure the issue does not recur.