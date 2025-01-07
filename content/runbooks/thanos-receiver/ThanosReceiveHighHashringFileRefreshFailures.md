---
title: ThanosReceiveHighHashringFileRefreshFailures
description: Troubleshooting for alert ThanosReceiveHighHashringFileRefreshFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHighHashringFileRefreshFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Receive {{$labels.job}} is failing to refresh hashring file, {{$value | humanize}} of attempts failed.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosReceiveHighHashringFileRefreshFailures
expr: (sum by (job) (rate(thanos_receive_hashrings_file_errors_total{job=~".*thanos-receive.*"}[5m])) / sum by (job) (rate(thanos_receive_hashrings_file_refreshes_total{job=~".*thanos-receive.*"}[5m])) > 0)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Receive High Hashring File Refresh Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} is failing to refresh hashring file, {{$value | humanize}} of attempts failed.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosReceiveHighHashringFileRefreshFailures

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
