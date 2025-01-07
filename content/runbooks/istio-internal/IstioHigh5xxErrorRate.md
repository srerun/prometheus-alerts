---
title: IstioHigh5xxErrorRate
description: Troubleshooting for alert IstioHigh5xxErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioHigh5xxErrorRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High percentage of HTTP 5xx responses in Istio (> 5%).

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: IstioHigh5xxErrorRate
expr: sum(rate(istio_requests_total{reporter="destination", response_code=~"5.*"}[5m])) / sum(rate(istio_requests_total{reporter="destination"}[5m])) * 100 > 5
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio high 5xx error rate (instance {{ $labels.instance }})
    description: |-
        High percentage of HTTP 5xx responses in Istio (> 5%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/IstioHigh5xxErrorRate

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
