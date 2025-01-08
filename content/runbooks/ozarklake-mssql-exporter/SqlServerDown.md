---
title: SqlServerDown
description: Troubleshooting for alert SqlServerDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SqlServerDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
SQL server instance is down

<details>
  <summary>Alert Rule</summary>

{{% rule "sql-server/ozarklake-mssql-exporter.yml" "SqlServerDown" %}}

{{% comment %}}

```yaml
alert: SqlServerDown
expr: mssql_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: SQL Server down (instance {{ $labels.instance }})
    description: |-
        SQL server instance is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ozarklake-mssql-exporter/SqlServerDown.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
