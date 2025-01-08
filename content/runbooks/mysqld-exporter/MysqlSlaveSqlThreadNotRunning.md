---
title: MysqlSlaveSqlThreadNotRunning
description: Troubleshooting for alert MysqlSlaveSqlThreadNotRunning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlSlaveSqlThreadNotRunning

MySQL Slave SQL thread not running on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlSlaveSqlThreadNotRunning" %}}

{{% comment %}}

```yaml
alert: MysqlSlaveSqlThreadNotRunning
expr: ( mysql_slave_status_slave_sql_running and ON (instance) mysql_slave_status_master_server_id > 0) == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: MySQL Slave SQL thread not running (instance {{ $labels.instance }})
    description: |-
        MySQL Slave SQL thread not running on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlSlaveSqlThreadNotRunning.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
