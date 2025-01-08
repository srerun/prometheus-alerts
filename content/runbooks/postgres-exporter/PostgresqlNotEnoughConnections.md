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


Here is the runbook for the Prometheus alert rule `PostgresqlNotEnoughConnections`:

## Meaning

The `PostgresqlNotEnoughConnections` alert is triggered when the total number of connections to a PostgreSQL instance falls below 5. This alert is raised to prevent potential issues with the database performance and availability.

## Impact

If this alert is not addressed, it may lead to:

* Degraded database performance
* Increased latency
* Errors or failures in applications relying on the database
* Potential data loss or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the PostgreSQL instance's connection count using the `pg_stat_activity` metric.
2. Verify that the instance is not experiencing any other issues, such as high CPU usage or disk space issues.
3. Review the database configuration and connection settings to ensure they are adequate for the workload.
4. Investigate recent changes to the database or application configuration that may have caused the connection count to drop.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the maximum number of connections allowed in the PostgreSQL configuration file (e.g., `postgresql.conf`).
2. Review and optimize database queries to reduce the load on the instance.
3. Consider load balancing or horizontal scaling to distribute the workload across multiple instances.
4. Implement connection pooling or other optimization techniques to improve database performance.
5. Monitor the instance's connection count and adjust as necessary to prevent future occurrences of this alert.