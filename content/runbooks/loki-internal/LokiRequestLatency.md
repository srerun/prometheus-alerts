---
title: LokiRequestLatency
description: Troubleshooting for alert LokiRequestLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiRequestLatency

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}s 99th percentile latency

<details>
  <summary>Alert Rule</summary>

{{% rule "loki/loki-internal.yml" "LokiRequestLatency" %}}

<!-- Rule when generated

```yaml
alert: LokiRequestLatency
expr: (histogram_quantile(0.99, sum(rate(loki_request_duration_seconds_bucket{route!~"(?i).*tail.*"}[5m])) by (le)))  > 1
for: 5m
labels:
    severity: critical
annotations:
    summary: Loki request latency (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}s 99th percentile latency
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/loki-internal/LokiRequestLatency.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
