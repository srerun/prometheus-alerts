---
title: SqlServerDown
description: Troubleshooting for alert SqlServerDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SqlServerDown

SQL server instance is down

<details>
  <summary>Alert Rule</summary>

{{% rule "sql-server/ozarklake-mssql-exporter.yml" "SqlServerDown" %}}

{{% comment %}}

```yaml
alert: SqlServerDown
expr: mssql_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: SQL Server down (instance {{ $labels.instance }})
    description: |-
        SQL server instance is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/ozarklake-mssql-exporter/SqlServerDown.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `SqlServerDown`:

## Meaning

The `SqlServerDown` alert is triggered when the `mssql_up` metric has a value of 0, indicating that the SQL Server instance is currently down or unavailable. This alert is critical, as it can impact database operations and potentially lead to data loss or corruption.

## Impact

The impact of this alert is high, as a downed SQL Server instance can:

* Cause database connections to fail, leading to application errors and downtime
* Result in data loss or corruption if the instance is not properly shut down
* Disrupt business operations and revenue if the database is critical to business processes

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the SQL Server instance status using the `mssql_up` metric in Prometheus
2. Verify that the SQL Server instance is not responding to queries or connections
3. Check the SQL Server error logs for any errors or warnings that may indicate the cause of the issue
4. Verify that the network connectivity to the SQL Server instance is stable and functioning correctly

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the SQL Server instance, if possible, to attempt to bring it back online
2. Investigate and resolve any underlying issues that may have caused the instance to go down (e.g. resource constraints, software issues, etc.)
3. Perform a database consistency check and repair any corrupted databases, if necessary
4. Verify that the SQL Server instance is properly configured and running with the correct settings and resources

Note: The runbook URL points to a external resource, which may provide additional information and specific steps for mitigation.