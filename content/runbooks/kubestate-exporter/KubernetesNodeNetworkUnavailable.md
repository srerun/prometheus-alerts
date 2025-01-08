---
title: KubernetesNodeNetworkUnavailable
description: Troubleshooting for alert KubernetesNodeNetworkUnavailable
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeNetworkUnavailable

Node {{ $labels.node }} has NetworkUnavailable condition

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeNetworkUnavailable" %}}

{{% comment %}}

```yaml
alert: KubernetesNodeNetworkUnavailable
expr: kube_node_status_condition{condition="NetworkUnavailable",status="true"} == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes Node network unavailable (instance {{ $labels.instance }})
    description: |-
        Node {{ $labels.node }} has NetworkUnavailable condition
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeNetworkUnavailable.md

```

{{% /comment %}}

</details>


Here is a runbook for the KubernetesNodeNetworkUnavailable alert:

## Meaning

The KubernetesNodeNetworkUnavailable alert is triggered when a Kubernetes node reports a NetworkUnavailable condition, indicating that the node's network is not available. This can cause pods running on the node to become unreachable, leading to service disruptions and potential data loss.

## Impact

* Pods running on the affected node may become unreachable, leading to service disruptions and potential data loss.
* Applications relying on the node's network connectivity may experience errors or failures.
* Cluster performance and reliability may be impacted if the node is critical to the cluster's operations.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the node's status using `kubectl get node <node_name>`.
2. Verify the node's network configuration using `kubectl describe node <node_name>`.
3. Review the node's system logs for network-related errors using `kubectl logs -f <node_name>`.
4. Check the node's network interface configuration using `ip addr show` or `ifconfig`.
5. Verify that the node's network cable is connected and functioning properly.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the node's network interface using `ip link set <interface_name> down && ip link set <interface_name> up`.
2. Check if the node's network configuration is correct and update it if necessary.
3. Verify that the node's network cable is connected and functioning properly.
4. If the issue persists, consider rebooting the node or replacing the network hardware if necessary.
5. Once the node's network is available again, verify that pods are running correctly and services are accessible.

Remember to refer to the Kubernetes documentation and your organization's specific guidelines for troubleshooting and resolving node network issues.