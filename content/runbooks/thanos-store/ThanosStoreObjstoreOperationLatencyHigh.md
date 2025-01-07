---
title: ThanosStoreObjstoreOperationLatencyHigh
description: Troubleshooting for alert ThanosStoreObjstoreOperationLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosStoreObjstoreOperationLatencyHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Store {{$labels.job}} Bucket has a 99th percentile latency of {{$value}} seconds for the bucket operations.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosStoreObjstoreOperationLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(thanos_objstore_bucket_operation_duration_seconds_bucket{job=~".*thanos-store.*"}[5m]))) > 2 and  sum by (job) (rate(thanos_objstore_bucket_operation_duration_seconds_count{job=~".*thanos-store.*"}[5m])) > 0)
for: 10m
labels:
    severity: warning
annotations:
    summary: Thanos Store Objstore Operation Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Store {{$labels.job}} Bucket has a 99th percentile latency of {{$value}} seconds for the bucket operations.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosStoreObjstoreOperationLatencyHigh

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
