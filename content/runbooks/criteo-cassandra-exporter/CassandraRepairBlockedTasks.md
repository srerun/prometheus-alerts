---
title: CassandraRepairBlockedTasks
description: Troubleshooting for alert CassandraRepairBlockedTasks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraRepairBlockedTasks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Cassandra repair tasks are blocked

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: CassandraRepairBlockedTasks
expr: cassandra_stats{name="org:apache:cassandra:metrics:threadpools:internal:antientropystage:currentlyblockedtasks:count"} > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra repair blocked tasks (instance {{ $labels.instance }})
    description: |-
        Some Cassandra repair tasks are blocked
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/CassandraRepairBlockedTasks

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
