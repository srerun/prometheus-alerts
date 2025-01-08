---
title: KubernetesPodNotHealthy
description: Troubleshooting for alert KubernetesPodNotHealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesPodNotHealthy

Pod {{ $labels.namespace }}/{{ $labels.pod }} has been in a non-running state for longer than 15 minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesPodNotHealthy" %}}

{{% comment %}}

```yaml
alert: KubernetesPodNotHealthy
expr: sum by (namespace, pod) (kube_pod_status_phase{phase=~"Pending|Unknown|Failed"}) > 0
for: 15m
labels:
    severity: critical
annotations:
    summary: Kubernetes Pod not healthy ({{ $labels.namespace }}/{{ $labels.pod }})
    description: |-
        Pod {{ $labels.namespace }}/{{ $labels.pod }} has been in a non-running state for longer than 15 minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesPodNotHealthy.md

```

{{% /comment %}}

</details>


## Meaning

The KubernetesPodNotHealthy alert is triggered when a Kubernetes pod has been in a non-running state (Pending, Unknown, or Failed) for more than 15 minutes. This alert indicates a potential issue with the pod or the underlying cluster resources that is preventing the pod from running successfully.

## Impact

The impact of this alert can be significant, as it may indicate a failure in the application or service that the pod is providing. This can lead to:

* Downtime or unavailability of the application or service
* Loss of data or inconsistent data
* Increased latency or errors
* Negative impact on user experience or business operations

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the pod's status using `kubectl describe pod <pod_name> -n <namespace>`
2. Check the pod's logs using `kubectl logs <pod_name> -n <namespace>`
3. Check the cluster's resource usage and node status using `kubectl top nodes` and `kubectl describe node <node_name>`
4. Check for any ongoing deployments or rollouts that may be affecting the pod
5. Check the pod's configuration and deployment YAML files for any errors or inconsistencies

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the pod using `kubectl rollout restart deployment <deployment_name> -n <namespace>`
2. Check and resolve any underlying issues with the cluster resources or nodes
3. Verify that the pod's configuration and deployment YAML files are correct and up-to-date
4. Check for any ongoing deployments or rollouts and pause or cancel them if necessary
5. If the issue persists, consider escalating to a senior engineer or devops team for further assistance.

Note: The runbook URL provided in the alert annotations points to a more detailed runbook that can be used for further guidance and troubleshooting.