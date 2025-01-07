---
title: ThanosRuleQueueIsDroppingAlerts
description: Troubleshooting for alert ThanosRuleQueueIsDroppingAlerts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleQueueIsDroppingAlerts

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Thanos Rule {{$labels.instance}} is failing to queue alerts.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosRuleQueueIsDroppingAlerts
expr: sum by (job, instance) (rate(thanos_alert_queue_alerts_dropped_total{job=~".*thanos-rule.*"}[5m])) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Rule Queue Is Dropping Alerts (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.instance}} is failing to queue alerts.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosRuleQueueIsDroppingAlerts

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
