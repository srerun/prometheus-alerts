---
title: PostgresqlLowXidConsumption
description: Troubleshooting for alert PostgresqlLowXidConsumption
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlLowXidConsumption

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Postgresql seems to be consuming transaction IDs very slowly

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PostgresqlLowXidConsumption
expr: rate(pg_txid_current[1m]) < 5
for: 2m
labels:
    severity: warning
annotations:
    summary: Postgresql low XID consumption (instance {{ $labels.instance }})
    description: |-
        Postgresql seems to be consuming transaction IDs very slowly
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/PostgresqlLowXidConsumption

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
