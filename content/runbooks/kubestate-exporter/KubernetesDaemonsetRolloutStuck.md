---
title: KubernetesDaemonsetRolloutStuck
description: Troubleshooting for alert KubernetesDaemonsetRolloutStuck
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesDaemonsetRolloutStuck

Some Pods of DaemonSet {{ $labels.namespace }}/{{ $labels.daemonset }} are not scheduled or not ready

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesDaemonsetRolloutStuck" %}}

{{% comment %}}

```yaml
alert: KubernetesDaemonsetRolloutStuck
expr: kube_daemonset_status_number_ready / kube_daemonset_status_desired_number_scheduled * 100 < 100 or kube_daemonset_status_desired_number_scheduled - kube_daemonset_status_current_number_scheduled > 0
for: 10m
labels:
    severity: warning
annotations:
    summary: Kubernetes DaemonSet rollout stuck ({{ $labels.namespace }}/{{ $labels.daemonset }})
    description: |-
        Some Pods of DaemonSet {{ $labels.namespace }}/{{ $labels.daemonset }} are not scheduled or not ready
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesDaemonsetRolloutStuck.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "KubernetesDaemonsetRolloutStuck":

## Meaning

The "KubernetesDaemonsetRolloutStuck" alert is triggered when a DaemonSet rollout is not progressing as expected. This can happen when some Pods of the DaemonSet are not scheduled or not ready, preventing the rollout from completing.

## Impact

A stuck DaemonSet rollout can have significant implications on the overall health and performance of the Kubernetes cluster. It can lead to:

* Inconsistent service availability
* Increased latency and errors
* Resource waste and inefficiency
* Difficulty in rolling out new features or fixes

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the DaemonSet status: Run `kubectl describe daemonset <daemonset_name> -n <namespace>` to view the current status of the DaemonSet.
2. Identify the stuck Pods: Run `kubectl get pods -l <daemonset_label> -n <namespace>` to list the Pods associated with the DaemonSet and identify which ones are not scheduled or not ready.
3. Check Pod logs: Run `kubectl logs <pod_name> -n <namespace>` to view the logs of the stuck Pods and identify any error messages or issues.
4. Verify node availability: Run `kubectl get nodes` to check if there are any issues with the nodes that are preventing Pods from being scheduled.
5. Check for resource constraints: Verify that there are sufficient resources (e.g. CPU, memory) available on the nodes to schedule the Pods.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve any underlying issues: Address any issues identified during diagnosis, such as node availability or resource constraints.
2. Force rollout: Run `kubectl rollout undo daemonset <daemonset_name> -n <namespace>` to force the rollout to retry.
3. Check and adjust DaemonSet configuration: Verify that the DaemonSet configuration is correct and adjust as needed.
4. Monitor DaemonSet status: Continuously monitor the DaemonSet status using `kubectl describe daemonset <daemonset_name> -n <namespace>` to ensure the rollout is progressing as expected.
5. Consider rolling back: If the issue persists, consider rolling back the DaemonSet to a previous version or configuration.