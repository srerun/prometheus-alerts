---
title: MysqlTooManyConnections(>80%)
description: Troubleshooting for alert MysqlTooManyConnections(>80%)
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlTooManyConnections(>80%)

## Meaning
[//]: # "Short paragraph that explains what the alert means"
More than 80% of MySQL connections are in use on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: MysqlTooManyConnections(>80%)
expr: max_over_time(mysql_global_status_threads_connected[1m]) / mysql_global_variables_max_connections * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: MySQL too many connections (> 80%) (instance {{ $labels.instance }})
    description: |-
        More than 80% of MySQL connections are in use on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/MysqlTooManyConnections(>80%)

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
