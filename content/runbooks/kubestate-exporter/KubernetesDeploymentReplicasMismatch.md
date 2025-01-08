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

Deployment {{ $labels.namespace }}/{{ $labels.deployment }} replicas mismatch

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesDeploymentReplicasMismatch" %}}

{{% comment %}}

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

{{% /comment %}}

</details>


Here is a runbook for the KubernetesDeploymentReplicasMismatch alert:

## Meaning

The KubernetesDeploymentReplicasMismatch alert is triggered when the number of replicas specified in a Kubernetes Deployment's spec does not match the number of available replicas reported by the Deployment's status. This mismatch can indicate a problem with the Deployment's scalability or availability.

## Impact

The impact of this alert can vary depending on the specific use case and requirements of the Deployment. Potential consequences include:

* Reduced application availability or performance due to insufficient replicas
* Increased resource utilization or costs due to excessive replicas
* Difficulty troubleshooting issues or identifying the root cause of problems

## Diagnosis

To diagnose the root cause of the KubernetesDeploymentReplicasMismatch alert, follow these steps:

1. Check the Deployment's spec and status to confirm the replica mismatch.
2. Review the Deployment's events and logs to identify any errors or issues that may be contributing to the mismatch.
3. Verify that the Deployment's configuration and resources (e.g., CPU, memory) are sufficient to support the desired number of replicas.
4. Check for any network connectivity or communication issues between the Kubernetes control plane and the Deployment's pods.

## Mitigation

To mitigate the KubernetesDeploymentReplicasMismatch alert, follow these steps:

1. Verify the intended replica count and adjust the Deployment's spec accordingly.
2. Check and resolve any underlying issues preventing the Deployment from scaling or updating correctly.
3. Consider implementing additional monitoring and alerting to detect and respond to replica mismatches more quickly.
4. Review and optimize the Deployment's configuration and resource allocation to ensure efficient use of resources.

Remember to refer to the Kubernetes documentation and your organization's specific guidelines for troubleshooting and resolving Deployment replica mismatches.