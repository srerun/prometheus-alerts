---
title: HadoopHdfsDiskSpaceLow
description: Troubleshooting for alert HadoopHdfsDiskSpaceLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopHdfsDiskSpaceLow

Available HDFS disk space is running low.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopHdfsDiskSpaceLow" %}}

{{% comment %}}

```yaml
alert: HadoopHdfsDiskSpaceLow
expr: (hadoop_hdfs_bytes_total - hadoop_hdfs_bytes_used) / hadoop_hdfs_bytes_total < 0.1
for: 15m
labels:
    severity: warning
annotations:
    summary: Hadoop HDFS Disk Space Low (instance {{ $labels.instance }})
    description: |-
        Available HDFS disk space is running low.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopHdfsDiskSpaceLow.md

```

{{% /comment %}}

</details>


Here is a runbook for the HadoopHdfsDiskSpaceLow alert:

## Meaning

The HadoopHdfsDiskSpaceLow alert is triggered when the available disk space on a Hadoop HDFS (Hadoop Distributed File System) instance falls below 10% of the total disk space. This alert is critical because HDFS is a fundamental component of the Hadoop ecosystem, and running out of disk space can cause data loss, corruption, or even bring down the entire cluster.

## Impact

If not addressed promptly, low disk space on HDFS can lead to:

* Data loss or corruption
* Inability to write new data to HDFS
* Cluster instability or crashes
* Delays or failures in data processing pipelines
* Impact on dependent applications and services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HDFS disk usage metrics in Prometheus to identify the affected instance(s) and the current usage levels.
2. Verify that the HDFS disk usage is not fluctuating due to temporary conditions, such as large file deletions or network issues.
3. Investigate recent changes to the HDFS configuration, such as changes to block sizes, replication factors, or data compression.
4. Review the HDFS usage patterns to identify any unusual activity or trends.
5. Check for any errors or warnings in the HDFS logs related to disk space issues.

## Mitigation

To mitigate the issue, follow these steps:

1. **Free up disk space**: Remove unnecessary files, delete unnecessary data, or archive cold data to free up disk space.
2. **Add more storage**: Add more disk storage to the HDFS instance to increase capacity.
3. **Optimize HDFS configuration**: Optimize HDFS configuration settings, such as block sizes, replication factors, or data compression, to reduce disk usage.
4. **Monitor HDFS usage**: Closely monitor HDFS disk usage to prevent future low disk space issues.
5. **Implement disk space alerts**: Set up alerts for low disk space thresholds to ensure prompt notification of potential issues.

Remember to also investigate the root cause of the low disk space issue to prevent it from occurring again in the future.