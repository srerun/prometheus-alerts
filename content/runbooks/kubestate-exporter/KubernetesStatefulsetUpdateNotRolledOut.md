---
title: KubernetesStatefulsetUpdateNotRolledOut
description: Troubleshooting for alert KubernetesStatefulsetUpdateNotRolledOut
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesStatefulsetUpdateNotRolledOut

## Meaning
[//]: # "Short paragraph that explains what the alert means"
StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} update has not been rolled out.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesStatefulsetUpdateNotRolledOut" %}}

{{% comment %}}

```yaml
alert: KubernetesStatefulsetUpdateNotRolledOut
expr: max without (revision) (kube_statefulset_status_current_revision unless kube_statefulset_status_update_revision) * (kube_statefulset_replicas != kube_statefulset_status_replicas_updated)
for: 10m
labels:
    severity: warning
annotations:
    summary: Kubernetes StatefulSet update not rolled out ({{ $labels.namespace }}/{{ $labels.statefulset }})
    description: |-
        StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} update has not been rolled out.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesStatefulsetUpdateNotRolledOut.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
