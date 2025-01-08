---
title: IstioHighTotalRequestRate
description: Troubleshooting for alert IstioHighTotalRequestRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioHighTotalRequestRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Global request rate in the service mesh is unusually high.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioHighTotalRequestRate" %}}

{{% comment %}}

```yaml
alert: IstioHighTotalRequestRate
expr: sum(rate(istio_requests_total{reporter="destination"}[5m])) > 1000
for: 2m
labels:
    severity: warning
annotations:
    summary: Istio high total request rate (instance {{ $labels.instance }})
    description: |-
        Global request rate in the service mesh is unusually high.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioHighTotalRequestRate.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
