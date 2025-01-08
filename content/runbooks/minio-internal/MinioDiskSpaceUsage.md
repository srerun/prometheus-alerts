---
title: MinioDiskSpaceUsage
description: Troubleshooting for alert MinioDiskSpaceUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MinioDiskSpaceUsage

Minio available free space is low (< 10%)

<details>
  <summary>Alert Rule</summary>

{{% rule "minio/minio-internal.yml" "MinioDiskSpaceUsage" %}}

{{% comment %}}

```yaml
alert: MinioDiskSpaceUsage
expr: disk_storage_available / disk_storage_total * 100 < 10
for: 0m
labels:
    severity: warning
annotations:
    summary: Minio disk space usage (instance {{ $labels.instance }})
    description: |-
        Minio available free space is low (< 10%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/minio-internal/MinioDiskSpaceUsage.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the MinioDiskSpaceUsage alert:

## Meaning

The MinioDiskSpaceUsage alert is triggered when the available free space on a Minio disk falls below 10%. This alert is critical as it may impact the performance and availability of Minio services.

## Impact

The impact of this alert includes:

* Minio services may become unresponsive or fail to write data
* Data loss or corruption may occur due to insufficient disk space
* Increased latency or errors may be experienced by users and applications relying on Minio services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Minio disk usage metrics to identify the specific disk with low available free space
2. Verify the total disk capacity and available free space using the `disk_storage_total` and `disk_storage_available` metrics
3. Review recent changes to Minio configuration, usage patterns, or data ingestion rates that may have contributed to the low disk space
4. Check the Minio logs for any error messages related to disk space or I/O operations

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately add more disk space to the affected Minio instance to prevent service disruptions
2. Identify and remove any unnecessary data or files to free up disk space
3. Consider implementing data retention policies or data lifecycle management practices to prevent future disk space issues
4. Monitor Minio disk usage metrics closely to ensure the issue does not recur
5. Consider implementing automated disk space monitoring and alerting to detect potential issues before they become critical