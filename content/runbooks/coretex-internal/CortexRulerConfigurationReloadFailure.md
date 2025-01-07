---
title: CortexRulerConfigurationReloadFailure
description: Troubleshooting for alert CortexRulerConfigurationReloadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CortexRulerConfigurationReloadFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Cortex ruler configuration reload failure (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CortexRulerConfigurationReloadFailure
expr: cortex_ruler_config_last_reload_successful != 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Cortex ruler configuration reload failure (instance {{ $labels.instance }})
    description: |-
        Cortex ruler configuration reload failure (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CortexRulerConfigurationReloadFailure

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
