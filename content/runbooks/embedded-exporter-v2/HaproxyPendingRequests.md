---
title: HaproxyPendingRequests
description: Troubleshooting for alert HaproxyPendingRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyPendingRequests

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some HAProxy requests are pending on {{ $labels.proxy }} - {{ $value | printf "%.2f"}}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyPendingRequests
expr: sum by (proxy) (rate(haproxy_backend_current_queue[2m])) > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: HAProxy pending requests (instance {{ $labels.instance }})
    description: |-
        Some HAProxy requests are pending on {{ $labels.proxy }} - {{ $value | printf "%.2f"}}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyPendingRequests

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
