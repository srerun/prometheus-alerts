---
title: LokiRequestPanic
description: Troubleshooting for alert LokiRequestPanic
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LokiRequestPanic

## Meaning
[//]: # "Short paragraph that explains what the alert means"
The {{ $labels.job }} is experiencing {{ printf "%.2f" $value }}% increase of panics

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: LokiRequestPanic
expr: sum(increase(loki_panic_total[10m])) by (namespace, job) > 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Loki request panic (instance {{ $labels.instance }})
    description: |-
        The {{ $labels.job }} is experiencing {{ printf "%.2f" $value }}% increase of panics
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/LokiRequestPanic

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
