---
title: ClickhouseDiskSpaceLowOnBackups
description: Troubleshooting for alert ClickhouseDiskSpaceLowOnBackups
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseDiskSpaceLowOnBackups

Disk space on backups is below 20%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseDiskSpaceLowOnBackups" %}}

{{% comment %}}

```yaml
alert: ClickhouseDiskSpaceLowOnBackups
expr: ClickHouseAsyncMetrics_DiskAvailable_backups / (ClickHouseAsyncMetrics_DiskAvailable_backups + ClickHouseAsyncMetrics_DiskUsed_backups) * 100 < 20
for: 2m
labels:
    severity: warning
annotations:
    summary: ClickHouse Disk Space Low on Backups (instance {{ $labels.instance }})
    description: |-
        Disk space on backups is below 20%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseDiskSpaceLowOnBackups.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ClickhouseDiskSpaceLowOnBackups`:

## Meaning

The `ClickhouseDiskSpaceLowOnBackups` alert is triggered when the available disk space on the ClickHouse backups falls below 20%. This alert indicates that the disk space on the backups is running low, which can lead to issues with ClickHouse's ability to store new data and perform writes.

## Impact

If this alert is not addressed, ClickHouse may become unable to write new data, leading to data loss and potential system downtime. Additionally, low disk space can cause ClickHouse to slow down or become unresponsive, impacting the performance of dependent services.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ClickHouse backups disk usage to identify the root cause of the low disk space.
2. Verify that the backups are not failing or stuck, causing the disk space to fill up rapidly.
3. Check the ClickHouse logs for any errors or warnings related to disk space or writes.
4. Review the disk usage trends to determine if this is a sudden issue or a gradual problem.

## Mitigation

To mitigate the issue, follow these steps:

1. **Free up disk space**: Immediately free up disk space on the backups by deleting unnecessary files or truncating log files.
2. **Increase disk space**: Increase the disk space allocated to the backups to prevent future instances of low disk space.
3. **Optimize backups**: Optimize the backup process to reduce the disk space usage, such as by compressing backups or using more efficient storage formats.
4. **Monitor disk space**: Implement proactive monitoring of disk space on the backups to detect low disk space issues before they become critical.

Note: This runbook is a general guideline and may need to be tailored to your specific environment and requirements.