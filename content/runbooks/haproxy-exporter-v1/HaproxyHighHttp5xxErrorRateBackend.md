---
title: HaproxyHighHttp5xxErrorRateBackend
description: Troubleshooting for alert HaproxyHighHttp5xxErrorRateBackend
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHighHttp5xxErrorRateBackend

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many HTTP requests with status 5xx (> 5%) on backend {{ $labels.fqdn }}/{{ $labels.backend }}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyHighHttp5xxErrorRateBackend" %}}

<!-- Rule when generated

```yaml
alert: HaproxyHighHttp5xxErrorRateBackend
expr: sum by (backend) (rate(haproxy_server_http_responses_total{code="5xx"}[1m])) / sum by (backend) (rate(haproxy_server_http_responses_total[1m])) > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy high HTTP 5xx error rate backend (instance {{ $labels.instance }})
    description: |-
        Too many HTTP requests with status 5xx (> 5%) on backend {{ $labels.fqdn }}/{{ $labels.backend }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyHighHttp5xxErrorRateBackend.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
