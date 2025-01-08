---
title: KubernetesHpaScaleMaximum
description: Troubleshooting for alert KubernetesHpaScaleMaximum
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesHpaScaleMaximum

HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} has hit maximum number of desired pods

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesHpaScaleMaximum" %}}

{{% comment %}}

```yaml
alert: KubernetesHpaScaleMaximum
expr: (kube_horizontalpodautoscaler_status_desired_replicas >= kube_horizontalpodautoscaler_spec_max_replicas) and (kube_horizontalpodautoscaler_spec_max_replicas > 1) and (kube_horizontalpodautoscaler_spec_min_replicas != kube_horizontalpodautoscaler_spec_max_replicas)
for: 2m
labels:
    severity: info
annotations:
    summary: Kubernetes HPA scale maximum (instance {{ $labels.instance }})
    description: |-
        HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} has hit maximum number of desired pods
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesHpaScaleMaximum.md

```

{{% /comment %}}

</details>


## Meaning

The KubernetesHpaScaleMaximum alert is triggered when a Horizontal Pod Autoscaler (HPA) has reached its maximum allowed number of replicas, and the minimum and maximum replica counts are not equal. This means that the HPA has scaled up to the maximum allowed number of pods, and further scaling up is not possible.

## Impact

When an HPA reaches its maximum scale, it may lead to:

* Increased latency or errors in the application due to insufficient resources
* Inability to handle increased traffic or demand, potentially leading to downtime or lost revenue
* Inefficient resource utilization, as the HPA is unable to scale down when not needed

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HPA configuration: Verify that the `maxReplicas` and `minReplicas` values are correctly set in the HPA configuration.
2. Check the current pod count: Verify the current number of pods running in the deployment/replicaset.
3. Check the HPA status: Verify the current status of the HPA, including the `desiredReplicas` value.
4. Check the deployment/replicaset resource usage: Verify the current resource usage (e.g., CPU, memory) of the deployment/replicaset.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the `maxReplicas` value: If the current `maxReplicas` value is too low, consider increasing it to allow for further scaling.
2. Check for resource bottlenecks: Identify and address any resource bottlenecks (e.g., insufficient CPU, memory, or storage) that may be preventing the HPA from scaling further.
3. Optimize deployment/replicaset configuration: Review and optimize the deployment/replicaset configuration to ensure efficient resource utilization.
4. Consider cluster autoscaling: If the HPA is consistently reaching its maximum scale, consider enabling cluster autoscaling to dynamically adjust the cluster size based on demand.

For more information, refer to the [Kubernetes HPA Scale Maximum Runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesHpaScaleMaximum.md).