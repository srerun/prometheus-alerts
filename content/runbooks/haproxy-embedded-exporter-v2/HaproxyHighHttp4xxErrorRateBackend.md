---
title: HaproxyHighHttp4xxErrorRateBackend
description: Troubleshooting for alert HaproxyHighHttp4xxErrorRateBackend
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHighHttp4xxErrorRateBackend

Too many HTTP requests with status 4xx (> 5%) on backend {{ $labels.fqdn }}/{{ $labels.backend }}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-embedded-exporter-v2.yml" "HaproxyHighHttp4xxErrorRateBackend" %}}

{{% comment %}}

```yaml
alert: HaproxyHighHttp4xxErrorRateBackend
expr: ((sum by (proxy) (rate(haproxy_server_http_responses_total{code="4xx"}[1m])) / sum by (proxy) (rate(haproxy_server_http_responses_total[1m]))) * 100) > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy high HTTP 4xx error rate backend (instance {{ $labels.instance }})
    description: |-
        Too many HTTP requests with status 4xx (> 5%) on backend {{ $labels.fqdn }}/{{ $labels.backend }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/HaproxyHighHttp4xxErrorRateBackend.md

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
