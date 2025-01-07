---
title: PrometheusTargetMissingWithWarmupTime
description: Troubleshooting for alert PrometheusTargetMissingWithWarmupTime
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusTargetMissingWithWarmupTime

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Allow a job time to start up (10 minutes) before alerting that it's down.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PrometheusTargetMissingWithWarmupTime
expr: sum by (instance, job) ((up == 0) * on (instance) group_right(job) (node_time_seconds - node_boot_time_seconds > 600))
for: 0m
labels:
    severity: critical
annotations:
    summary: Prometheus target missing with warmup time (instance {{ $labels.instance }})
    description: |-
        Allow a job time to start up (10 minutes) before alerting that it's down.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PrometheusTargetMissingWithWarmupTime

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
