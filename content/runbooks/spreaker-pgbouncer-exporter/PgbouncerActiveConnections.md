---
title: PgbouncerActiveConnections
description: Troubleshooting for alert PgbouncerActiveConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PgbouncerActiveConnections

PGBouncer pools are filling up

<details>
  <summary>Alert Rule</summary>

{{% rule "pgbouncer/spreaker-pgbouncer-exporter.yml" "PgbouncerActiveConnections" %}}

{{% comment %}}

```yaml
alert: PgbouncerActiveConnections
expr: pgbouncer_pools_server_active_connections > 200
for: 2m
labels:
    severity: warning
annotations:
    summary: PGBouncer active connections (instance {{ $labels.instance }})
    description: |-
        PGBouncer pools are filling up
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/spreaker-pgbouncer-exporter/PgbouncerActiveConnections.md

```

{{% /comment %}}

</details>


Here is a runbook for the PgbouncerActiveConnections alert rule:

## Meaning

The PgbouncerActiveConnections alert is triggered when the number of active connections to a Pgbouncer pool exceeds 200 for more than 2 minutes. This indicates that the Pgbouncer instance is experiencing a high load, which may impact the performance and availability of the PostgreSQL database.

## Impact

* High latency or timeouts for database queries
* Increased load on the Pgbouncer instance, potentially leading to crashes or errors
* Impacted database performance, leading to cascading failures in dependent services
* Potential data loss or inconsistencies due to failed or delayed transactions

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Pgbouncer metrics to identify the specific pool experiencing high active connections.
2. Verify that the PostgreSQL database is not experiencing any issues, such as high CPU usage or disk I/O bottlenecks.
3. Investigate recent changes to the application or service using the Pgbouncer instance, such as increased traffic or new feature deployments.
4. Review the Pgbouncer configuration to ensure it is properly tuned for the current workload.

## Mitigation

To mitigate the issue, follow these steps:

1. **Temporary solution**: Reduce the load on the Pgbouncer instance by shedding connections or implementing rate limiting on the application or service.
2. **Scaling**: Provision additional Pgbouncer instances or increase the resources (e.g., CPU, memory) of the existing instance to handle the increased load.
3. **Configuration adjustments**: Adjust the Pgbouncer configuration to optimize performance, such as increasing the maximum number of connections or adjusting the connection pooling settings.
4. **Root cause analysis**: Identify and address the underlying cause of the increased load, such as optimizing database queries or improving application performance.

Remember to revert any temporary solutions once the root cause has been addressed and the issue is resolved.