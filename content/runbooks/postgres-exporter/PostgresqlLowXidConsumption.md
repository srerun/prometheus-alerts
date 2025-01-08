---
title: PostgresqlLowXidConsumption
description: Troubleshooting for alert PostgresqlLowXidConsumption
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlLowXidConsumption

Postgresql seems to be consuming transaction IDs very slowly

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlLowXidConsumption" %}}

{{% comment %}}

```yaml
alert: PostgresqlLowXidConsumption
expr: rate(pg_txid_current[1m]) < 5
for: 2m
labels:
    severity: warning
annotations:
    summary: Postgresql low XID consumption (instance {{ $labels.instance }})
    description: |-
        Postgresql seems to be consuming transaction IDs very slowly
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlLowXidConsumption.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `PostgresqlLowXidConsumption`:

## Meaning

The `PostgresqlLowXidConsumption` alert is triggered when the rate of consumption of PostgreSQL transaction IDs (XIDs) is lower than 5 per minute. This alert indicates that the database is not utilizing transactions efficiently, which can lead to issues with transaction ID wraparound.

## Impact

If left unaddressed, low XID consumption can cause the following issues:

* Transaction ID wraparound: If the XID counter reaches its maximum value, it will wrap around to zero, causing data inconsistencies and potential data loss.
* Decreased database performance: Inefficient transaction management can lead to slower query execution times and decreased overall database performance.
* Increased risk of errors: Low XID consumption can lead to errors during transaction processing, resulting in failed transactions and potential data inconsistencies.

## Diagnosis

To diagnose the root cause of low XID consumption, follow these steps:

1. Check the PostgreSQL logs for any errors or warnings related to transaction processing.
2. Verify that the `max_connections` parameter is set correctly and that the number of active connections is within the recommended range.
3. Check the `autovacuum` settings to ensure that they are configured correctly and running regularly.
4. Analyze the query patterns and transaction workload to identify any inefficiencies or bottlenecks.
5. Verify that the PostgreSQL version is up-to-date and that any relevant patches have been applied.

## Mitigation

To mitigate the effects of low XID consumption, follow these steps:

1. Increase the `max_connections` parameter to allow for more concurrent connections.
2. Optimize queries and transactions to reduce the load on the database.
3. Adjust the `autovacuum` settings to run more frequently or with a higher level of aggressiveness.
4. Consider implementing connection pooling or load balancing to distribute the load across multiple database instances.
5. Monitor the XID consumption rate and adjust the alert threshold accordingly to ensure that it is triggered early enough to prevent wraparound.

Additional resources:

* PostgreSQL documentation on transaction IDs and wraparound
* PostgreSQL documentation on autovacuum settings
* PostgreSQL tuning and optimization guides