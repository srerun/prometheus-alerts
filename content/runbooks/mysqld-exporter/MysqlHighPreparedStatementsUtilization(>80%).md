---
title: MysqlHighPreparedStatementsUtilization(>80%)
description: Troubleshooting for alert MysqlHighPreparedStatementsUtilization(>80%)
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlHighPreparedStatementsUtilization(>80%)

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High utilization of prepared statements (>80%) on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlHighPreparedStatementsUtilization(>80%)" %}}

{{% comment %}}

```yaml
alert: MysqlHighPreparedStatementsUtilization(>80%)
expr: max_over_time(mysql_global_status_prepared_stmt_count[1m]) / mysql_global_variables_max_prepared_stmt_count * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: MySQL high prepared statements utilization (> 80%) (instance {{ $labels.instance }})
    description: |-
        High utilization of prepared statements (>80%) on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlHighPreparedStatementsUtilization(>80%).md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
