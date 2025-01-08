---
title: PgbouncerErrors
description: Troubleshooting for alert PgbouncerErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PgbouncerErrors

PGBouncer is logging errors. This may be due to a a server restart or an admin typing commands at the pgbouncer console.

<details>
  <summary>Alert Rule</summary>

{{% rule "pgbouncer/spreaker-pgbouncer-exporter.yml" "PgbouncerErrors" %}}

{{% comment %}}

```yaml
alert: PgbouncerErrors
expr: increase(pgbouncer_errors_count{errmsg!="server conn crashed?"}[1m]) > 10
for: 0m
labels:
    severity: warning
annotations:
    summary: PGBouncer errors (instance {{ $labels.instance }})
    description: |-
        PGBouncer is logging errors. This may be due to a a server restart or an admin typing commands at the pgbouncer console.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/spreaker-pgbouncer-exporter/PgbouncerErrors.md

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
