---
title: KubernetesDeploymentGenerationMismatch
description: Troubleshooting for alert KubernetesDeploymentGenerationMismatch
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesDeploymentGenerationMismatch

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Deployment {{ $labels.namespace }}/{{ $labels.deployment }} has failed but has not been rolled back.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesDeploymentGenerationMismatch" %}}

{{% comment %}}

```yaml
alert: KubernetesDeploymentGenerationMismatch
expr: kube_deployment_status_observed_generation != kube_deployment_metadata_generation
for: 10m
labels:
    severity: critical
annotations:
    summary: Kubernetes Deployment generation mismatch ({{ $labels.namespace }}/{{ $labels.deployment }})
    description: |-
        Deployment {{ $labels.namespace }}/{{ $labels.deployment }} has failed but has not been rolled back.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesDeploymentGenerationMismatch.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
