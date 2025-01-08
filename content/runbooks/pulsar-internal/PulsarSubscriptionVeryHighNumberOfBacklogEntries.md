---
title: PulsarSubscriptionVeryHighNumberOfBacklogEntries
description: Troubleshooting for alert PulsarSubscriptionVeryHighNumberOfBacklogEntries
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarSubscriptionVeryHighNumberOfBacklogEntries

The number of subscription backlog entries is over 100k

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarSubscriptionVeryHighNumberOfBacklogEntries" %}}

{{% comment %}}

```yaml
alert: PulsarSubscriptionVeryHighNumberOfBacklogEntries
expr: sum(pulsar_subscription_back_log) by (subscription) > 100000
for: 1h
labels:
    severity: critical
annotations:
    summary: Pulsar subscription very high number of backlog entries (instance {{ $labels.instance }})
    description: |-
        The number of subscription backlog entries is over 100k
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarSubscriptionVeryHighNumberOfBacklogEntries.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule:

## Meaning
This alert is triggered when the total number of backlog entries for a Pulsar subscription exceeds 100,000. This indicates that the subscription is not able to keep up with the incoming messages, leading to a large backlog of unprocessed messages.

## Impact
A high backlog of unprocessed messages can cause:

* Delays in processing messages, potentially leading to data loss or staleness
* Increased memory usage on the Pulsar brokers, potentially leading to performance issues or even crashes
* Decreased overall system performance and reliability

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Pulsar subscription logs to identify the root cause of the backlog buildup.
2. Verify that the subscription is properly configured and that the consumers are functioning correctly.
3. Check the message rate and size to identify if there are any unusual patterns or spikes.
4. Verify that the Pulsar brokers have sufficient resources (e.g., memory, CPU) to handle the message load.

## Mitigation
To mitigate the issue, follow these steps:

1. Increase the number of consumers for the subscription to help process the backlog.
2. Check for any message processing bottlenecks and optimize the processing pipeline as needed.
3. Consider increasing the resources (e.g., memory, CPU) available to the Pulsar brokers to handle the message load.
4. Implement message retention policies to prevent backlog buildup in the future.

Note: This is just a sample runbook, and you may need to customize it to fit your specific use case and environment.