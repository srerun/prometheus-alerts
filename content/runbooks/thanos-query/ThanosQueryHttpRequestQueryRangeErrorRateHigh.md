---
title: ThanosQueryHttpRequestQueryRangeErrorRateHigh
description: Troubleshooting for alert ThanosQueryHttpRequestQueryRangeErrorRateHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryHttpRequestQueryRangeErrorRateHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Query {{$labels.job}} is failing to handle {{$value | humanize}}% of "query_range" requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryHttpRequestQueryRangeErrorRateHigh" %}}

<!-- Rule when generated

```yaml
alert: ThanosQueryHttpRequestQueryRangeErrorRateHigh
expr: (sum by (job) (rate(http_requests_total{code=~"5..", job=~".*thanos-query.*", handler="query_range"}[5m]))/  sum by (job) (rate(http_requests_total{job=~".*thanos-query.*", handler="query_range"}[5m]))) * 100 > 5
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Query Http Request Query Range Error Rate High (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} is failing to handle {{$value | humanize}}% of "query_range" requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryHttpRequestQueryRangeErrorRateHigh.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
