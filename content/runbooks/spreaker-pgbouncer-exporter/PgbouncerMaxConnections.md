---
title: PgbouncerMaxConnections
description: Troubleshooting for alert PgbouncerMaxConnections
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PgbouncerMaxConnections

The number of PGBouncer client connections has reached max_client_conn.

<details>
  <summary>Alert Rule</summary>

{{% rule "pgbouncer/spreaker-pgbouncer-exporter.yml" "PgbouncerMaxConnections" %}}

{{% comment %}}

```yaml
alert: PgbouncerMaxConnections
expr: increase(pgbouncer_errors_count{errmsg="no more connections allowed (max_client_conn)"}[30s]) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: PGBouncer max connections (instance {{ $labels.instance }})
    description: |-
        The number of PGBouncer client connections has reached max_client_conn.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/spreaker-pgbouncer-exporter/PgbouncerMaxConnections.md

```

{{% /comment %}}

</details>


Here is the runbook for the PgbouncerMaxConnections alert:

## Meaning
The PgbouncerMaxConnections alert is triggered when the number of PGBouncer client connections reaches the maximum allowed limit (max_client_conn). This means that PGBouncer is refusing new connections, which can impact the availability and performance of dependent services.

## Impact
The impact of this alert is critical, as it can cause:

* Service unavailability: New connections to the database may be refused, leading to errors and downtime for dependent services.
* Performance degradation: Existing connections may be affected, leading to slow performance and latency issues.
* Data loss: In extreme cases, data may be lost or corrupted due to the inability to establish new connections.

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the PGBouncer dashboard for the instance in question to confirm the max connections threshold has been reached.
2. Review the PGBouncer logs to identify the reason for the connection limit being reached (e.g., high traffic, resource constraints, etc.).
3. Verify that the max_client_conn setting is set correctly and adjust it if necessary.
4. Check for any other alerts or issues related to PGBouncer or dependent services.

## Mitigation
To mitigate the issue, follow these steps:

1. Increase the max_client_conn setting in PGBouncer configuration to allow more connections.
2. Identify and address the root cause of the high connection demand (e.g., optimizing queries, reducing traffic, etc.).
3. Implement connection pooling or other optimization techniques to reduce the load on PGBouncer.
4. Consider scaling up or load-balancing PGBouncer instances to distribute the connection load.
5. Monitor the situation closely and be prepared to roll back changes if they do not resolve the issue.