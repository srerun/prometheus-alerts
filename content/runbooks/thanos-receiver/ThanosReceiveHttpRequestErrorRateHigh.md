---
title: ThanosReceiveHttpRequestErrorRateHigh
description: Troubleshooting for alert ThanosReceiveHttpRequestErrorRateHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHttpRequestErrorRateHigh

Thanos Receive {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveHttpRequestErrorRateHigh" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveHttpRequestErrorRateHigh
expr: (sum by (job) (rate(http_requests_total{code=~"5..", job=~".*thanos-receive.*", handler="receive"}[5m]))/  sum by (job) (rate(http_requests_total{job=~".*thanos-receive.*", handler="receive"}[5m]))) * 100 > 5
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Receive Http Request Error Rate High (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} is failing to handle {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveHttpRequestErrorRateHigh.md

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
