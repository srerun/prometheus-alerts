---
title: PostgresqlDown
description: Troubleshooting for alert PostgresqlDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Postgresql instance is down

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PostgresqlDown
expr: pg_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql down (instance {{ $labels.instance }})
    description: |-
        Postgresql instance is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PostgresqlDown

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
