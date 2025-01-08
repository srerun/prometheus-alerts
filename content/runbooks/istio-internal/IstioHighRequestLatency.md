---
title: IstioHighRequestLatency
description: Troubleshooting for alert IstioHighRequestLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioHighRequestLatency

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Istio average requests execution is longer than 100ms.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioHighRequestLatency" %}}

{{% comment %}}

```yaml
alert: IstioHighRequestLatency
expr: rate(istio_request_duration_milliseconds_sum{reporter="destination"}[1m]) / rate(istio_request_duration_milliseconds_count{reporter="destination"}[1m]) > 100
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio high request latency (instance {{ $labels.instance }})
    description: |-
        Istio average requests execution is longer than 100ms.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioHighRequestLatency.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
