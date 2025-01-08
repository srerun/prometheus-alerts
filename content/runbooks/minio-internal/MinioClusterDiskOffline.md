---
title: MinioClusterDiskOffline
description: Troubleshooting for alert MinioClusterDiskOffline
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MinioClusterDiskOffline

Minio cluster disk is offline

<details>
  <summary>Alert Rule</summary>

{{% rule "minio/minio-internal.yml" "MinioClusterDiskOffline" %}}

{{% comment %}}

```yaml
alert: MinioClusterDiskOffline
expr: minio_cluster_drive_offline_total > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Minio cluster disk offline (instance {{ $labels.instance }})
    description: |-
        Minio cluster disk is offline
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/minio-internal/MinioClusterDiskOffline.md

```

{{% /comment %}}

</details>


Here is a runbook for the MinioClusterDiskOffline alert rule:

## Meaning

The MinioClusterDiskOffline alert is triggered when one or more disks in a Minio cluster are offline. This means that the disk is not available for reading or writing, which can impact the availability and performance of the Minio cluster.

## Impact

The impact of this alert is high, as it can lead to:

* Data loss or corruption
* Inconsistent data across the cluster
* Reduced performance or unavailability of the Minio cluster
* Potential data inconsistency or loss for applications relying on the Minio cluster

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Minio cluster status using the Minio dashboard or CLI tools
2. Identify which disk(s) are offline and investigate the reason for the outage
3. Check the system logs for any errors or warnings related to the disk or Minio cluster
4. Verify that the disk is properly connected and configured
5. Check for any underlying hardware or software issues that may be causing the disk to be offline

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately investigate the cause of the disk offline and take corrective action
2. Attempt to bring the offline disk back online
3. If the disk cannot be brought back online, replace it with a new disk
4. Verify that the Minio cluster is stable and functional after the disk is brought back online or replaced
5. Consider implementing redundancy or failover mechanisms to prevent similar issues in the future

Additional resources:

* Refer to the Minio documentation for troubleshooting and troubleshooting guides
* Consult with the Minio cluster administrators or subject matter experts for further assistance if needed.