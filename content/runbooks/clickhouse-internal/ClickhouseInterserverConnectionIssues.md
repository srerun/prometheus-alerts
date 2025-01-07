---
title: ClickhouseInterserverConnectionIssues
description: Troubleshooting for alert ClickhouseInterserverConnectionIssues
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseInterserverConnectionIssues

## Meaning
[//]: # "Short paragraph that explains what the alert means"
An increase in interserver connections may indicate replication or distributed query handling issues.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: ClickhouseInterserverConnectionIssues
expr: increase(ClickHouseMetrics_InterserverConnection[5m]) > 0
for: 1m
labels:
    severity: warning
annotations:
    summary: ClickHouse Interserver Connection Issues (instance {{ $labels.instance }})
    description: |-
        An increase in interserver connections may indicate replication or distributed query handling issues.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/ClickhouseInterserverConnectionIssues

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
