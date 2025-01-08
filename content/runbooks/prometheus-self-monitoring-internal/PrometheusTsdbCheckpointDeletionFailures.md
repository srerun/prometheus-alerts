---
title: PrometheusTsdbCheckpointDeletionFailures
description: Troubleshooting for alert PrometheusTsdbCheckpointDeletionFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbCheckpointDeletionFailures

Prometheus encountered {{ $value }} checkpoint deletion failures

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTsdbCheckpointDeletionFailures" %}}

{{% comment %}}

```yaml
alert: PrometheusTsdbCheckpointDeletionFailures
expr: increase(prometheus_tsdb_checkpoint_deletions_failed_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus TSDB checkpoint deletion failures (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} checkpoint deletion failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTsdbCheckpointDeletionFailures.md

```

{{% /comment %}}

</details>


## Meaning

The PrometheusTsdbCheckpointDeletionFailures alert is triggered when Prometheus fails to delete checkpoint files from its TSDB (Time Series Database) storage. Checkpoint files are used to store the current state of Prometheus's internal data structures, such as the in-memory index and the WAL (Write-Ahead Log). The deletion of these files is crucial to prevent disk space exhaustion and maintain a healthy Prometheus instance.

## Impact

If this alert is not addressed, it can lead to:

* Disk space exhaustion, causing Prometheus to run out of storage capacity, leading to performance issues, and potentially even crashes.
* Inconsistent data, as the failed checkpoint deletion can cause inconsistencies in the TSDB, leading to incorrect query results or even data loss.
* Increased memory usage, as the system may retain unnecessary data in memory, further exacerbating performance issues.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus server's disk usage and available storage capacity.
2. Verify that the TSDB is functioning correctly and that the WAL is being properly truncated.
3. Review the Prometheus server logs for any errors or warnings related to checkpoint deletion failures.
4. Check the $labels.instance label to identify the specific Prometheus instance affected.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately free up disk space by deleting unnecessary files and data.
2. Restart the Prometheus server to allow it to recover and retry the checkpoint deletion.
3. Verify that the TSDB is functioning correctly and that the WAL is being properly truncated.
4. Implement a regular maintenance schedule to ensure that disk space is regularly cleaned up and that the TSDB is healthy.
5. Consider increasing the storage capacity or implementing a more robust storage solution to prevent future disk space exhaustion issues.