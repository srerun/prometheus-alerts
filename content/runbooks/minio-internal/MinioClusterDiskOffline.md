---
title: MinioClusterDiskOffline
description: Troubleshooting for alert MinioClusterDiskOffline
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# MinioClusterDiskOffline

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Minio cluster disk is offline

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/MinioClusterDiskOffline

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
