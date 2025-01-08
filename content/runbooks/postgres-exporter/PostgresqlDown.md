---
title: PostgresqlDown
description: Troubleshooting for alert PostgresqlDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlDown

Postgresql instance is down

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlDown" %}}

{{% comment %}}

```yaml
alert: PostgresqlDown
expr: pg_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql down (instance {{ $labels.instance }})
    description: |-
        Postgresql instance is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlDown.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the PostgresqlDown alert:

## Meaning

The PostgresqlDown alert is triggered when the `pg_up` metric is 0, indicating that the PostgreSQL instance is not responding or unreachable. This alert is critical, as it can lead to data loss, application downtime, and other severe consequences.

## Impact

The impact of this alert can be severe, as it may cause:

* Data loss or corruption
* Application downtime or instability
* Loss of business critical functionality
* Revenue loss or other financial implications

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the PostgreSQL instance status using the `pg_top` or `pg_stat_activity` commands.
2. Verify that the PostgreSQL process is running and listening on the expected port.
3. Check the system logs for any error messages related to PostgreSQL.
4. Validate that the PostgreSQL configuration is correct and up-to-date.
5. Check for any network connectivity issues between the Prometheus server and the PostgreSQL instance.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the PostgreSQL instance or the entire server if necessary.
2. Perform a database consistency check to ensure data integrity.
3. Verify that the PostgreSQL configuration is correct and up-to-date.
4. Implement additional monitoring and logging to detect similar issues in the future.
5. Perform a root cause analysis to identify the underlying cause of the issue and prevent it from happening again.

Remember to follow your organization's specific procedures and guidelines for responding to critical alerts. This runbook is meant to serve as a starting point and should be tailored to your specific environment and needs.