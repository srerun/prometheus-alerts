---
title: PostgresqlReplicationBehind
description: Troubleshooting for alert PostgresqlReplicationBehind
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlReplicationBehind

"Replication lag is greater than 60 seconds on server {{$labels.instance}}.  Currently {{ $value }} seconds behind"

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlReplicationBehind" %}}

{{% comment %}}

```yaml
alert: PostgresqlReplicationBehind
expr: pg_replication_lag_seconds > 60
for: 5m
labels:
    severity: warning
annotations:
    summary: Postgresql replication is more than 60s behind
    description: |-
        "Replication lag is greater than 60 seconds on server {{$labels.instance}}.  Currently {{ $value }} seconds behind"
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/postgres-exporter/postgresqlreplicationbehind/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The `PostgresqlReplicationBehind` alert is triggered when the replication lag of a PostgreSQL server exceeds 60 seconds. This means that the standby server is not keeping up with the primary server, and data may not be properly replicated.

## Impact

If left unresolved, this issue can lead to:

* Data inconsistencies between the primary and standby servers
* Increased risk of data loss in the event of a failover
* Performance degradation due to increased load on the primary server
* Potential for prolonged downtime during failover or maintenance

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the PostgreSQL server logs for any errors or warnings related to replication
2. Verify that the standby server is properly configured and running
3. Check the network connection between the primary and standby servers for any issues
4. Verify that the replication lag is not caused by a slow network connection
5. Check the PostgreSQL server metrics (e.g. `pg_stat_replication`) to determine the cause of the replication lag

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve any underlying issues causing the replication lag
2. Adjust the PostgreSQL configuration to optimize replication performance (e.g. increasing the `wal_sender_timeout` parameter)
3. Consider increasing the resources (e.g. CPU, memory) of the standby server to improve replication performance
4. Implement additional monitoring and alerting to detect replication issues earlier
5. Perform a failover to the standby server to ensure data consistency and availability

For more detailed instructions and troubleshooting steps, please refer to the [PostgresqlReplicationBehind runbook](https://srerun.github.io/prometheus-alerts/runbooks/postgres-exporter/postgresqlreplicationbehind/).