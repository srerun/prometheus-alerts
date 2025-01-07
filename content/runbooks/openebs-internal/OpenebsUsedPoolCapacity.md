---
title: OpenebsUsedPoolCapacity
description: Troubleshooting for alert OpenebsUsedPoolCapacity
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OpenebsUsedPoolCapacity

## Meaning
[//]: # "Short paragraph that explains what the alert means"
OpenEBS Pool use more than 80% of his capacity

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: OpenebsUsedPoolCapacity
expr: openebs_used_pool_capacity_percent > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: OpenEBS used pool capacity (instance {{ $labels.instance }})
    description: |-
        OpenEBS Pool use more than 80% of his capacity
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/OpenebsUsedPoolCapacity

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
