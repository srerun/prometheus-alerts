---
title: KubernetesHpaScaleInability
description: Troubleshooting for alert KubernetesHpaScaleInability
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesHpaScaleInability

HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is unable to scale

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesHpaScaleInability" %}}

{{% comment %}}

```yaml
alert: KubernetesHpaScaleInability
expr: (kube_horizontalpodautoscaler_spec_max_replicas - kube_horizontalpodautoscaler_status_desired_replicas) * on (horizontalpodautoscaler,namespace) (kube_horizontalpodautoscaler_status_condition{condition="ScalingLimited", status="true"} == 1) == 0
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes HPA scale inability (instance {{ $labels.instance }})
    description: |-
        HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is unable to scale
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesHpaScaleInability.md

```

{{% /comment %}}

</details>


## Meaning

The KubernetesHpaScaleInability alert is triggered when a Horizontal Pod Autoscaler (HPA) is unable to scale. This occurs when the desired number of replicas is less than or equal to the current number of replicas, and the HPA is in a "ScalingLimited" state. This alert indicates that the HPA is not able to scale up or down as needed, which can lead to issues with pod availability and resource utilization.

## Impact

The impact of this alert is that the affected HPA will not be able to adjust the number of replicas to match changing workload demands. This can result in:

* Insufficient resources to handle increased workload, leading to performance degradation or errors
* Waste of resources due to over-provisioning, leading to increased costs and inefficiency
* Potential pod unavailability or crashes due to resource constraints

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the HPA configuration and ensure that the `spec.maxReplicas` value is set correctly.
2. Verify that the `status.desiredReplicas` value is up-to-date and accurate.
3. Investigate any external factors that may be preventing the HPA from scaling, such as:
	* Resource constraints (e.g., CPU, memory, or disk space)
	* Network connectivity issues
	* Other Kubernetes components or controllers that may be interfering with the HPA
4. Review the HPA's event history to identify any errors or warnings related to scaling.

## Mitigation

To mitigate the effects of this alert, take the following steps:

1. Adjust the `spec.maxReplicas` value to ensure it is greater than the current number of replicas.
2. Verify that the HPA is properly configured to scale based on the desired metrics (e.g., CPU utilization).
3. Identify and address any external factors preventing the HPA from scaling (e.g., resource constraints, network issues).
4. Consider implementing additional monitoring and logging to detect future scaling limitations.
5. Consider implementing a fallback or contingency plan to handle situations where the HPA is unable to scale.

Note: This runbook is a general guideline, and the specific steps may vary depending on the environment and configuration. It is recommended to review the specific HPA configuration and environment to determine the best course of action.