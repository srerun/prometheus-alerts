---
title: PostgresqlConfigurationChanged
description: Troubleshooting for alert PostgresqlConfigurationChanged
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlConfigurationChanged

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Postgres Database configuration change has occurred

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlConfigurationChanged" %}}

<!-- Rule when generated

```yaml
alert: PostgresqlConfigurationChanged
expr: '{__name__=~"pg_settings_.*"} != ON(__name__, instance) {__name__=~"pg_settings_([^t]|t[^r]|tr[^a]|tra[^n]|tran[^s]|trans[^a]|transa[^c]|transac[^t]|transact[^i]|transacti[^o]|transactio[^n]|transaction[^_]|transaction_[^r]|transaction_r[^e]|transaction_re[^a]|transaction_rea[^d]|transaction_read[^_]|transaction_read_[^o]|transaction_read_o[^n]|transaction_read_on[^l]|transaction_read_onl[^y]).*"} OFFSET 5m'
for: 0m
labels:
    severity: info
annotations:
    summary: Postgresql configuration changed (instance {{ $labels.instance }})
    description: |-
        Postgres Database configuration change has occurred
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlConfigurationChanged.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
