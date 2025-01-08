---
title: NginxLatencyHigh
description: Troubleshooting for alert NginxLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NginxLatencyHigh

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Nginx p99 latency is higher than 3 seconds

<details>
  <summary>Alert Rule</summary>

{{% rule "nginx/knyar-nginx-exporter.yml" "NginxLatencyHigh" %}}

{{% comment %}}

```yaml
alert: NginxLatencyHigh
expr: histogram_quantile(0.99, sum(rate(nginx_http_request_duration_seconds_bucket[2m])) by (host, node, le)) > 3
for: 2m
labels:
    severity: warning
annotations:
    summary: Nginx latency high (instance {{ $labels.instance }})
    description: |-
        Nginx p99 latency is higher than 3 seconds
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/knyar-nginx-exporter/NginxLatencyHigh.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
