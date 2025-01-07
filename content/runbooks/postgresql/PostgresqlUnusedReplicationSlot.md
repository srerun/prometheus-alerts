---
title: PostgresqlUnusedReplicationSlot
description: Troubleshooting for alert PostgresqlUnusedReplicationSlot
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlUnusedReplicationSlot

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Unused Replication Slots

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: PostgresqlUnusedReplicationSlot
expr: pg_replication_slots_active == 0
for: 1m
labels:
    severity: warning
annotations:
    summary: Postgresql unused replication slot (instance {{ $labels.instance }})
    description: |-
        Unused Replication Slots
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/PostgresqlUnusedReplicationSlot

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
