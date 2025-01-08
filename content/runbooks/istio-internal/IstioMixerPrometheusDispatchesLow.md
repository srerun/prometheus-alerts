---
title: IstioMixerPrometheusDispatchesLow
description: Troubleshooting for alert IstioMixerPrometheusDispatchesLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioMixerPrometheusDispatchesLow

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Number of Mixer dispatches to Prometheus is too low. Istio metrics might not be being exported properly.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioMixerPrometheusDispatchesLow" %}}

<!-- Rule when generated

```yaml
alert: IstioMixerPrometheusDispatchesLow
expr: sum(rate(mixer_runtime_dispatches_total{adapter=~"prometheus"}[1m])) < 180
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio Mixer Prometheus dispatches low (instance {{ $labels.instance }})
    description: |-
        Number of Mixer dispatches to Prometheus is too low. Istio metrics might not be being exported properly.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioMixerPrometheusDispatchesLow.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
