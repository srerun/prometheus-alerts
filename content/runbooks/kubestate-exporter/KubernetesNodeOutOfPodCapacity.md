---
title: KubernetesNodeOutOfPodCapacity
description: Troubleshooting for alert KubernetesNodeOutOfPodCapacity
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeOutOfPodCapacity

Node {{ $labels.node }} is out of pod capacity

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeOutOfPodCapacity" %}}

{{% comment %}}

```yaml
alert: KubernetesNodeOutOfPodCapacity
expr: sum by (node) ((kube_pod_status_phase{phase="Running"} == 1) + on(uid, instance) group_left(node) (0 * kube_pod_info{pod_template_hash=""})) / sum by (node) (kube_node_status_allocatable{resource="pods"}) * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes Node out of pod capacity (instance {{ $labels.instance }})
    description: |-
        Node {{ $labels.node }} is out of pod capacity
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeOutOfPodCapacity.md

```

{{% /comment %}}

</details>


Here is a runbook for the KubernetesNodeOutOfPodCapacity alert:

## Meaning

The KubernetesNodeOutOfPodCapacity alert is triggered when a Kubernetes node is running low on available pod capacity. This means that the node is almost out of resources to schedule new pods, which can lead to pod failures and decreased cluster availability.

## Impact

If left unaddressed, this alert can have significant impacts on the availability and reliability of the Kubernetes cluster. Some possible consequences include:

* Pod failures and restarts
* Decreased cluster performance and responsiveness
* Increased latency and error rates
* Potential for cascading failures and outage

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected node(s) using the `node` label in the alert.
2. Check the node's current resource utilization using `kubectl top node <node_name>` or a similar command.
3. Verify that the node is not experiencing any hardware or software issues that could be contributing to the capacity issue.
4. Review the pod scheduling and deployment strategies to ensure they are efficient and not causing unnecessary pod churn.
5. Check the node's allocatable resources using `kubectl describe node <node_name> | grep Allocatable`.

## Mitigation

To mitigate the issue, follow these steps:

1. Scale up the node by adding more resources (e.g., increase the number of available pods) or upgrading the node's hardware.
2. Identify and terminate any unused or unnecessary pods to free up resources.
3. Implement pod disruption budgets to ensure that critical pods are not terminated during maintenance or upgrades.
4. Review and optimize pod resource requests and limits to ensure efficient resource utilization.
5. Consider implementing a cluster autoscaler to dynamically adjust node capacity based on demand.

Note: The specific mitigation steps may vary depending on the cluster configuration, deployment strategy, and application requirements. The runbook should be reviewed and updated accordingly.