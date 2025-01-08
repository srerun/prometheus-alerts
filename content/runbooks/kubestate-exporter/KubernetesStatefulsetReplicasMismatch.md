---
title: KubernetesStatefulsetReplicasMismatch
description: Troubleshooting for alert KubernetesStatefulsetReplicasMismatch
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesStatefulsetReplicasMismatch

## Meaning
[//]: # "Short paragraph that explains what the alert means"
StatefulSet does not match the expected number of replicas.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesStatefulsetReplicasMismatch" %}}

{{% comment %}}

```yaml
alert: KubernetesStatefulsetReplicasMismatch
expr: kube_statefulset_status_replicas_ready != kube_statefulset_status_replicas
for: 10m
labels:
    severity: warning
annotations:
    summary: Kubernetes StatefulSet replicas mismatch (instance {{ $labels.instance }})
    description: |-
        StatefulSet does not match the expected number of replicas.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesStatefulsetReplicasMismatch.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
