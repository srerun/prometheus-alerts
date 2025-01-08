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

Postgres Database configuration change has occurred

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlConfigurationChanged" %}}

{{% comment %}}

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

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `PostgresqlConfigurationChanged`:

## Meaning

This alert is triggered when a change is detected in the PostgreSQL configuration. The change is detected by monitoring the `pg_settings` metrics, which are exposed by the PostgreSQL exporter. The alert rule uses a complex expression to filter out certain settings that are expected to change, such as transaction settings.

## Impact

A change in the PostgreSQL configuration can have significant implications on the performance, security, and availability of the database. Unintended changes can lead to data corruption, performance degradation, or even data loss. This alert is critical in detecting and responding to unauthorized or unintentional changes to the database configuration.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `pg_settings` metrics to identify which specific setting has changed.
2. Verify that the change was intended and authorized.
3. Review the PostgreSQL logs to identify the source of the change (e.g., which user or process made the change).
4. Check for any related errors or issues in the database or application logs.

## Mitigation

To mitigate the issue, follow these steps:

1. Revert the change to the previous known good state, if possible.
2. Investigate and address the root cause of the unintended change (e.g., unauthorized access, incorrect configuration management).
3. Implement additional monitoring and auditing measures to detect and prevent future changes.
4. Review and update the PostgreSQL configuration management process to prevent similar issues in the future.

Note: The mitigation steps may vary depending on the specific change and impact. This runbook provides a general guideline, and the actual mitigation steps may require customization based on the specific situation.