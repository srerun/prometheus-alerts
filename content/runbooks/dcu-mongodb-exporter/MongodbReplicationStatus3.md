---
title: MongodbReplicationStatus3
description: Troubleshooting for alert MongodbReplicationStatus3
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MongodbReplicationStatus3

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MongoDB Replication set member either perform startup self-checks, or transition from completing a rollback or resync

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: MongodbReplicationStatus3
expr: mongodb_replset_member_state == 3
for: 0m
labels:
    severity: critical
annotations:
    summary: MongoDB replication Status 3 (instance {{ $labels.instance }})
    description: |-
        MongoDB Replication set member either perform startup self-checks, or transition from completing a rollback or resync
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/MongodbReplicationStatus3

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
