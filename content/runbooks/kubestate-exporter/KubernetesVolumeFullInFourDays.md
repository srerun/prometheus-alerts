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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Volume under {{ $labels.namespace }}/{{ $labels.persistentvolumeclaim }} is expected to fill up within four days. Currently {{ $value | humanize }}% is available.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesVolumeFullInFourDays" %}}

<!-- Rule when generated

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

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
