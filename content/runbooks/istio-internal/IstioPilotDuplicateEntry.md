---
title: IstioPilotDuplicateEntry
description: Troubleshooting for alert IstioPilotDuplicateEntry
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioPilotDuplicateEntry

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Istio pilot duplicate entry error.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: IstioPilotDuplicateEntry
expr: sum(rate(pilot_duplicate_envoy_clusters{}[5m])) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Istio Pilot Duplicate Entry (instance {{ $labels.instance }})
    description: |-
        Istio pilot duplicate entry error.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/IstioPilotDuplicateEntry

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
