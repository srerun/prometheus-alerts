---
title: KubernetesVolumeOutOfDiskSpace
description: Troubleshooting for alert KubernetesVolumeOutOfDiskSpace
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesVolumeOutOfDiskSpace

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Volume is almost full (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesVolumeOutOfDiskSpace" %}}

{{% comment %}}

```yaml
alert: KubernetesVolumeOutOfDiskSpace
expr: kubelet_volume_stats_available_bytes / kubelet_volume_stats_capacity_bytes * 100 < 10
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes Volume out of disk space (instance {{ $labels.instance }})
    description: |-
        Volume is almost full (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesVolumeOutOfDiskSpace.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
