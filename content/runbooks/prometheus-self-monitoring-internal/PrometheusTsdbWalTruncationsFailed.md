---
title: PrometheusTsdbWalTruncationsFailed
description: Troubleshooting for alert PrometheusTsdbWalTruncationsFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbWalTruncationsFailed

Prometheus encountered {{ $value }} TSDB WAL truncation failures

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTsdbWalTruncationsFailed" %}}

{{% comment %}}

```yaml
alert: PrometheusTsdbWalTruncationsFailed
expr: increase(prometheus_tsdb_wal_truncations_failed_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus TSDB WAL truncations failed (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} TSDB WAL truncation failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTsdbWalTruncationsFailed.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusTsdbWalTruncationsFailed alert rule:

## Meaning

The PrometheusTsdbWalTruncationsFailed alert is triggered when Prometheus encounters failures during the truncation of its Write-Ahead Log (WAL). The WAL is a critical component of Prometheus's storage system, and failures during truncation can lead to data loss or inconsistencies.

## Impact

If this alert is triggered, it may indicate that Prometheus is experiencing issues with its storage system, which can lead to:

* Data loss or inconsistencies
* Inaccurate or incomplete metrics
* Performance degradation
* Increased risk of data corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus logs for errors related to WAL truncation.
2. Verify that the disk space is sufficient and not running out.
3. Check the disk I/O performance and verify that it's within acceptable limits.
4. Verify that the WAL directory is not corrupted and is writable.
5. Check the Prometheus configuration to ensure that it's correctly configured for WAL truncation.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the disk space and ensure that it's sufficient.
2. Restart the Prometheus instance to see if it resolves the issue.
3. Verify that the WAL directory is not corrupted and is writable.
4. Check the disk I/O performance and optimize it if necessary.
5. If the issue persists, consider increasing the `storage.tsdb.retention` configuration option to reduce the frequency of WAL truncations.
6. Consider increasing the `storage.tsdb.wal-compression` configuration option to reduce the WAL size.
7. If none of the above steps resolve the issue, consider seeking assistance from a Prometheus administrator or a qualified engineer.