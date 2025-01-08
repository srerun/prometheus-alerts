---
title: KubernetesHpaUnderutilized
description: Troubleshooting for alert KubernetesHpaUnderutilized
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesHpaUnderutilized

HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is constantly at minimum replicas for 50% of the time. Potential cost saving here.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesHpaUnderutilized" %}}

{{% comment %}}

```yaml
alert: KubernetesHpaUnderutilized
expr: max(quantile_over_time(0.5, kube_horizontalpodautoscaler_status_desired_replicas[1d]) == kube_horizontalpodautoscaler_spec_min_replicas) by (horizontalpodautoscaler) > 3
for: 0m
labels:
    severity: info
annotations:
    summary: Kubernetes HPA underutilized (instance {{ $labels.instance }})
    description: |-
        HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is constantly at minimum replicas for 50% of the time. Potential cost saving here.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesHpaUnderutilized.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the KubernetesHpaUnderutilized alert:

## Meaning

The KubernetesHpaUnderutilized alert is triggered when a Horizontal Pod Autoscaler (HPA) is constantly at minimum replicas for 50% of the time, indicating potential underutilization of resources. This alert is triggered when the desired replicas of the HPA are equal to the minimum replicas for more than 50% of the time in the past 24 hours.

## Impact

* Underutilization of resources (e.g., CPU, memory) allocated to the pods
* Potential cost savings by right-sizing the resources
* Inefficient use of cluster resources, which can lead to resource shortages for other applications

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HPA configuration to ensure that the minimum replicas are set correctly.
2. Verify that the application is not experiencing high traffic or usage patterns that require more resources.
3. Investigate the pod utilization metrics (e.g., CPU, memory) to determine if the pods are indeed underutilized.
4. Check for any recent changes to the application or deployment that may have caused the underutilization.

## Mitigation

To mitigate the issue, follow these steps:

1. Review and adjust the HPA configuration to set a more optimal minimum replica count.
2. Consider implementing a more dynamic scaling strategy based on resource utilization metrics.
3. Right-size the resources allocated to the pods to match the actual usage patterns.
4. Consider implementing cost-saving measures, such as resource pooling or bin packing, to optimize resource utilization.
5. Monitor the HPA and pod utilization metrics regularly to detect any future underutilization or overutilization issues.