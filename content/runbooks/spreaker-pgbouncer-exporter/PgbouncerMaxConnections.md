---
title: PgbouncerMaxConnections
description: Troubleshooting for alert PgbouncerMaxConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PgbouncerMaxConnections

The number of PGBouncer client connections has reached max_client_conn.

<details>
  <summary>Alert Rule</summary>

{{% rule "pgbouncer/spreaker-pgbouncer-exporter.yml" "PgbouncerMaxConnections" %}}

{{% comment %}}

```yaml
alert: PgbouncerMaxConnections
expr: increase(pgbouncer_errors_count{errmsg="no more connections allowed (max_client_conn)"}[30s]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: PGBouncer max connections (instance {{ $labels.instance }})
    description: |-
        The number of PGBouncer client connections has reached max_client_conn.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/spreaker-pgbouncer-exporter/PgbouncerMaxConnections.md

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
