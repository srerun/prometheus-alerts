---
title: PostgresqlRestarted
description: Troubleshooting for alert PostgresqlRestarted
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlRestarted

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Postgresql restarted

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PostgresqlRestarted
expr: time() - pg_postmaster_start_time_seconds < 60
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql restarted (instance {{ $labels.instance }})
    description: |-
        Postgresql restarted
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PostgresqlRestarted

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
