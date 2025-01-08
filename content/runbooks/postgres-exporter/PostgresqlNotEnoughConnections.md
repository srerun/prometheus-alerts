---
title: PostgresqlNotEnoughConnections
description: Troubleshooting for alert PostgresqlNotEnoughConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlNotEnoughConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
PostgreSQL instance should have more connections (> 5)

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlNotEnoughConnections" %}}

{{% comment %}}

```yaml
alert: PostgresqlNotEnoughConnections
expr: sum by (datname) (pg_stat_activity_count{datname!~"template.*|postgres"}) < 5
for: 2m
labels:
    severity: warning
annotations:
    summary: Postgresql not enough connections (instance {{ $labels.instance }})
    description: |-
        PostgreSQL instance should have more connections (> 5)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlNotEnoughConnections.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
