---
title: MinioNodeDiskOffline
description: Troubleshooting for alert MinioNodeDiskOffline
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MinioNodeDiskOffline

Minio cluster node disk is offline

<details>
  <summary>Alert Rule</summary>

{{% rule "minio/minio-internal.yml" "MinioNodeDiskOffline" %}}

{{% comment %}}

```yaml
alert: MinioNodeDiskOffline
expr: minio_cluster_nodes_offline_total > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Minio node disk offline (instance {{ $labels.instance }})
    description: |-
        Minio cluster node disk is offline
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/minio-internal/MinioNodeDiskOffline.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the MinioNodeDiskOffline alert:

## Meaning
The MinioNodeDiskOffline alert is triggered when one or more disks in a Minio cluster node are offline. This alert is critical because it can cause data unavailability and potential data loss.

## Impact
The impact of this alert is high because it can:

* Cause data unavailability to users
* Lead to data loss if the offline disk is not brought back online promptly
* Affect the overall performance and reliability of the Minio cluster

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the Minio cluster node status using the Minio dashboard or CLI
2. Identify the specific disk that is offline using the LABELS output in the alert
3. Check the disk's health status using disk utility commands (e.g. `smartctl`, `fsck`)
4. Check the system logs for any error messages related to the disk or Minio node
5. Verify that the disk is properly connected and configured

## Mitigation
To mitigate the issue, follow these steps:

1. Immediately investigate the cause of the disk offline and take corrective action (e.g. replace the disk, check cables, etc.)
2. Bring the offline disk back online as soon as possible
3. Verify that the Minio cluster node is healthy and data is available
4. Consider adding additional redundancy to the Minio cluster to prevent similar issues in the future
5. Update the runbook and documentation to prevent similar issues in the future

Note: The mitigation steps may vary depending on the specific environment and setup. This runbook is meant to provide general guidance and may need to be tailored to the specific use case.