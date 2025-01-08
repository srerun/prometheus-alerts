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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Minio cluster node disk is offline

<details>
  <summary>Alert Rule</summary>

{{% rule "minio/minio-internal.yml" "MinioNodeDiskOffline" %}}

<!-- Rule when generated

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

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
