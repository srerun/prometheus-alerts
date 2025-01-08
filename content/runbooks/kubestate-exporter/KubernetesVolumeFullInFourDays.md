---
title: KubernetesVolumeFullInFourDays
description: Troubleshooting for alert KubernetesVolumeFullInFourDays
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesVolumeFullInFourDays

Volume under {{ $labels.namespace }}/{{ $labels.persistentvolumeclaim }} is expected to fill up within four days. Currently {{ $value | humanize }}% is available.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesVolumeFullInFourDays" %}}

{{% comment %}}

```yaml
alert: KubernetesVolumeFullInFourDays
expr: predict_linear(kubelet_volume_stats_available_bytes[6h:5m], 4 * 24 * 3600) < 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Kubernetes Volume full in four days (instance {{ $labels.instance }})
    description: |-
        Volume under {{ $labels.namespace }}/{{ $labels.persistentvolumeclaim }} is expected to fill up within four days. Currently {{ $value | humanize }}% is available.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesVolumeFullInFourDays.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the KubernetesVolumeFullInFourDays alert:

## Meaning

The KubernetesVolumeFullInFourDays alert is triggered when a Kubernetes volume is predicted to run out of available space within the next four days. This alert is critical, as it can cause disruptions to applications and services running on the affected nodes.

## Impact

If left unaddressed, a full volume can cause:

* Disruptions to application and service availability
* Data loss or corruption
* Increased latency and errors
* Potential downstream impacts to dependent services and systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected namespace and Persistent Volume Claim (PVC) using the `{{ $labels.namespace }}` and `{{ $labels.persistentvolumeclaim }}` labels.
2. Check the current available space on the volume using the `kubelet_volume_stats_available_bytes` metric.
3. Investigate the rate of disk usage growth using the `predict_linear` function.
4. Review recent changes to the application or service using the volume, such as increased traffic or data ingestion.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the root cause of the disk usage growth and address it, if possible (e.g., reduce data ingestion, optimize application storage usage).
2. Increase the volume size or add additional storage capacity to the node.
3. Consider implementing data retention policies or cleaning up unnecessary data to free up space.
4. Monitor the volume usage closely and adjust the alert threshold as needed.

Remember to also refer to the GitHub runbook linked in the alert annotations for additional information and guidance.