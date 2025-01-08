---
title: MysqlInnodbLogWaits
description: Troubleshooting for alert MysqlInnodbLogWaits
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MysqlInnodbLogWaits

MySQL innodb log writes stalling

<details>
  <summary>Alert Rule</summary>

{{% rule "mysql/mysqld-exporter.yml" "MysqlInnodbLogWaits" %}}

{{% comment %}}

```yaml
alert: MysqlInnodbLogWaits
expr: rate(mysql_global_status_innodb_log_waits[15m]) > 10
for: 0m
labels:
    severity: warning
annotations:
    summary: MySQL InnoDB log waits (instance {{ $labels.instance }})
    description: |-
        MySQL innodb log writes stalling
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/mysqld-exporter/MysqlInnodbLogWaits.md

```

{{% /comment %}}

</details>


Here is a runbook for the MysqlInnodbLogWaits alert:

## Meaning

The MysqlInnodbLogWaits alert is triggered when the rate of InnoDB log waits exceeds 10 per 15 minutes. This indicates that the MySQL database is experiencing issues with writing logs to disk, which can lead to performance problems and even crashes.

## Impact

* Increased latency and slow query performance
* Higher risk of database crashes or corruption
* Potential for data loss or inconsistencies

If left unchecked, this issue can have a significant impact on the availability and reliability of the MySQL database, leading to frustrated users and potentially even financial losses.

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

1. Check the MySQL error log for any indications of disk I/O issues or errors related to log writing.
2. Verify that the disk is not full and that there is sufficient free space for the log files.
3. Check the InnoDB log file size and verify that it is not too large, causing write stalls.
4. Investigate any recent changes to the MySQL configuration or environment that may be contributing to the issue.
5. Run the `SHOW ENGINE INNODB STATUS` command to gather more information about the current InnoDB log wait situation.

## Mitigation

To mitigate this issue, follow these steps:

1. Increase the `innodb_log_file_size` variable to reduce the frequency of log writes.
2. Verify that the disk is optimized for high I/O performance, and consider upgrading to a faster disk if necessary.
3. Implement regular backups and ensure that the backups are completing successfully.
4. Consider increasing the `innodb_log_buffer_size` variable to improve log write performance.
5. If the issue persists, consider consulting with a MySQL expert or seeking additional support to resolve the issue.

Remember to monitor the situation closely and adjust the mitigation strategies as needed to ensure the stability and performance of the MySQL database.