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

Database connections with SSL compression enabled. This may add significant jitter in replication delay. Replicas should turn off SSL compression via `sslcompression=0` in `recovery.conf`.

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlSslCompressionActive" %}}

{{% comment %}}

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

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `PostgresqlSslCompressionActive`:

## Meaning

This alert is triggered when SSL compression is active on a PostgreSQL instance. SSL compression can add significant jitter to replication delay, which can impact the performance and reliability of the database.

## Impact

The impact of this alert is high, as it can lead to:

* Increased latency and variability in database replication
* Decreased performance and throughput of the database
* Potential data inconsistencies and errors due to delayed replication

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the PostgreSQL instance configuration to confirm that SSL compression is enabled.
2. Verify that the `sslcompression` parameter is set to a non-zero value in the `recovery.conf` file.
3. Review database performance metrics to identify any signs of replication delay or inconsistency.
4. Check the PostgreSQL error logs for any related errors or warnings.

## Mitigation

To mitigate the issue, follow these steps:

1. Update the `recovery.conf` file to set `sslcompression=0` to disable SSL compression.
2. Restart the PostgreSQL instance to apply the configuration change.
3. Monitor database performance metrics to ensure that replication delay and inconsistency have been resolved.
4. Verify that the alert has been resolved and the database is operating normally.

Note: It is recommended to test the changes in a non-production environment before applying them to production systems.