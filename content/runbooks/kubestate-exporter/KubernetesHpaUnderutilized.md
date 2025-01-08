---
title: KubernetesHpaUnderutilized
description: Troubleshooting for alert KubernetesHpaUnderutilized
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesHpaUnderutilized

HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is constantly at minimum replicas for 50% of the time. Potential cost saving here.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesHpaUnderutilized" %}}

{{% comment %}}

```yaml
alert: KubernetesHpaUnderutilized
expr: max(quantile_over_time(0.5, kube_horizontalpodautoscaler_status_desired_replicas[1d]) == kube_horizontalpodautoscaler_spec_min_replicas) by (horizontalpodautoscaler) > 3
for: 0m
labels:
    severity: info
annotations:
    summary: Kubernetes HPA underutilized (instance {{ $labels.instance }})
    description: |-
        HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is constantly at minimum replicas for 50% of the time. Potential cost saving here.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesHpaUnderutilized.md

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
