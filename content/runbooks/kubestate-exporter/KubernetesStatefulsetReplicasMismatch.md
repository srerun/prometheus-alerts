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


Here is a runbook for the Prometheus alert rule:

## Meaning

The `KubernetesStatefulsetReplicasMismatch` alert is triggered when the number of ready replicas in a StatefulSet does not match the expected number of replicas. This alert is raised when the `kube_statefulset_status_replicas_ready` metric does not equal the `kube_statefulset_status_replicas` metric for a period of 10 minutes.

## Impact

A mismatch in the number of replicas can lead to:

* Inconsistent application behavior
* Loss of data or inconsistent data
* Inability to scale or upgrade the application
* Potential security risks due to unexpected changes in the cluster

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the StatefulSet configuration to ensure it is correctly set up and the desired number of replicas is correct.
2. Verify that the `kube_statefulset_status_replicas_ready` and `kube_statefulset_status_replicas` metrics are correct and up-to-date.
3. Check the pod logs for any errors or issues that may be preventing the pods from being created or becoming ready.
4. Check the cluster events to see if there are any errors or issues related to the StatefulSet or pods.
5. Verify that the StatefulSet is not stuck in a rolling update or deployment.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the StatefulSet configuration and update it if necessary to ensure the correct number of replicas.
2. Verify that the pods are correctly configured and update them if necessary.
3. Check for any issues with the cluster or nodes that may be preventing the pods from being created or becoming ready.
4. Check for any errors or issues related to the StatefulSet or pods and resolve them.
5. If the issue persists, consider rolling back to a previous version of the StatefulSet or pods, or seeking assistance from a Kubernetes administrator.