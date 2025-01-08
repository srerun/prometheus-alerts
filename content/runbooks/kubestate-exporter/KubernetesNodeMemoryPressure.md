---
title: KubernetesNodeMemoryPressure
description: Troubleshooting for alert KubernetesNodeMemoryPressure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeMemoryPressure

Node {{ $labels.node }} has MemoryPressure condition

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeMemoryPressure" %}}

{{% comment %}}

```yaml
alert: KubernetesNodeMemoryPressure
expr: kube_node_status_condition{condition="MemoryPressure",status="true"} == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes memory pressure (node {{ $labels.node }})
    description: |-
        Node {{ $labels.node }} has MemoryPressure condition
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeMemoryPressure.md

```

{{% /comment %}}

</details>


## Meaning

The KubernetesNodeMemoryPressure alert is triggered when a Kubernetes node is experiencing memory pressure, indicating that the node is running low on available memory. This can lead to performance issues, slow downs, and even node failures.

## Impact

The impact of this alert is critical, as memory pressure on a Kubernetes node can:

* Cause pods to restart or fail
* Impact the performance of applications running on the node
* Lead to node failures, causing downtime and service disruptions
* Affect the overall stability and reliability of the Kubernetes cluster

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the node's resource utilization using `kubectl top` or a monitoring tool like Prometheus.
2. Verify that the node is not experiencing any other issues, such as disk pressure or CPU throttling.
3. Review the node's configuration and resource allocation to ensure that it is suitable for the workloads running on it.
4. Investigate if there are any memory-intensive workloads or pods running on the node that may be contributing to the memory pressure.

## Mitigation

To mitigate the issue, perform the following steps:

1. Identify and terminate any unnecessary or idle pods consuming excessive memory.
2. Scale up or spread out memory-intensive workloads to reduce the load on the node.
3. Add more nodes to the cluster or upgrade existing nodes with more memory.
4. Implement memory-related constraints and limits in the pod configurations to prevent memory over-allocation.
5. Consider implementing a cluster autoscaler to dynamically adjust the node count based on resource utilization.

For further guidance, refer to the runbook at [https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeMemoryPressure.md](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeMemoryPressure.md)