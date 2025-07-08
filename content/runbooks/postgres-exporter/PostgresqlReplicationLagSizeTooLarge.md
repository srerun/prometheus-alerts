---
title: PostgresqlReplicationLagSizeTooLarge
description: Troubleshooting for alert PostgresqlReplicationLagSizeTooLarge
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlReplicationLagSizeTooLarge

"Replication lag size on server {{$labels.instance}} ({{$labels.application_name}}) is currently {{ $value | humanize1024}}B behind the leader in cluster {{$labels.cluster_name}}"

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlReplicationLagSizeTooLarge" %}}

{{% comment %}}

```yaml
alert: PostgresqlReplicationLagSizeTooLarge
expr: pg_replication_status_lag_size > 1e+09
for: 5m
labels:
    severity: critical
annotations:
    summary: Postgresql is more than 1G behind (instance {{ $labels.instance }})
    description: |-
        "Replication lag size on server {{$labels.instance}} ({{$labels.application_name}}) is currently {{ $value | humanize1024}}B behind the leader in cluster {{$labels.cluster_name}}"
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/postgres-exporter/postgresqlreplicationlagsizetoolarge/

```

{{% /comment %}}

</details>


## Meaning

The PostgresqlReplicationLagSizeTooLarge alert is triggered when the replication lag size of a PostgreSQL instance exceeds 1GB. This means that the standby server is falling behind the primary server in terms of data replication, which can lead to data inconsistencies and potential data loss.

## Impact

The impact of this alert is critical, as it can result in:

* Data inconsistencies between the primary and standby servers
* Potential data loss if the standby server is not brought up to date
* Downtime for applications relying on the standby server
* Increased risk of data corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the PostgreSQL logs for any errors or warnings related to replication
2. Verify the network connectivity between the primary and standby servers
3. Check the disk space and I/O performance of the standby server
4. Review the PostgreSQL configuration files to ensure that the replication settings are correct
5. Use the `pg_replication_status_lag_size` metric to monitor the replication lag size and identify the root cause of the issue

## Mitigation

To mitigate the issue, follow these steps:

1. Check the PostgreSQL replication settings and adjust them as necessary
2. Verify that the standby server has sufficient disk space and I/O performance
3. Restart the PostgreSQL service on the standby server to re-establish replication
4. Consider increasing the replication timeout or adjusting the replication strategy to reduce the lag size
5. Implement a backup and restore process to ensure data consistency in case of a failover
6. Follow the runbook provided in the annotations for detailed instructions on resolving the issue: <https://srerun.github.io/prometheus-alerts/runbooks/postgres-exporter/postgresqlreplicationlagsizetoolarge/>