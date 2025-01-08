---
title: ThanosStoreSeriesGateLatencyHigh
description: Troubleshooting for alert ThanosStoreSeriesGateLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosStoreSeriesGateLatencyHigh

Thanos Store {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for store series gate requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-store.yml" "ThanosStoreSeriesGateLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosStoreSeriesGateLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(thanos_bucket_store_series_gate_duration_seconds_bucket{job=~".*thanos-store.*"}[5m]))) > 2 and sum by (job) (rate(thanos_bucket_store_series_gate_duration_seconds_count{job=~".*thanos-store.*"}[5m])) > 0)
for: 10m
labels:
    severity: warning
annotations:
    summary: Thanos Store Series Gate Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Store {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for store series gate requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-store/ThanosStoreSeriesGateLatencyHigh.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
