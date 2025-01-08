---
title: KubernetesDeploymentReplicasMismatch
description: Troubleshooting for alert KubernetesDeploymentReplicasMismatch
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesDeploymentReplicasMismatch

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Deployment {{ $labels.namespace }}/{{ $labels.deployment }} replicas mismatch

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesDeploymentReplicasMismatch" %}}

<!-- Rule when generated

```yaml
alert: KubernetesDeploymentReplicasMismatch
expr: kube_deployment_spec_replicas != kube_deployment_status_replicas_available
for: 10m
labels:
    severity: warning
annotations:
    summary: Kubernetes Deployment replicas mismatch ({{ $labels.namespace }}/{{ $labels.deployment }})
    description: |-
        Deployment {{ $labels.namespace }}/{{ $labels.deployment }} replicas mismatch
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesDeploymentReplicasMismatch.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
