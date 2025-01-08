---
title: KubernetesNodeNotReady
description: Troubleshooting for alert KubernetesNodeNotReady
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeNotReady

Node {{ $labels.node }} has been unready for a long time

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeNotReady" %}}

{{% comment %}}

```yaml
alert: KubernetesNodeNotReady
expr: kube_node_status_condition{condition="Ready",status="true"} == 0
for: 10m
labels:
    severity: critical
annotations:
    summary: Kubernetes Node ready (node {{ $labels.node }})
    description: |-
        Node {{ $labels.node }} has been unready for a long time
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeNotReady.md

```

{{% /comment %}}

</details>


## Meaning

The KubernetesNodeNotReady alert is triggered when a Kubernetes node is not ready for an extended period of time (10 minutes in this case). This means that the node is not reporting as healthy and is not accepting new pods or performing other essential node functions. This can lead to service disruptions, pod failures, and other issues that can impact the overall health and stability of the Kubernetes cluster.

## Impact

The impact of a Kubernetes node being not ready can be significant, including:

* Service disruptions: Pods running on the affected node may fail, leading to service outages or reduced capacity.
* Resource waste: Other nodes in the cluster may need to take on additional workload, leading to resource waste and potential performance issues.
* Monitoring and logging issues: Node unavailability can disrupt monitoring and logging capabilities, making it harder to diagnose and troubleshoot issues.
* Security risks: Unhealthy nodes can pose security risks if they are not receiving security updates or are vulnerable to attacks.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Kubernetes node status: Run the command `kubectl get nodes` to view the current status of all nodes in the cluster. Look for the node that is reporting as "NotReady".
2. Check the node's condition: Run the command `kubectl describe node <node_name>` to view detailed information about the node's condition, including any error messages or events that may indicate the cause of the issue.
3. Check system logs: Review system logs for errors or warnings related to the node or its components, such as the kubelet or Docker.
4. Check for resource issues: Verify that the node has sufficient resources (e.g., CPU, memory, disk space) and that there are no issues with resource allocation.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the node: If the node is not responding, try restarting it to see if it comes back online.
2. Investigate and resolve underlying issues: Identify and address any underlying issues that may be causing the node to be not ready, such as software or hardware failures, network connectivity issues, or resource constraints.
3. Perform node maintenance: Perform routine maintenance tasks on the node, such as updating software or replacing faulty hardware.
4. Consider node replacement: If the node is persistently failing or cannot be restored, consider replacing it with a new node to ensure cluster stability and availability.

Remember to refer to the Kubernetes documentation and your organization's specific guidelines for further instructions on troubleshooting and resolving node issues.