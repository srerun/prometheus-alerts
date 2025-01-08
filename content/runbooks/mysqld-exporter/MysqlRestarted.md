---
title: MysqlRestarted
description: Troubleshooting for alert MysqlRestarted
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlRestarted

MySQL has just been restarted, less than one minute ago on {{ $labels.instance }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlRestarted" %}}

{{% comment %}}

```yaml
alert: MysqlRestarted
expr: mysql_global_status_uptime < 60
for: 0m
labels:
    severity: info
annotations:
    summary: MySQL restarted (instance {{ $labels.instance }})
    description: |-
        MySQL has just been restarted, less than one minute ago on {{ $labels.instance }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlRestarted.md

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
