---
title: HaproxyBackendMaxActiveSession
description: Troubleshooting for alert HaproxyBackendMaxActiveSession
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyBackendMaxActiveSession

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HAproxy backend {{ $labels.fqdn }}/{{ $labels.backend }} is reaching session limit (> 80%).

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyBackendMaxActiveSession
expr: ((sum by (backend) (avg_over_time(haproxy_backend_current_sessions[2m]) * 100) / sum by (backend) (avg_over_time(haproxy_backend_limit_sessions[2m])))) > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: HAProxy backend max active session (instance {{ $labels.instance }})
    description: |-
        HAproxy backend {{ $labels.fqdn }}/{{ $labels.backend }} is reaching session limit (> 80%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyBackendMaxActiveSession

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
