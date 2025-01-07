---
title: LokiProcessTooManyRestarts
description: Troubleshooting for alert LokiProcessTooManyRestarts
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiProcessTooManyRestarts

## Meaning
[//]: # "Short paragraph that explains what the alert means"
A loki process had too many restarts (target {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: LokiProcessTooManyRestarts
expr: changes(process_start_time_seconds{job=~".*loki.*"}[15m]) > 2
for: 0m
labels:
    severity: warning
annotations:
    summary: Loki process too many restarts (instance {{ $labels.instance }})
    description: |-
        A loki process had too many restarts (target {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/LokiProcessTooManyRestarts

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
