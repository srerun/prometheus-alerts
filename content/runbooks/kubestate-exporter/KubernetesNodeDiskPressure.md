---
title: KubernetesNodeDiskPressure
description: Troubleshooting for alert KubernetesNodeDiskPressure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeDiskPressure

Node {{ $labels.node }} has DiskPressure condition

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeDiskPressure" %}}

{{% comment %}}

```yaml
alert: KubernetesNodeDiskPressure
expr: kube_node_status_condition{condition="DiskPressure",status="true"} == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes disk pressure (node {{ $labels.node }})
    description: |-
        Node {{ $labels.node }} has DiskPressure condition
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeDiskPressure.md

```

{{% /comment %}}

</details>


Here is a runbook for the KubernetesNodeDiskPressure alert:

## Meaning

The KubernetesNodeDiskPressure alert is triggered when a Kubernetes node is experiencing disk pressure, meaning that the available disk space is below a certain threshold. This alert is critical because it can cause problems with pod scheduling, deployment, and overall cluster stability.

## Impact

The impact of disk pressure on a Kubernetes node can be significant:

* Pods may not be scheduled on the node due to lack of disk space
* Deployments may fail or be stuck in a pending state
* The node may become unresponsive or crash due to disk full errors
* Cluster performance and reliability may be affected

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the node's disk usage using `kubectl describe node <node_name>` or `kubectl top node <node_name> --resources`
2. Identify the pods and containers consuming the most disk space using `kubectl describe pod <pod_name>` or `kubectl exec -it <pod_name> -- df -h`
3. Check the node's disk capacity and available space using `kubectl describe node <node_name> | grep -i capacity`
4. Verify that there are no stuck or pending deployments using `kubectl get deployments --all-namespaces`

## Mitigation

To mitigate the issue, follow these steps:

1. **Identify and remove unnecessary files and directories**: Use `kubectl exec` to connect to the node and remove any unnecessary files and directories consuming disk space.
2. **Scale down or terminate resource-intensive pods**: Identify pods consuming excessive disk space and scale them down or terminate them if possible.
3. **Increase disk space**: If possible, increase the disk space available on the node by adding more storage or upgrading the node's hardware.
4. **Check and adjust pod resource requests and limits**: Review pod resource requests and limits to ensure they are reasonable and not causing excessive disk usage.
5. **Consider implementing a cleaning process**: Schedule a regular cleaning process to remove unnecessary files and directories from the node.

Remember to monitor the node's disk usage and adjust your mitigation strategy as needed.