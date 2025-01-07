---
title: MysqlSlaveReplicationLag
description: Troubleshooting for alert MysqlSlaveReplicationLag
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlSlaveReplicationLag

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MySQL replication lag on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: MysqlSlaveReplicationLag
expr: ( (mysql_slave_status_seconds_behind_master - mysql_slave_status_sql_delay) and ON (instance) mysql_slave_status_master_server_id > 0 ) > 30
for: 1m
labels:
    severity: critical
annotations:
    summary: MySQL Slave replication lag (instance {{ $labels.instance }})
    description: |-
        MySQL replication lag on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/MysqlSlaveReplicationLag

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
