---
title: MysqlSlowQueries
description: Troubleshooting for alert MysqlSlowQueries
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlSlowQueries

MySQL server mysql has some new slow query.

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlSlowQueries" %}}

{{% comment %}}

```yaml
alert: MysqlSlowQueries
expr: increase(mysql_global_status_slow_queries[1m]) > 0
for: 2m
labels:
    severity: warning
annotations:
    summary: MySQL slow queries (instance {{ $labels.instance }})
    description: |-
        MySQL server mysql has some new slow query.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlSlowQueries.md

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
