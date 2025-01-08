---
title: KubernetesNodeOutOfPodCapacity
description: Troubleshooting for alert KubernetesNodeOutOfPodCapacity
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeOutOfPodCapacity

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Node {{ $labels.node }} is out of pod capacity

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeOutOfPodCapacity" %}}

<!-- Rule when generated

```yaml
alert: KubernetesNodeOutOfPodCapacity
expr: sum by (node) ((kube_pod_status_phase{phase="Running"} == 1) + on(uid, instance) group_left(node) (0 * kube_pod_info{pod_template_hash=""})) / sum by (node) (kube_node_status_allocatable{resource="pods"}) * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes Node out of pod capacity (instance {{ $labels.instance }})
    description: |-
        Node {{ $labels.node }} is out of pod capacity
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeOutOfPodCapacity.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
