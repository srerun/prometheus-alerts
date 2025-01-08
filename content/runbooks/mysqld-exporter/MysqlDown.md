---
title: MysqlDown
description: Troubleshooting for alert MysqlDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
MySQL instance is down on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlDown" %}}

<!-- Rule when generated

```yaml
alert: MysqlDown
expr: mysql_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: MySQL down (instance {{ $labels.instance }})
    description: |-
        MySQL instance is down on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlDown.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
