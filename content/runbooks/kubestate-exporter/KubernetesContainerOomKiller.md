---
title: KubernetesContainerOomKiller
description: Troubleshooting for alert KubernetesContainerOomKiller
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesContainerOomKiller

Container {{ $labels.container }} in pod {{ $labels.namespace }}/{{ $labels.pod }} has been OOMKilled {{ $value }} times in the last 10 minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesContainerOomKiller" %}}

{{% comment %}}

```yaml
alert: KubernetesContainerOomKiller
expr: (kube_pod_container_status_restarts_total - kube_pod_container_status_restarts_total offset 10m >= 1) and ignoring (reason) min_over_time(kube_pod_container_status_last_terminated_reason{reason="OOMKilled"}[10m]) == 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes container oom killer ({{ $labels.namespace }}/{{ $labels.pod }}:{{ $labels.container }})
    description: |-
        Container {{ $labels.container }} in pod {{ $labels.namespace }}/{{ $labels.pod }} has been OOMKilled {{ $value }} times in the last 10 minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesContainerOomKiller.md

```

{{% /comment %}}

</details>


## Meaning

The KubernetesContainerOomKiller alert rule is triggered when a container in a Kubernetes pod is terminated due to an Out-of-Memory (OOM) error. This means that the container has consumed too much memory, and the system has killed it to prevent further memory exhaustion.

## Impact

The impact of this alert is that the affected container will not be able to continue running, and any services or applications that rely on it may be disrupted or become unavailable. Additionally, if this is a recurring issue, it can lead to instability and downtime in the cluster.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the container logs to identify the root cause of the OOM error.
2. Verify that the container is not consuming excessive resources (e.g., CPU, memory) by checking the pod's resource utilization metrics (e.g., `kube_pod_container_resource_requests`, `kube_pod_container_resource_limits`).
3. Check the cluster's overall resource utilization to ensure that there is sufficient capacity to run the container.
4. Review the pod's configuration to ensure that the container is properly configured and deployed.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and address the root cause of the OOM error, such as fixing memory leaks or optimizing resource utilization.
2. Adjust the container's resource requests and limits to ensure that it has sufficient resources to run without exhausting memory.
3. Consider implementing memory-based horizontal pod autoscaling to ensure that the cluster can scale to meet increasing memory demands.
4. Implement monitoring and alerting for OOM errors to catch and respond to issues before they become critical.

Remember to refer to the runbook URL provided in the alert annotations for more detailed guidance and steps specific to your environment.