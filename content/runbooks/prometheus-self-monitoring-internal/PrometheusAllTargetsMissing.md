---
title: PrometheusAllTargetsMissing
description: Troubleshooting for alert PrometheusAllTargetsMissing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAllTargetsMissing

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A Prometheus job does not have living target anymore.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusAllTargetsMissing
expr: sum by (job) (up) == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus all targets missing (instance {{ $labels.instance }})
    description: |-
        A Prometheus job does not have living target anymore.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PrometheusAllTargetsMissing

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
