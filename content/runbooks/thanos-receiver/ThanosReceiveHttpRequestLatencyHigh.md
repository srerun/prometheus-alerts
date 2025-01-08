---
title: ThanosReceiveHttpRequestLatencyHigh
description: Troubleshooting for alert ThanosReceiveHttpRequestLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHttpRequestLatencyHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Receive {{$labels.job}} has a 99th percentile latency of {{ $value }} seconds for requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveHttpRequestLatencyHigh" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveHttpRequestLatencyHigh
expr: (histogram_quantile(0.99, sum by (job, le) (rate(http_request_duration_seconds_bucket{job=~".*thanos-receive.*", handler="receive"}[5m]))) > 10 and sum by (job) (rate(http_request_duration_seconds_count{job=~".*thanos-receive.*", handler="receive"}[5m])) > 0)
for: 10m
labels:
    severity: critical
annotations:
    summary: Thanos Receive Http Request Latency High (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} has a 99th percentile latency of {{ $value }} seconds for requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveHttpRequestLatencyHigh.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
