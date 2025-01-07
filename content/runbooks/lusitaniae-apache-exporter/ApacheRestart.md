---
title: ApacheRestart
description: Troubleshooting for alert ApacheRestart
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ApacheRestart

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Apache has just been restarted.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ApacheRestart
expr: apache_uptime_seconds_total / 60 < 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Apache restart (instance {{ $labels.instance }})
    description: |-
        Apache has just been restarted.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ApacheRestart

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
