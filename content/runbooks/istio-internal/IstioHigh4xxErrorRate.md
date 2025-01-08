---
title: IstioHigh4xxErrorRate
description: Troubleshooting for alert IstioHigh4xxErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioHigh4xxErrorRate

High percentage of HTTP 5xx responses in Istio (> 5%).

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioHigh4xxErrorRate" %}}

{{% comment %}}

```yaml
alert: IstioHigh4xxErrorRate
expr: sum(rate(istio_requests_total{reporter="destination", response_code=~"4.*"}[5m])) / sum(rate(istio_requests_total{reporter="destination"}[5m])) * 100 > 5
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio high 4xx error rate (instance {{ $labels.instance }})
    description: |-
        High percentage of HTTP 5xx responses in Istio (> 5%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioHigh4xxErrorRate.md

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
