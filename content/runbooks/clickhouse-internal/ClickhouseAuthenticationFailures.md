---
title: ClickhouseAuthenticationFailures
description: Troubleshooting for alert ClickhouseAuthenticationFailures
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseAuthenticationFailures

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Authentication failures detected, indicating potential security issues or misconfiguration.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ClickhouseAuthenticationFailures
expr: increase(ClickHouseErrorMetric_AUTHENTICATION_FAILED[5m]) > 0
for: 0m
labels:
    severity: info
annotations:
    summary: ClickHouse Authentication Failures (instance {{ $labels.instance }})
    description: |-
        Authentication failures detected, indicating potential security issues or misconfiguration.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ClickhouseAuthenticationFailures

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
