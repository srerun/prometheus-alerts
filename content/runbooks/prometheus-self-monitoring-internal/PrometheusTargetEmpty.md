---
title: PrometheusTargetEmpty
description: Troubleshooting for alert PrometheusTargetEmpty
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTargetEmpty

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Prometheus has no target in service discovery

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusTargetEmpty
expr: prometheus_sd_discovered_targets == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus target empty (instance {{ $labels.instance }})
    description: |-
        Prometheus has no target in service discovery
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PrometheusTargetEmpty

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
