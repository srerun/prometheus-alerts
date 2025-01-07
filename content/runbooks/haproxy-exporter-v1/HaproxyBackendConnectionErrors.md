---
title: HaproxyBackendConnectionErrors
description: Troubleshooting for alert HaproxyBackendConnectionErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyBackendConnectionErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Too many connection errors to {{ $labels.fqdn }}/{{ $labels.backend }} backend (> 100 req/s). Request throughput may be too high.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyBackendConnectionErrors
expr: sum by (backend) (rate(haproxy_backend_connection_errors_total[1m])) > 100
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy backend connection errors (instance {{ $labels.instance }})
    description: |-
        Too many connection errors to {{ $labels.fqdn }}/{{ $labels.backend }} backend (> 100 req/s). Request throughput may be too high.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyBackendConnectionErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
