---
title: KubernetesStatefulsetDown
description: Troubleshooting for alert KubernetesStatefulsetDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesStatefulsetDown

StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} went down

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesStatefulsetDown" %}}

{{% comment %}}

```yaml
alert: KubernetesStatefulsetDown
expr: kube_statefulset_replicas != kube_statefulset_status_replicas_ready > 0
for: 1m
labels:
    severity: critical
annotations:
    summary: Kubernetes StatefulSet down ({{ $labels.namespace }}/{{ $labels.statefulset }})
    description: |-
        StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} went down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesStatefulsetDown.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `KubernetesStatefulsetDown`:

## Meaning
This alert indicates that a StatefulSet in a Kubernetes cluster has gone down, meaning that the number of replicas specified in the StatefulSet definition does not match the number of ready replicas. This can cause issues with application availability and data consistency.

## Impact
The impact of this alert is that the affected StatefulSet may not be able to serve requests, leading to:

* Downtime or reduced availability of the application
* Potential data loss or inconsistency
* Increased latency or errors for users

## Diagnosis
To diagnose the issue, follow these steps:

1. Check the StatefulSet logs for errors or exceptions
2. Verify that the StatefulSet pods are running and healthy
3. Check the Kubernetes cluster events for any errors or warnings related to the StatefulSet
4. Verify that the StatefulSet configuration is correct and up-to-date
5. Check the underlying infrastructure (e.g. nodes, disks) for any issues

## Mitigation
To mitigate the issue, follow these steps:

1. Check the StatefulSet status and verify that the replicas are running and healthy
2. Investigate and resolve any issues with the StatefulSet pods, such as pending or terminating pods
3. Scale up the StatefulSet to the desired number of replicas
4. Verify that the application is functioning correctly and serving requests
5. Perform a rolling update of the StatefulSet to ensure that all replicas are running the latest version of the application

Note: If the issue persists after following these steps, it may be necessary to escalate to a Kubernetes or application administrator for further assistance.