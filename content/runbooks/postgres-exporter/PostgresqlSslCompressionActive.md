---
title: PostgresqlSslCompressionActive
description: Troubleshooting for alert PostgresqlSslCompressionActive
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlSslCompressionActive

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Database connections with SSL compression enabled. This may add significant jitter in replication delay. Replicas should turn off SSL compression via `sslcompression=0` in `recovery.conf`.

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlSslCompressionActive" %}}

<!-- Rule when generated

```yaml
alert: PostgresqlSslCompressionActive
expr: sum(pg_stat_ssl_compression) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql SSL compression active (instance {{ $labels.instance }})
    description: |-
        Database connections with SSL compression enabled. This may add significant jitter in replication delay. Replicas should turn off SSL compression via `sslcompression=0` in `recovery.conf`.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlSslCompressionActive.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
