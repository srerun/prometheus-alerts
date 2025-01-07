---
title: ClickhouseAccessDeniedErrors
description: Troubleshooting for alert ClickhouseAccessDeniedErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseAccessDeniedErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Access denied errors have been logged, which could indicate permission issues or unauthorized access attempts.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ClickhouseAccessDeniedErrors
expr: increase(ClickHouseErrorMetric_RESOURCE_ACCESS_DENIED[5m]) > 0
for: 0m
labels:
    severity: info
annotations:
    summary: ClickHouse Access Denied Errors (instance {{ $labels.instance }})
    description: |-
        Access denied errors have been logged, which could indicate permission issues or unauthorized access attempts.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ClickhouseAccessDeniedErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
