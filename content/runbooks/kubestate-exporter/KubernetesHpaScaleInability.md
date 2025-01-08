---
title: KubernetesHpaScaleInability
description: Troubleshooting for alert KubernetesHpaScaleInability
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesHpaScaleInability

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is unable to scale

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesHpaScaleInability" %}}

<!-- Rule when generated

```yaml
alert: KubernetesHpaScaleInability
expr: (kube_horizontalpodautoscaler_spec_max_replicas - kube_horizontalpodautoscaler_status_desired_replicas) * on (horizontalpodautoscaler,namespace) (kube_horizontalpodautoscaler_status_condition{condition="ScalingLimited", status="true"} == 1) == 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes HPA scale inability (instance {{ $labels.instance }})
    description: |-
        HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is unable to scale
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesHpaScaleInability.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
