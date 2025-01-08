---
title: PostgresqlTooManyConnections
description: Troubleshooting for alert PostgresqlTooManyConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlTooManyConnections

PostgreSQL instance has too many connections (> 80%).

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlTooManyConnections" %}}

{{% comment %}}

```yaml
alert: PostgresqlTooManyConnections
expr: sum by (instance, job, server) (pg_stat_activity_count) > min by (instance, job, server) (pg_settings_max_connections * 0.8)
for: 2m
labels:
    severity: warning
annotations:
    summary: Postgresql too many connections (instance {{ $labels.instance }})
    description: |-
        PostgreSQL instance has too many connections (> 80%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlTooManyConnections.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The `PostgresqlTooManyConnections` alert is triggered when the number of active connections to a PostgreSQL instance exceeds 80% of the maximum allowed connections. This can indicate a potential performance issue or resource bottleneck, as excessive connections can lead to increased memory usage, slower query execution, and even crashes.

## Impact

If left unaddressed, this issue can lead to:

* Slow query performance or timeouts
* Increased memory usage, potentially causing OOM errors
* PostgreSQL instance crashes or restarts
* Downtime and loss of application availability

## Diagnosis

To diagnose the cause of the alert, follow these steps:

1. Check the PostgreSQL instance's connection count and max connections using the `pg_stat_activity_count` and `pg_settings_max_connections` metrics.
2. Investigate recent changes to the application or workload that may be contributing to the increased connection count.
3. Review the PostgreSQL instance's configuration and resource utilization (e.g., CPU, memory, disk usage).
4. Check for any slow or long-running queries that may be holding connections open.
5. Verify that the PostgreSQL instance is properly configured for connection pooling and that the connection limit is set appropriately.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and terminate any unnecessary or idle connections using the `pg_terminate_backend` function.
2. Optimize application code to reduce the number of connections required or implement connection pooling.
3. Adjust the PostgreSQL instance's configuration to increase the max connections limit, if necessary.
4. Implement query optimization techniques to reduce the duration of slow queries.
5. Consider scaling the PostgreSQL instance or adding additional resources to handle increased workload demands.

Remember to investigate the root cause of the issue and address it accordingly to prevent future occurrences of this alert.