---
title: KubernetesNodeDiskPressure
description: Troubleshooting for alert KubernetesNodeDiskPressure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeDiskPressure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Node {{ $labels.node }} has DiskPressure condition

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeDiskPressure" %}}

{{% comment %}}

```yaml
alert: KubernetesNodeDiskPressure
expr: kube_node_status_condition{condition="DiskPressure",status="true"} == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes disk pressure (node {{ $labels.node }})
    description: |-
        Node {{ $labels.node }} has DiskPressure condition
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeDiskPressure.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
