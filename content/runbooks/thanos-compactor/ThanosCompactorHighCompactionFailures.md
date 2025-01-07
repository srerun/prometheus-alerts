---
title: ThanosCompactorHighCompactionFailures
description: Troubleshooting for alert ThanosCompactorHighCompactionFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosCompactorHighCompactionFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Compact {{$labels.job}} is failing to execute {{$value | humanize}}% of compactions.

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosCompactorHighCompactionFailures

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
