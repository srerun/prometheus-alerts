---
title: PrometheusTsdbWalCorruptions
description: Troubleshooting for alert PrometheusTsdbWalCorruptions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbWalCorruptions

Prometheus encountered {{ $value }} TSDB WAL corruptions

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTsdbWalCorruptions" %}}

{{% comment %}}

```yaml
alert: PrometheusTsdbWalCorruptions
expr: increase(prometheus_tsdb_wal_corruptions_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus TSDB WAL corruptions (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} TSDB WAL corruptions
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTsdbWalCorruptions.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the PrometheusTsdbWalCorruptions alert:

## Meaning

The PrometheusTsdbWalCorruptions alert is triggered when Prometheus encounters corruptions in its Write-Ahead Log (WAL) storage, which is used to store incoming samples temporarily before they are written to the long-term storage. This corruption can lead to data loss and inconsistencies in the Prometheus storage.

## Impact

The impact of this alert is critical, as it can result in:

* Data loss: Corrupted WAL data can lead to the loss of metric data, which can impact the reliability of monitoring and alerting.
* Inconsistent data: Corrupted WAL data can also lead to inconsistencies in the Prometheus storage, making it difficult to trust the accuracy of the data.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus server logs for errors related to WAL corruptions.
2. Verify that the WAL storage is properly configured and has sufficient disk space.
3. Check the Prometheus instance's disk usage and ensure that it has enough available disk space.
4. Run the `promtool tsdb` command to inspect the WAL and identify any corrupted segments.
5. Check the `prometheus_tsdb_wal_corruptions_total` metric to see the total number of corruptions encountered.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately stop the Prometheus instance to prevent further data corruption.
2. Identify and fix the underlying cause of the corruption, such as disk space issues or configuration errors.
3. Run the `promtool tsdb` command to repair the corrupted WAL segments.
4. Restart the Prometheus instance and verify that it is functioning correctly.
5. Monitor the `prometheus_tsdb_wal_corruptions_total` metric to ensure that no further corruptions occur.

Note: It is essential to act quickly to mitigate this issue to prevent further data loss and inconsistencies.