---
title: MysqlHighThreadsRunning
description: Troubleshooting for alert MysqlHighThreadsRunning
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlHighThreadsRunning

## Meaning
[//]: # "Short paragraph that explains what the alert means"
More than 60% of MySQL connections are in running state on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlHighThreadsRunning" %}}

<!-- Rule when generated

```yaml
alert: MysqlHighThreadsRunning
expr: max_over_time(mysql_global_status_threads_running[1m]) / mysql_global_variables_max_connections * 100 > 60
for: 2m
labels:
    severity: warning
annotations:
    summary: MySQL high threads running (instance {{ $labels.instance }})
    description: |-
        More than 60% of MySQL connections are in running state on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlHighThreadsRunning.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
