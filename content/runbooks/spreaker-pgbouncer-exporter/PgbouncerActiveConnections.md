---
title: PgbouncerActiveConnections
description: Troubleshooting for alert PgbouncerActiveConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PgbouncerActiveConnections

## Meaning
[//]: # "Short paragraph that explains what the alert means"
PGBouncer pools are filling up

<details>
  <summary>Alert Rule</summary>

{{% rule "pgbouncer/spreaker-pgbouncer-exporter.yml" "PgbouncerActiveConnections" %}}

<!-- Rule when generated

```yaml
alert: PgbouncerActiveConnections
expr: pgbouncer_pools_server_active_connections > 200
for: 2m
labels:
    severity: warning
annotations:
    summary: PGBouncer active connections (instance {{ $labels.instance }})
    description: |-
        PGBouncer pools are filling up
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/spreaker-pgbouncer-exporter/PgbouncerActiveConnections.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
