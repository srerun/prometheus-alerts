---
title: HaproxyServerDown
description: Troubleshooting for alert HaproxyServerDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyServerDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HAProxy server is down

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyServerDown
expr: haproxy_server_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: HAProxy server down (instance {{ $labels.instance }})
    description: |-
        HAProxy server is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyServerDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
