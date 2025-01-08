---
title: ClickhouseDiskSpaceLowOnDefault
description: Troubleshooting for alert ClickhouseDiskSpaceLowOnDefault
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseDiskSpaceLowOnDefault

Disk space on default is below 20%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseDiskSpaceLowOnDefault" %}}

{{% comment %}}

```yaml
alert: ClickhouseDiskSpaceLowOnDefault
expr: ClickHouseAsyncMetrics_DiskAvailable_default / (ClickHouseAsyncMetrics_DiskAvailable_default + ClickHouseAsyncMetrics_DiskUsed_default) * 100 < 20
for: 2m
labels:
    severity: warning
annotations:
    summary: ClickHouse Disk Space Low on Default (instance {{ $labels.instance }})
    description: |-
        Disk space on default is below 20%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseDiskSpaceLowOnDefault.md

```

{{% /comment %}}

</details>


## Meaning

The ClickhouseDiskSpaceLowOnDefault alert is triggered when the available disk space on the default storage of a ClickHouse node falls below 20%. This indicates that the node is running low on disk space, which can lead to performance issues, data loss, or even database crashes.

## Impact

The impact of this alert is moderate to high, as it can:

* Cause slow query performance or timeouts due to low disk space
* Lead to data loss or corruption if the disk becomes full
* Prevent ClickHouse from writing new data, leading to gaps in data collection
* Potentially cause the ClickHouse node to crash or become unavailable

## Diagnosis

To diagnose this issue, follow these steps:

1. Check the ClickHouse node's disk usage and available space using the `df -h` command or a equivalent tool.
2. Verify that the default storage is indeed running low on disk space.
3. Check the ClickHouse logs for any errors or warnings related to disk space issues.
4. Identify any potential causes for the low disk space, such as:
	* High data ingestion rates
	* Inefficient data compression or storage
	* Lack of disk space monitoring or maintenance

## Mitigation

To mitigate this issue, follow these steps:

1. Immediately free up disk space by:
	* Deleting unnecessary data or files
	* Compressing or optimizing data storage
	* Moving data to a different storage location
2. Implement disk space monitoring and alerting to catch low disk space issues before they become critical.
3. Increase the disk space available to the ClickHouse node by:
	* Adding more disk space to the node
	* Migrating to a larger storage instance
	* Implementing a more efficient data storage strategy
4. Review and adjust data ingestion rates and compression settings to prevent future disk space issues.
5. Consider implementing automated disk space maintenance tasks, such as regular cleanups or data pruning.