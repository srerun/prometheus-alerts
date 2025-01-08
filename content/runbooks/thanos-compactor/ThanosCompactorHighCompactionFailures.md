---
title: ThanosCompactorHighCompactionFailures
description: Troubleshooting for alert ThanosCompactorHighCompactionFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactorHighCompactionFailures

Thanos Compact {{$labels.job}} is failing to execute {{$value | humanize}}% of compactions.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-compactor.yml" "ThanosCompactorHighCompactionFailures" %}}

{{% comment %}}

```yaml
alert: ThanosCompactorHighCompactionFailures
expr: (sum by (job) (rate(thanos_compact_group_compactions_failures_total{job=~".*thanos-compact.*"}[5m])) / sum by (job) (rate(thanos_compact_group_compactions_total{job=~".*thanos-compact.*"}[5m])) * 100 > 5)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Compactor High Compaction Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Compact {{$labels.job}} is failing to execute {{$value | humanize}}% of compactions.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-compactor/ThanosCompactorHighCompactionFailures.md

```

{{% /comment %}}

</details>


## Meaning

The ThanosCompactorHighCompactionFailures alert is triggered when the percentage of failed compactions in a Thanos compactor exceeds 5% over a 5-minute period. This alert indicates that the compactor is experiencing issues with executing compactions, which can lead to data inconsistencies and affect the overall performance of the system.

## Impact

The impact of this alert includes:

* Data inconsistencies: Failed compactions can result in incomplete or corrupted data, leading to inaccuracies in metrics and dashboards.
* Performance degradation: High compaction failure rates can cause the compactor to slow down, leading to increased latency and decreased system performance.
* Alert fatigue: Frequent or persistent compaction failures can lead to alert fatigue, causing teams to ignore or overlook critical issues.

## Diagnosis

To diagnose the root cause of the ThanosCompactorHighCompactionFailures alert, follow these steps:

1. Check the Thanos compactor logs for errors or exceptions related to compaction failures.
2. Verify that the compactor has sufficient resources (e.g., CPU, memory, and disk space) to execute compactions successfully.
3. Investigate any recent changes to the compactor configuration or deployment that may be contributing to the failures.
4. Check the status of the underlying storage system to ensure it is healthy and responsive.

## Mitigation

To mitigate the ThanosCompactorHighCompactionFailures alert, follow these steps:

1. Investigate and resolve the underlying cause of the compaction failures, as identified during diagnosis.
2. Increase the resources allocated to the compactor to ensure it can handle the compaction workload.
3. Implement retries or fallback mechanisms to handle temporary failures or exceptions during compaction.
4. Consider scaling out the compactor deployment to distribute the compaction workload and improve system resilience.
5. Update the alert threshold or notification settings to reduce alert fatigue and ensure that critical issues are addressed promptly.