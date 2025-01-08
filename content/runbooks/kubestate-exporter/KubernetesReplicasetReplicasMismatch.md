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


Here is a sample runbook for the Prometheus alert rule:

## Meaning

The `KubernetesReplicasetReplicasMismatch` alert is triggered when the number of replicas specified in a ReplicaSet's configuration (`kube_replicaset_spec_replicas`) does not match the number of ready replicas reported by the ReplicaSet's status (`kube_replicaset_status_ready_replicas`). This mismatch indicates that the ReplicaSet is not in a healthy state, and may lead to issues with the application or service being deployed.

## Impact

A mismatch between the specified and actual number of replicas can have several consequences, including:

* Under or over-provisioning of resources, leading to wasted resources or decreased performance
* Inconsistent application behavior, potentially leading to errors or downtime
* Difficulty in troubleshooting and debugging issues, due to the mismatch between expected and actual replica counts

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the ReplicaSet's configuration and status using `kubectl`:
```
kubectl get rs <replicaset_name> -n <namespace> -o yaml
```
2. Verify the number of replicas specified in the configuration (`spec.replicas`) and the number of ready replicas reported in the status (`status.readyReplicas`).
3. Investigate the ReplicaSet's events and logs to determine the cause of the mismatch:
```
kubectl describe rs <replicaset_name> -n <namespace>
```
4. Check the cluster's resource utilization and node availability to ensure that there are no underlying issues that could be contributing to the mismatch.

## Mitigation

To mitigate the issue, perform the following steps:

1. Update the ReplicaSet's configuration to reflect the correct number of replicas:
```
kubectl patch rs <replicaset_name> -n <namespace> -p='{"spec":{"replicas":<correct_number>}}'
```
2. Verify that the ReplicaSet's status updates to reflect the new configuration.
3. Monitor the ReplicaSet's status and events to ensure that the issue does not recur.
4. Consider implementing automated self-healing mechanisms, such as a Kubernetes PodDisruptionBudget, to help maintain the desired number of replicas.