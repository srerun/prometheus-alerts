---
title: IstioLatency99Percentile
description: Troubleshooting for alert IstioLatency99Percentile
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioLatency99Percentile

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Istio 1% slowest requests are longer than 1000ms.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioLatency99Percentile" %}}

<!-- Rule when generated

```yaml
alert: IstioLatency99Percentile
expr: histogram_quantile(0.99, sum(rate(istio_request_duration_milliseconds_bucket[1m])) by (destination_canonical_service, destination_workload_namespace, source_canonical_service, source_workload_namespace, le)) > 1000
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio latency 99 percentile (instance {{ $labels.instance }})
    description: |-
        Istio 1% slowest requests are longer than 1000ms.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioLatency99Percentile.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
