---
title: KubernetesStatefulsetGenerationMismatch
description: Troubleshooting for alert KubernetesStatefulsetGenerationMismatch
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesStatefulsetGenerationMismatch

StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} has failed but has not been rolled back.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesStatefulsetGenerationMismatch" %}}

{{% comment %}}

```yaml
alert: KubernetesStatefulsetGenerationMismatch
expr: kube_statefulset_status_observed_generation != kube_statefulset_metadata_generation
for: 10m
labels:
    severity: critical
annotations:
    summary: Kubernetes StatefulSet generation mismatch ({{ $labels.namespace }}/{{ $labels.statefulset }})
    description: |-
        StatefulSet {{ $labels.namespace }}/{{ $labels.statefulset }} has failed but has not been rolled back.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesStatefulsetGenerationMismatch.md

```

{{% /comment %}}

</details>


Here is a runbook for the KubernetesStatefulsetGenerationMismatch alert:

## Meaning

The KubernetesStatefulsetGenerationMismatch alert is triggered when the observed generation of a StatefulSet does not match the expected generation. This can occur when a StatefulSet update fails, causing the observed generation to diverge from the expected generation.

## Impact

If left unresolved, a StatefulSet generation mismatch can lead to:

* Unexpected behavior or errors in the application
* Inconsistent data or state across replicas
* Difficulty in rolling back to a previous version
* Potential data loss or corruption

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the StatefulSet's status using `kubectl describe statefulset <statefulset_name> -n <namespace>`
2. Verify the observed generation and expected generation values
3. Check the StatefulSet's update history using `kubectl rollout history statefulset <statefulset_name> -n <namespace>`
4. Inspect the Pods' status and logs for any errors or issues

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the root cause of the generation mismatch (e.g., failed update, network issues, etc.)
2. Roll back the StatefulSet to a previous version using `kubectl rollout undo statefulset <statefulset_name> -n <namespace>`
3. Verify the StatefulSet's status and generation values after the rollback
4. Investigate and address any underlying issues that caused the generation mismatch
5. Consider implementing additional monitoring and logging to detect similar issues in the future