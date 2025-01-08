---
title: PostgresqlRestarted
description: Troubleshooting for alert PostgresqlRestarted
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlRestarted

Postgresql restarted

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlRestarted" %}}

{{% comment %}}

```yaml
alert: PostgresqlRestarted
expr: time() - pg_postmaster_start_time_seconds < 60
for: 0m
labels:
    severity: critical
annotations:
    summary: Postgresql restarted (instance {{ $labels.instance }})
    description: |-
        Postgresql restarted
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlRestarted.md

```

{{% /comment %}}

</details>


## Meaning

The `PostgresqlRestarted` alert is triggered when the Postgresql instance restarts. This alert indicates that the database server has been restarted, which may cause temporary unavailability of the database and potential data loss.

## Impact

The impact of this alert can be significant, as it may:

* Cause temporary unavailability of the database, leading to errors or timeouts for applications that rely on it.
* Result in potential data loss or corruption if the database was not properly shut down.
* Require manual intervention to assess and address the cause of the restart.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Postgresql logs to determine the reason for the restart.
2. Verify that the Postgresql instance is running and available.
3. Check for any error messages or alerts related to the restart.
4. Review system logs for any signs of system crashes or other issues that may have contributed to the restart.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify that the Postgresql instance is running and available.
2. Run a database consistency check to ensure that the database is intact.
3. Review system logs and Postgresql logs to determine the root cause of the restart.
4. Take corrective action to prevent future restarts, such as addressing system crashes or configuration issues.
5. Consider implementing automatic failover or high availability solutions to minimize the impact of future restarts.
6. Notify stakeholders and application owners of the outage and provide an estimated time to recovery.

Note: For detailed instructions and troubleshooting steps, refer to the [PostgresqlRestarted runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlRestarted.md).