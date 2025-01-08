---
title: NatsHighJetstreamStoreUsage
description: Troubleshooting for alert NatsHighJetstreamStoreUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsHighJetstreamStoreUsage

JetStream store usage is over 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsHighJetstreamStoreUsage" %}}

{{% comment %}}

```yaml
alert: NatsHighJetstreamStoreUsage
expr: gnatsd_varz_jetstream_stats_storage / gnatsd_varz_jetstream_config_max_storage > 0.8
for: 5m
labels:
    severity: warning
annotations:
    summary: Nats high JetStream store usage (instance {{ $labels.instance }})
    description: |-
        JetStream store usage is over 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsHighJetstreamStoreUsage.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `NatsHighJetstreamStoreUsage`:

## Meaning

The `NatsHighJetstreamStoreUsage` alert is triggered when the JetStream store usage of a NATS instance exceeds 80% of the configured maximum storage capacity. This alert indicates that the NATS instance is running low on storage space, which can lead to performance issues and potentially cause data loss.

## Impact

If left unaddressed, high JetStream store usage can cause:

* Performance degradation: As storage usage increases, NATS may slow down, leading to delayed message processing and increased latency.
* Data loss: If the storage capacity is exceeded, NATS may start dropping messages, resulting in data loss and potential business disruptions.
* Instability: High storage usage can lead to instability in the NATS cluster, making it more prone to crashes and failures.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `gnatsd_varz_jetstream_stats_storage` metric to determine the current storage usage of the NATS instance.
2. Verify that the `gnatsd_varz_jetstream_config_max_storage` metric is set correctly and reflects the intended maximum storage capacity.
3. Review the NATS instance configuration and JetStream settings to ensure they are optimized for storage efficiency.
4. Investigate recent changes to the NATS instance or JetStream settings that may have contributed to the high storage usage.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the maximum storage capacity: If possible, increase the `gnatsd_varz_jetstream_config_max_storage` value to provide more storage space for the NATS instance.
2. Optimize JetStream settings: Review and optimize JetStream settings to reduce storage usage, such as adjusting retention policies, message sizes, and compression settings.
3. Purge unnecessary data: Remove any unnecessary data from the JetStream store to free up storage space.
4. Monitor storage usage: Closely monitor storage usage and adjust configurations as needed to prevent future high usage occurrences.

Remember to investigate the root cause of the high storage usage and address it to prevent recurrence.