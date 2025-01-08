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


Here is the runbook for the Prometheus alert rule `KubernetesVolumeOutOfDiskSpace`:

## Meaning

This alert is triggered when a Kubernetes volume has less than 10% of its total capacity available. This means that the volume is almost full, and immediate action is required to prevent disk space exhaustion, which can lead to pod evictions, data loss, and system instability.

## Impact

If this alert is not addressed promptly, it can lead to:

* Pod evictions and restarts, resulting in service disruptions
* Data loss or corruption due to insufficient disk space
* System instability and potential crashes
* Decreased overall cluster performance and reliability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the volume usage metrics: `kubelet_volume_stats_available_bytes` and `kubelet_volume_stats_capacity_bytes` to identify the affected volume(s).
2. Verify the pod(s) using the affected volume(s) and check their resource usage, such as container logs and disk usage.
3. Check the Kubernetes node(s) hosting the affected pod(s) for any signs of disk space issues or resource contention.
4. Review the cluster's capacity planning and resource allocation to identify potential bottlenecks.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and terminate or remove any unnecessary pods or resources consuming excessive disk space.
2. Expand the affected volume(s) or add new storage capacity to the cluster, if possible.
3. Implement volume quotas or resource limits to prevent similar issues in the future.
4. Consider migrating data to an external storage solution, such as an object store or a cloud-based storage service.
5. Monitor the affected volume(s) closely to ensure the issue is resolved and does not recur.

Remember to address the underlying causes of the issue and implement long-term solutions to prevent similar alerts in the future.