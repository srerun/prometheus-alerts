---
title: MysqlDown
description: Troubleshooting for alert MysqlDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlDown

MySQL instance is down on {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlDown" %}}

{{% comment %}}

```yaml
alert: MysqlDown
expr: mysql_up == 0
for: 0m
labels:
    severity: critical
annotations:
    summary: MySQL down (instance {{ $labels.instance }})
    description: |-
        MySQL instance is down on {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlDown.md

```

{{% /comment %}}

</details>


Here is a runbook for the MysqlDown alert:

## Meaning

The MysqlDown alert is triggered when the Prometheus `mysql_up` metric returns a value of 0, indicating that the MySQL instance is not responding or is down. This alert has a severity of critical, indicating that immediate attention is required to prevent data loss or application downtime.

## Impact

The impact of a MySQL instance being down can be severe, including:

* Loss of access to critical data and systems
* Disruption to business operations and revenue
* Inability to perform critical tasks or transactions
* Potential for data corruption or loss

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the MySQL instance logs for errors or issues that may be causing the downstate.
2. Verify that the MySQL instance is not overloaded or experiencing high CPU or memory usage.
3. Check the disk space and verify that it is not full or corrupted.
4. Verify that the network connectivity to the MySQL instance is stable and not experiencing issues.
5. Check the Prometheus and MySQL exporter logs for any errors or issues that may be causing the alert.

## Mitigation

To mitigate the issue, perform the following steps:

1. Restart the MySQL instance to attempt to restore service.
2. Verify that the MySQL instance is properly configured and running with the correct permissions and settings.
3. Check and resolve any underlying issues causing the downstate, such as disk space or network connectivity issues.
4. Monitor the MySQL instance for a period of time to ensure that it remains stable and responsive.
5. If the issue persists, consider escalating to a senior engineer or DBA for further assistance.