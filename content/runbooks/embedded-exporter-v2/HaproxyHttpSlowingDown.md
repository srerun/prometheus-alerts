---
title: HaproxyHttpSlowingDown
description: Troubleshooting for alert HaproxyHttpSlowingDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHttpSlowingDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Average request time is increasing - {{ $value | printf "%.2f"}}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyHttpSlowingDown
expr: avg by (instance, proxy) (haproxy_backend_max_total_time_seconds) > 1
for: 1m
labels:
    severity: warning
annotations:
    summary: HAProxy HTTP slowing down (instance {{ $labels.instance }})
    description: |-
        Average request time is increasing - {{ $value | printf "%.2f"}}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyHttpSlowingDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
