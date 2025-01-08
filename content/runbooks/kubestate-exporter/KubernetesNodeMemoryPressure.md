---
title: KubernetesNodeMemoryPressure
description: Troubleshooting for alert KubernetesNodeMemoryPressure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeMemoryPressure

Node {{ $labels.node }} has MemoryPressure condition

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeMemoryPressure" %}}

{{% comment %}}

```yaml
alert: KubernetesNodeMemoryPressure
expr: kube_node_status_condition{condition="MemoryPressure",status="true"} == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes memory pressure (node {{ $labels.node }})
    description: |-
        Node {{ $labels.node }} has MemoryPressure condition
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeMemoryPressure.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
