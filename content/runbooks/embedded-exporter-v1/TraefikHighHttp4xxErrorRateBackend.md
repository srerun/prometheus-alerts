---
title: TraefikHighHttp4xxErrorRateBackend
description: Troubleshooting for alert TraefikHighHttp4xxErrorRateBackend
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikHighHttp4xxErrorRateBackend

Traefik backend 4xx error rate is above 5%

<details>
  <summary>Alert Rule</summary>

{{% rule "traefik/embedded-exporter-v1.yml" "TraefikHighHttp4xxErrorRateBackend" %}}

{{% comment %}}

```yaml
alert: TraefikHighHttp4xxErrorRateBackend
expr: sum(rate(traefik_backend_requests_total{code=~"4.*"}[3m])) by (backend) / sum(rate(traefik_backend_requests_total[3m])) by (backend) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: Traefik high HTTP 4xx error rate backend (instance {{ $labels.instance }})
    description: |-
        Traefik backend 4xx error rate is above 5%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v1/TraefikHighHttp4xxErrorRateBackend.md

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
