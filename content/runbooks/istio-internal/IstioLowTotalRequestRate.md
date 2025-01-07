---
title: IstioLowTotalRequestRate
description: Troubleshooting for alert IstioLowTotalRequestRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioLowTotalRequestRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Global request rate in the service mesh is unusually low.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: IstioLowTotalRequestRate
expr: sum(rate(istio_requests_total{reporter="destination"}[5m])) < 100
for: 2m
labels:
    severity: warning
annotations:
    summary: Istio low total request rate (instance {{ $labels.instance }})
    description: |-
        Global request rate in the service mesh is unusually low.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/IstioLowTotalRequestRate

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
