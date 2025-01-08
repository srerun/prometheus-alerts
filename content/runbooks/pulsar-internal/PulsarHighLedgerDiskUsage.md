---
title: PulsarHighLedgerDiskUsage
description: Troubleshooting for alert PulsarHighLedgerDiskUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PulsarHighLedgerDiskUsage

Observing Ledger Disk Usage (> 75%)

<details>
  <summary>Alert Rule</summary>

{{% rule "pulsar/pulsar-internal.yml" "PulsarHighLedgerDiskUsage" %}}

{{% comment %}}

```yaml
alert: PulsarHighLedgerDiskUsage
expr: sum(bookie_ledger_dir__pulsar_data_bookkeeper_ledgers_usage) by (kubernetes_pod_name) > 75
for: 1h
labels:
    severity: critical
annotations:
    summary: Pulsar high ledger disk usage (instance {{ $labels.instance }})
    description: |-
        Observing Ledger Disk Usage (> 75%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pulsar-internal/PulsarHighLedgerDiskUsage.md

```

{{% /comment %}}

</details>


## Meaning

The PulsarHighLedgerDiskUsage alert is triggered when the sum of ledger disk usage for all bookkeepers in a Pulsar cluster exceeds 75% for more than 1 hour. This indicates that the available disk space for ledgers is critically low, which can lead to performance issues, data loss, or even cluster unavailability.

## Impact

* Performance degradation: High ledger disk usage can cause slow write speeds, leading to increased latency and decreased throughput.
* Data loss: If the disk becomes full, new data may not be written, resulting in data loss.
* Cluster unavailability: In extreme cases, the cluster may become unavailable due to the lack of disk space, causing a complete outage.

## Diagnosis

To diagnose the root cause of the high ledger disk usage, follow these steps:

1. Check the Prometheus metrics: Verify the `bookie_ledger_dir__pulsar_data_bookkeeper_ledgers_usage` metric to confirm the high disk usage.
2. Identify the affected pod: Look for the `kubernetes_pod_name` label to determine which pod is experiencing high ledger disk usage.
3. Investigate recent changes: Check the cluster's deployment history and recent changes to the Pulsar configuration to see if any recent updates may have contributed to the high disk usage.
4. Check ledger retention policies: Review the ledger retention policies to ensure they are not too aggressive, causing ledgers to grow rapidly.
5. Verify disk capacity: Check the disk capacity and available space to ensure it is sufficient for the current workload.

## Mitigation

To mitigate the high ledger disk usage, follow these steps:

1. Increase disk capacity: Add more disk space to the affected pod or cluster to alleviate the disk usage pressure.
2. Adjust ledger retention policies: Modify the ledger retention policies to reduce the rate of ledger growth or increase the frequency of ledger compaction.
3. Optimize Pulsar configuration: Review and optimize the Pulsar configuration to reduce disk usage, such as adjusting the `managedLedgerDefaultEnsembleSize` or `managedLedgerMaxEntriesPerLedger` settings.
4. Implement disk usage monitoring: Set up monitoring to track disk usage and receive early warnings of potential issues.
5. Consider load balancing: If the high disk usage is due to uneven load distribution, consider load balancing the workload across multiple pods or clusters.

Remember to follow the links in the alert annotations for more detailed instructions and troubleshooting steps.