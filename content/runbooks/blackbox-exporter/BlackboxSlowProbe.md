---
title: BlackboxSlowProbe
description: Troubleshooting for alert BlackboxSlowProbe
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxSlowProbe

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Blackbox probe took more than 1s to complete

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: BlackboxSlowProbe
expr: avg_over_time(probe_duration_seconds[1m]) > 1
for: 1m
labels:
    severity: warning
annotations:
    summary: Blackbox slow probe (instance {{ $labels.instance }})
    description: |-
        Blackbox probe took more than 1s to complete
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/BlackboxSlowProbe

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
