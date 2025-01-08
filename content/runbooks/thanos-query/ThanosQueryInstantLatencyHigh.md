---
title: ThanosQueryInstantLatencyHigh
description: Troubleshooting for alert ThanosQueryInstantLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryInstantLatencyHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Query {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for instant queries.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryInstantLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosQueryInstantLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(http_request_duration_seconds_bucket{job=~".*thanos-query.*", handler="query"}[5m]))) > 40 and sum by (job) (rate(http_request_duration_seconds_bucket{job=~".*thanos-query.*", handler="query"}[5m])) > 0)
for: 10m
labels:
    severity: critical
annotations:
    summary: Thanos Query Instant Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} has a 99th percentile latency of {{$value}} seconds for instant queries.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryInstantLatencyHigh.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
