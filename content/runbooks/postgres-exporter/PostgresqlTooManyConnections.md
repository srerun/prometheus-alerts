---
title: PostgresqlTooManyConnections
description: Troubleshooting for alert PostgresqlTooManyConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlTooManyConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
PostgreSQL instance has too many connections (> 80%).

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlTooManyConnections" %}}

{{% comment %}}

```yaml
alert: PostgresqlTooManyConnections
expr: sum by (instance, job, server) (pg_stat_activity_count) > min by (instance, job, server) (pg_settings_max_connections * 0.8)
for: 2m
labels:
    severity: warning
annotations:
    summary: Postgresql too many connections (instance {{ $labels.instance }})
    description: |-
        PostgreSQL instance has too many connections (> 80%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlTooManyConnections.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
