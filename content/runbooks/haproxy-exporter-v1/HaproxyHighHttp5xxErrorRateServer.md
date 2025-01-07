---
title: HaproxyHighHttp5xxErrorRateServer
description: Troubleshooting for alert HaproxyHighHttp5xxErrorRateServer
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHighHttp5xxErrorRateServer

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many HTTP requests with status 5xx (> 5%) on server {{ $labels.server }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyHighHttp5xxErrorRateServer
expr: sum by (server) (rate(haproxy_server_http_responses_total{code="5xx"}[1m]) * 100) / sum by (server) (rate(haproxy_server_http_responses_total[1m])) > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy high HTTP 5xx error rate server (instance {{ $labels.instance }})
    description: |-
        Too many HTTP requests with status 5xx (> 5%) on server {{ $labels.server }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyHighHttp5xxErrorRateServer

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
