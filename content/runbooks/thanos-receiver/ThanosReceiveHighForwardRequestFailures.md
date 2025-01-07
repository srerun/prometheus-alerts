---
title: ThanosReceiveHighForwardRequestFailures
description: Troubleshooting for alert ThanosReceiveHighForwardRequestFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveHighForwardRequestFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Receive {{$labels.job}} is failing to forward {{$value | humanize}}% of requests.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosReceiveHighForwardRequestFailures
expr: (sum by (job) (rate(thanos_receive_forward_requests_total{result="error", job=~".*thanos-receive.*"}[5m]))/  sum by (job) (rate(thanos_receive_forward_requests_total{job=~".*thanos-receive.*"}[5m]))) * 100 > 20
for: 5m
labels:
    severity: info
annotations:
    summary: Thanos Receive High Forward Request Failures (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.job}} is failing to forward {{$value | humanize}}% of requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosReceiveHighForwardRequestFailures

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
