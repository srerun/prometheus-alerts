---
title: ThanosQueryGrpcClientErrorRate
description: Troubleshooting for alert ThanosQueryGrpcClientErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryGrpcClientErrorRate

Thanos Query {{$labels.job}} is failing to send {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryGrpcClientErrorRate" %}}

{{% comment %}}

```yaml
alert: ThanosQueryGrpcClientErrorRate
expr: (sum by (job) (rate(grpc_client_handled_total{grpc_code!="OK", job=~".*thanos-query.*"}[5m])) / sum by (job) (rate(grpc_client_started_total{job=~".*thanos-query.*"}[5m]))) * 100 > 5
for: 5m
labels:
    severity: warning
annotations:
    summary: Thanos Query Grpc Client Error Rate (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} is failing to send {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryGrpcClientErrorRate.md

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
