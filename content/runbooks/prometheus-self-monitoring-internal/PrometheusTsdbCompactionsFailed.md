---
title: PrometheusTsdbCompactionsFailed
description: Troubleshooting for alert PrometheusTsdbCompactionsFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTsdbCompactionsFailed

Prometheus encountered {{ $value }} TSDB compactions failures

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusTsdbCompactionsFailed" %}}

{{% comment %}}

```yaml
alert: PrometheusTsdbCompactionsFailed
expr: increase(prometheus_tsdb_compactions_failed_total[1m]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus TSDB compactions failed (instance {{ $labels.instance }})
    description: |-
        Prometheus encountered {{ $value }} TSDB compactions failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusTsdbCompactionsFailed.md

```

{{% /comment %}}

</details>


Here is a runbook for the PrometheusTsdbCompactionsFailed alert rule:

## Meaning

The PrometheusTsdbCompactionsFailed alert rule is triggered when Prometheus encounters failures while compacting its time series database (TSDB). TSDB compactions are a critical maintenance task that ensures efficient storage and query performance. Failure to compact the TSDB can lead to performance degradation, increased storage usage, and even crashes.

## Impact

The impact of this alert is critical, as it can cause:

* Performance degradation: Uncompacted TSDB can lead to slower query responses and increased latency.
* Storage usage increase: Uncompacted data can occupy more disk space, leading to storage capacity issues.
* Instability: In extreme cases, TSDB compaction failures can cause Prometheus to crash or become unresponsive.

## Diagnosis

To diagnose the root cause of the TSDB compaction failures, follow these steps:

1. Check the Prometheus logs for error messages related to TSDB compactions.
2. Verify that the disk space is sufficient, and the file system is not full.
3. Check for any recent changes or updates to the Prometheus configuration or environment.
4. Review the TSDB compaction metrics, such as `prometheus_tsdb_compactions_failed_total`, to identify any patterns or trends.

## Mitigation

To mitigate the TSDB compaction failures, follow these steps:

1. Check the Prometheus configuration to ensure that the TSDB compaction settings are correct and sufficient for the current data volume.
2. Verify that the Prometheus instance has sufficient resources (CPU, memory, and disk space) to perform compactions efficiently.
3. Consider increasing the `storage.local.target-avail-bytes` configuration option to allow for more aggressive compaction.
4. If the issue persists, consider restarting the Prometheus instance to recover from any potential internal state issues.
5. If none of the above steps resolve the issue, consider seeking assistance from the Prometheus community or a qualified administrator.

Remember to update the Prometheus configuration and environment according to the findings and resolutions to prevent future TSDB compaction failures.