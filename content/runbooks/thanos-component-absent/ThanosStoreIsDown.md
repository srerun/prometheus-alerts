---
title: ThanosStoreIsDown
description: Troubleshooting for alert ThanosStoreIsDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosStoreIsDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
ThanosStore has disappeared. Prometheus target for the component cannot be discovered.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ThanosStoreIsDown
expr: absent(up{job=~".*thanos-store.*"} == 1)
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Store Is Down (instance {{ $labels.instance }})
    description: |-
        ThanosStore has disappeared. Prometheus target for the component cannot be discovered.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ThanosStoreIsDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
