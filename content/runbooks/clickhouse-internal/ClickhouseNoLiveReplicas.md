---
title: ClickhouseNoLiveReplicas
description: Troubleshooting for alert ClickhouseNoLiveReplicas
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseNoLiveReplicas

## Meaning
[//]: # "Short paragraph that explains what the alert means"
There are too few live replicas available, risking data loss and service disruption.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ClickhouseNoLiveReplicas
expr: ClickHouseErrorMetric_TOO_FEW_LIVE_REPLICAS == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: ClickHouse No Live Replicas (instance {{ $labels.instance }})
    description: |-
        There are too few live replicas available, risking data loss and service disruption.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ClickhouseNoLiveReplicas

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
