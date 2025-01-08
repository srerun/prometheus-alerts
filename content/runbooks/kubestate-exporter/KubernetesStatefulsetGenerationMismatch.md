---
title: KubernetesStatefulsetGenerationMismatch
description: Troubleshooting for alert KubernetesStatefulsetGenerationMismatch
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesStatefulsetGenerationMismatch

StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} has failed but has not been rolled back.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesStatefulsetGenerationMismatch" %}}

{{% comment %}}

```yaml
alert: KubernetesStatefulsetGenerationMismatch
expr: kube_statefulset_status_observed_generation != kube_statefulset_metadata_generation
for: 10m
labels:
    severity: critical
annotations:
    summary: Kubernetes StatefulSet generation mismatch ({{ $labels.namespace }}/{{ $labels.statefulset }})
    description: |-
        StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} has failed but has not been rolled back.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesStatefulsetGenerationMismatch.md

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
