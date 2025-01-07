---
title: NginxHighHttp4xxErrorRate
description: Troubleshooting for alert NginxHighHttp4xxErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NginxHighHttp4xxErrorRate

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many HTTP requests with status 4xx (> 5%)

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: NginxHighHttp4xxErrorRate
expr: sum(rate(nginx_http_requests_total{status=~"^4.."}[1m])) / sum(rate(nginx_http_requests_total[1m])) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: Nginx high HTTP 4xx error rate (instance {{ $labels.instance }})
    description: |-
        Too many HTTP requests with status 4xx (> 5%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/NginxHighHttp4xxErrorRate

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
