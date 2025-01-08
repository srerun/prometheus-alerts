---
title: KubernetesStatefulsetDown
description: Troubleshooting for alert KubernetesStatefulsetDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesStatefulsetDown

## Meaning
[//]: # "Short paragraph that explains what the alert means"
StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} went down

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesStatefulsetDown" %}}

{{% comment %}}

```yaml
alert: KubernetesStatefulsetDown
expr: kube_statefulset_replicas != kube_statefulset_status_replicas_ready > 0
for: 1m
labels:
    severity: critical
annotations:
    summary: Kubernetes StatefulSet down ({{ $labels.namespace }}/{{ $labels.statefulset }})
    description: |-
        StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} went down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesStatefulsetDown.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
