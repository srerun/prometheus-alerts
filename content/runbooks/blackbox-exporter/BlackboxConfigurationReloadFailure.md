---
title: BlackboxConfigurationReloadFailure
description: Troubleshooting for alert BlackboxConfigurationReloadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxConfigurationReloadFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Blackbox configuration reload failure

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: BlackboxConfigurationReloadFailure
expr: blackbox_exporter_config_last_reload_successful != 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Blackbox configuration reload failure (instance {{ $labels.instance }})
    description: |-
        Blackbox configuration reload failure
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/BlackboxConfigurationReloadFailure

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
