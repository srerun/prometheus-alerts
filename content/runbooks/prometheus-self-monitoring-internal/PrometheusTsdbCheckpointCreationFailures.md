---
title: PrometheusTsdbCheckpointCreationFailures
description: Troubleshooting for alert PrometheusTsdbCheckpointCreationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbCheckpointCreationFailures

Prometheus encountered {{ $value }} checkpoint creation failures

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTsdbCheckpointCreationFailures" %}}

{{% comment %}}

```yaml
alert: PrometheusTsdbCheckpointCreationFailures
expr: increase(prometheus_tsdb_checkpoint_creations_failed_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus TSDB checkpoint creation failures (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} checkpoint creation failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTsdbCheckpointCreationFailures.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusTsdbCheckpointCreationFailures alert:

## Meaning

The PrometheusTsdbCheckpointCreationFailures alert is triggered when Prometheus encounters failures while creating TSDB checkpoints. TSDB checkpoints are essential for data durability and recovery in Prometheus. A failure to create these checkpoints can lead to data loss or corruption.

## Impact

The impact of this alert is critical, as it can result in:

* Data loss or corruption
* Incomplete or inaccurate metrics data
* Downtime or slow performance of Prometheus and dependent services
* Inability to recover data in case of a failure or restart

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus logs for errors related to TSDB checkpoint creation.
2. Verify that the disk space is sufficient and not running out of disk space.
3. Check the underlying storage system for any issues or errors.
4. Investigate any recent changes to the Prometheus configuration or deployment.
5. Check the value of `prometheus_tsdb_checkpoint_creations_failed_total` metric to understand the extent of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Prometheus configuration and ensure that the `storage.tsdb.retention=time` setting is correct and not too aggressive.
2. Increase the disk space available to Prometheus if it's running low.
3. Investigate and resolve any underlying storage system issues.
4. Restart the Prometheus service to attempt to recreate the checkpoints.
5. If the issue persists, consider rolling back recent changes to the Prometheus configuration or deployment.
6. If none of the above steps resolve the issue, seek assistance from a Prometheus expert or the development team.