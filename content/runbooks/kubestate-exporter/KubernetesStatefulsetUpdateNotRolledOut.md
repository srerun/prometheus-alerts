---
title: KubernetesStatefulsetUpdateNotRolledOut
description: Troubleshooting for alert KubernetesStatefulsetUpdateNotRolledOut
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesStatefulsetUpdateNotRolledOut

StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} update has not been rolled out.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesStatefulsetUpdateNotRolledOut" %}}

{{% comment %}}

```yaml
alert: KubernetesStatefulsetUpdateNotRolledOut
expr: max without (revision) (kube_statefulset_status_current_revision unless kube_statefulset_status_update_revision) * (kube_statefulset_replicas != kube_statefulset_status_replicas_updated)
for: 10m
labels:
    severity: warning
annotations:
    summary: Kubernetes StatefulSet update not rolled out ({{ $labels.namespace }}/{{ $labels.statefulset }})
    description: |-
        StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} update has not been rolled out.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesStatefulsetUpdateNotRolledOut.md

```

{{% /comment %}}

</details>


## Meaning

The KubernetesStatefulsetUpdateNotRolledOut alert is triggered when a StatefulSet update has not been rolled out to all replicas within a 10-minute window. This alert indicates that there is an issue with the deployment or scaling of the StatefulSet, which may lead to inconsistent or outdated application behavior.

## Impact

The impact of this alert can be significant, as it may result in:

* Inconsistent application behavior across replicas
* Rollback or failure to apply updates, potentially leading to security vulnerabilities or feature gaps
* Increased latency or errors due to outdated or inconsistent application versions
* Decreased confidence in the reliability and scalability of the application

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the StatefulSet's rollout history to identify the update that is not rolling out.
2. Verify that the update is correctly configured and applied to the StatefulSet.
3. Check the replica's status and logs to identify any errors or issues that may be preventing the update from rolling out.
4. Verify that the Kubernetes cluster has sufficient resources (e.g., CPU, memory) to support the update.
5. Check for any network connectivity issues or firewalls that may be blocking the update.

## Mitigation

To mitigate this alert, follow these steps:

1. Roll back the update to a previous version that was successfully deployed.
2. Identify and fix any configuration issues or errors in the update.
3. Verify that the replica's status is updated correctly and the update is successfully applied.
4. Increase the replicas to ensure that the application is still functional, if necessary.
5. Monitor the StatefulSet's rollout history and logs to ensure that the update is rolling out correctly.

Remember to refer to the provided runbook link for more detailed steps and best practices for resolving this alert.