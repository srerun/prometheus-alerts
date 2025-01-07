---
title: HaproxyHasNoAliveBackends
description: Troubleshooting for alert HaproxyHasNoAliveBackends
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHasNoAliveBackends

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HAProxy has no alive active or backup backends for {{ $labels.proxy }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyHasNoAliveBackends
expr: haproxy_backend_active_servers + haproxy_backend_backup_servers == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: HAproxy has no alive backends (instance {{ $labels.instance }})
    description: |-
        HAProxy has no alive active or backup backends for {{ $labels.proxy }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyHasNoAliveBackends

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
