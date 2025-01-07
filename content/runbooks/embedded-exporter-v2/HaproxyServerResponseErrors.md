---
title: HaproxyServerResponseErrors
description: Troubleshooting for alert HaproxyServerResponseErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyServerResponseErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many response errors to {{ $labels.server }} server (> 5%).

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyServerResponseErrors
expr: (sum by (server) (rate(haproxy_server_response_errors_total[1m])) / sum by (server) (rate(haproxy_server_http_responses_total[1m]))) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy server response errors (instance {{ $labels.instance }})
    description: |-
        Too many response errors to {{ $labels.server }} server (> 5%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyServerResponseErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
