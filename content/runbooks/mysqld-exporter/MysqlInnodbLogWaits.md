---
title: MysqlInnodbLogWaits
description: Troubleshooting for alert MysqlInnodbLogWaits
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlInnodbLogWaits

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MySQL innodb log writes stalling

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlInnodbLogWaits" %}}

<!-- Rule when generated

```yaml
alert: MysqlInnodbLogWaits
expr: rate(mysql_global_status_innodb_log_waits[15m]) > 10
for: 0m
labels:
    severity: warning
annotations:
    summary: MySQL InnoDB log waits (instance {{ $labels.instance }})
    description: |-
        MySQL innodb log writes stalling
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlInnodbLogWaits.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
