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


## Meaning

The KubernetesDeploymentGenerationMismatch alert is triggered when the observed generation of a Kubernetes deployment does not match the expected generation. This mismatch indicates that the deployment has failed, but the rolling back process has not been completed. This alert is critical because it can lead to unexpected behavior, errors, and potential data loss.

## Impact

The impact of this alert is high because it can cause:

* Deployment instability and potential downtime
* Data loss or inconsistency
* Rolling update failures
* Incorrect resource configuration

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Kubernetes deployment status using `kubectl get deployments <deployment_name> -n <namespace> -o yaml`
2. Verify the observed generation of the deployment using `kubectl get deployments <deployment_name> -n <namespace> -o jsonpath='{.status.observedGeneration}'`
3. Check the deployment's rolling update strategy and configuration
4. Review the deployment's events and logs to identify any errors or issues
5. Check the Kubernetes cluster's overall health and resource utilization

## Mitigation

To mitigate the issue, follow these steps:

1. Roll back the deployment to a stable version using `kubectl rollout undo deployments <deployment_name> -n <namespace>`
2. Verify that the deployment's observed generation matches the expected generation
3. Investigate and resolve any underlying issues that caused the deployment failure
4. Update the deployment's configuration and rolling update strategy as necessary
5. Monitor the deployment's status and performance to ensure stability and reliability