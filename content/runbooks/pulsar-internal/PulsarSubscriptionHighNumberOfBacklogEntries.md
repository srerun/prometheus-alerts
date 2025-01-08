---
title: PulsarSubscriptionHighNumberOfBacklogEntries
description: Troubleshooting for alert PulsarSubscriptionHighNumberOfBacklogEntries
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarSubscriptionHighNumberOfBacklogEntries

The number of subscription backlog entries is over 5k

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarSubscriptionHighNumberOfBacklogEntries" %}}

{{% comment %}}

```yaml
alert: PulsarSubscriptionHighNumberOfBacklogEntries
expr: sum(pulsar_subscription_back_log) by (subscription) > 5000
for: 1h
labels:
    severity: warning
annotations:
    summary: Pulsar subscription high number of backlog entries (instance {{ $labels.instance }})
    description: |-
        The number of subscription backlog entries is over 5k
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarSubscriptionHighNumberOfBacklogEntries.md

```

{{% /comment %}}

</details>


## Meaning

The PulsarSubscriptionHighNumberOfBacklogEntries alert is triggered when the number of backlog entries for a Pulsar subscription exceeds 5000. This alert indicates that the subscription is experiencing a high volume of unprocessed messages, which can lead to performance issues, latency, and potential message loss.

## Impact

The impact of this alert can be significant, as a high number of backlog entries can:

* Cause delays in message processing, affecting the overall performance of the system
* Lead to increased memory usage, potentially causing node crashes or restarts
* Result in message loss or duplication, affecting data integrity and consistency
* Affect the overall reliability and availability of the Pulsar cluster

## Diagnosis

To diagnose the cause of this alert, follow these steps:

1. Check the Pulsar subscription configuration to ensure it is properly configured and scaled to handle the expected message volume.
2. Verify that the Pulsar broker nodes have sufficient resources (e.g., CPU, memory, and disk space) to handle the backlog.
3. Investigate if there are any issues with the message producers or consumers that may be contributing to the backlog.
4. Check the Pulsar metrics to identify any trends or patterns that may indicate the root cause of the issue.
5. Review the Pulsar subscription logs to identify any errors or exceptions that may be related to the backlog.

## Mitigation

To mitigate this alert, follow these steps:

1. Increase the resources (e.g., CPU, memory, and disk space) of the Pulsar broker nodes to handle the backlog.
2. Adjust the Pulsar subscription configuration to optimize message processing and reduce the backlog.
3. Implement message consolidation or deduplication to reduce the volume of messages being processed.
4. Investigate and resolve any issues with message producers or consumers that may be contributing to the backlog.
5. Monitor the Pulsar subscription metrics and logs to ensure the backlog is decreasing and the system is returning to a stable state.