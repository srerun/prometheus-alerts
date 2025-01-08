---
title: MysqlRestarted
description: Troubleshooting for alert MysqlRestarted
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlRestarted

MySQL has just been restarted, less than one minute ago on {{ $labels.instance }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlRestarted" %}}

{{% comment %}}

```yaml
alert: MysqlRestarted
expr: mysql_global_status_uptime < 60
for: 0m
labels:
    severity: info
annotations:
    summary: MySQL restarted (instance {{ $labels.instance }})
    description: |-
        MySQL has just been restarted, less than one minute ago on {{ $labels.instance }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlRestarted.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the `MysqlRestarted` alert:

## Meaning

The `MysqlRestarted` alert is triggered when the MySQL uptime is less than 60 seconds, indicating that the MySQL instance has been restarted recently. This alert is informational in nature, and its purpose is to notify the operation team that a MySQL restart has occurred.

## Impact

The impact of this alert is moderate, as a MySQL restart can cause temporary disruptions to dependent services and applications. Additionally, a restart may indicate underlying issues with the MySQL instance, such as hardware or software failures, that need to be investigated and addressed to prevent future restarts.

## Diagnosis

To diagnose the cause of the MySQL restart, follow these steps:

1. Check the MySQL error logs for any errors or warnings that may have led to the restart.
2. Verify that the MySQL instance is running and accepting connections.
3. Check the system logs for any signs of hardware or software failures that may have caused the restart.
4. Review the MySQL configuration files to ensure that they are correct and up-to-date.
5. Check the MySQL metrics, such as CPU usage, memory usage, and disk usage, to identify any resource bottlenecks that may have contributed to the restart.

## Mitigation

To mitigate the impact of the MySQL restart, follow these steps:

1. Verify that dependent services and applications are functioning correctly and restart them if necessary.
2. Investigate and address the underlying cause of the restart to prevent future occurrences.
3. Perform a thorough review of the MySQL instance to ensure that it is configured correctly and optimized for performance.
4. Consider implementing monitoring and alerting for MySQL metrics to detect potential issues before they cause a restart.
5. Develop a plan for regularly scheduled maintenance and upgrades to minimize the likelihood of unplanned restarts.