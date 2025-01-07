---
title: TraefikHighHttp5xxErrorRateService
description: Troubleshooting for alert TraefikHighHttp5xxErrorRateService
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikHighHttp5xxErrorRateService

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Traefik service 5xx error rate is above 5%

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: TraefikHighHttp5xxErrorRateService
expr: sum(rate(traefik_service_requests_total{code=~"5.*"}[3m])) by (service) / sum(rate(traefik_service_requests_total[3m])) by (service) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: Traefik high HTTP 5xx error rate service (instance {{ $labels.instance }})
    description: |-
        Traefik service 5xx error rate is above 5%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/TraefikHighHttp5xxErrorRateService

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
