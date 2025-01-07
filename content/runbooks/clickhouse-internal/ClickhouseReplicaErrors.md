---
title: ClickhouseReplicaErrors
description: Troubleshooting for alert ClickhouseReplicaErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseReplicaErrors

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Critical replica errors detected, either all replicas are stale or lost.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ClickhouseReplicaErrors
expr: ClickHouseErrorMetric_ALL_REPLICAS_ARE_STALE == 1 or ClickHouseErrorMetric_ALL_REPLICAS_LOST == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: ClickHouse Replica Errors (instance {{ $labels.instance }})
    description: |-
        Critical replica errors detected, either all replicas are stale or lost.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ClickhouseReplicaErrors

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
