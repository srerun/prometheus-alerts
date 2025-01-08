---
title: KubernetesReplicasetReplicasMismatch
description: Troubleshooting for alert KubernetesReplicasetReplicasMismatch
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesReplicasetReplicasMismatch

## Meaning
[//]: # "Short paragraph that explains what the alert means"
ReplicaSet {{ $labels.namespace }}/{{ $labels.replicaset }} replicas mismatch

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesReplicasetReplicasMismatch" %}}

{{% comment %}}

```yaml
alert: KubernetesReplicasetReplicasMismatch
expr: kube_replicaset_spec_replicas != kube_replicaset_status_ready_replicas
for: 10m
labels:
    severity: warning
annotations:
    summary: Kubernetes ReplicasSet mismatch ({{ $labels.namespace }}/{{ $labels.replicaset }})
    description: |-
        ReplicaSet {{ $labels.namespace }}/{{ $labels.replicaset }} replicas mismatch
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesReplicasetReplicasMismatch.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
