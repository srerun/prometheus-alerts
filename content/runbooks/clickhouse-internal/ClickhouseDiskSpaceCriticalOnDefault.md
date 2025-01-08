---
title: ClickhouseDiskSpaceCriticalOnDefault
description: Troubleshooting for alert ClickhouseDiskSpaceCriticalOnDefault
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseDiskSpaceCriticalOnDefault

Disk space on default disk is critically low, below 10%.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseDiskSpaceCriticalOnDefault" %}}

{{% comment %}}

```yaml
alert: ClickhouseDiskSpaceCriticalOnDefault
expr: ClickHouseAsyncMetrics_DiskAvailable_default / (ClickHouseAsyncMetrics_DiskAvailable_default + ClickHouseAsyncMetrics_DiskUsed_default) * 100 < 10
for: 2m
labels:
    severity: critical
annotations:
    summary: ClickHouse Disk Space Critical on Default (instance {{ $labels.instance }})
    description: |-
        Disk space on default disk is critically low, below 10%.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseDiskSpaceCriticalOnDefault.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

This alert is triggered when the available disk space on the default disk of a ClickHouse instance falls below 10%. This is a critical threshold, indicating that the disk is almost full and may cause issues with data ingestion, queries, and overall system performance.

## Impact

If left unaddressed, this issue can lead to:

* Data loss or corruption due to insufficient disk space
* Query timeouts or failures
* Increased error rates
* System instability or crashes
* Potential data inconsistencies or losses

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ClickHouse instance's disk usage and available space using the `ClickHouseAsyncMetrics_DiskAvailable_default` and `ClickHouseAsyncMetrics_DiskUsed_default` metrics.
2. Verify that the default disk is indeed the one running low on space.
3. Investigate recent data ingestion patterns and query loads to identify potential causes for the disk space depletion.
4. Check the ClickHouse instance's configuration and data retention policies to ensure they are set correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately add more disk space to the default disk or consider migrating data to a larger disk.
2. Investigate and address the root cause of the disk space depletion, such as high data ingestion rates or inefficient data compression.
3. Consider implementing data retention policies or purging unnecessary data to free up space.
4. Monitor the disk space usage and adjust the alert threshold or add more disk space as needed to prevent future occurrences.
5. If necessary, consider scaling up the ClickHouse instance or distributing data across multiple instances to alleviate disk space pressure.

Remember to update the ClickHouse instance's configuration and data retention policies to prevent similar issues in the future.