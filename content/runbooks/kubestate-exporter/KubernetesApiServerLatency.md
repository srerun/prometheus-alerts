---
title: KubernetesApiServerLatency
description: Troubleshooting for alert KubernetesApiServerLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesApiServerLatency

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Kubernetes API server has a 99th percentile latency of {{ $value }} seconds for {{ $labels.verb }} {{ $labels.resource }}.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: KubernetesApiServerLatency
expr: histogram_quantile(0.99, sum(rate(apiserver_request_duration_seconds_bucket{verb!~"(?:CONNECT|WATCHLIST|WATCH|PROXY)"} [10m])) WITHOUT (subresource)) > 1
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes API server latency (instance {{ $labels.instance }})
    description: |-
        Kubernetes API server has a 99th percentile latency of {{ $value }} seconds for {{ $labels.verb }} {{ $labels.resource }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/KubernetesApiServerLatency

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
