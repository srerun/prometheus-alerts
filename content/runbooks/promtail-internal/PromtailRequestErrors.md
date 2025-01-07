---
title: PromtailRequestErrors
description: Troubleshooting for alert PromtailRequestErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PromtailRequestErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}% errors.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PromtailRequestErrors
expr: 100 * sum(rate(promtail_request_duration_seconds_count{status_code=~"5..|failed"}[1m])) by (namespace, job, route, instance) / sum(rate(promtail_request_duration_seconds_count[1m])) by (namespace, job, route, instance) > 10
for: 5m
labels:
    severity: critical
annotations:
    summary: Promtail request errors (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} {{ $labels.route }} is experiencing {{ printf "%.2f" $value }}% errors.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PromtailRequestErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
