---
title: KubernetesPersistentvolumeError
description: Troubleshooting for alert KubernetesPersistentvolumeError
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesPersistentvolumeError

Persistent volume {{ $labels.persistentvolume }} is in bad state

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesPersistentvolumeError" %}}

{{% comment %}}

```yaml
alert: KubernetesPersistentvolumeError
expr: kube_persistentvolume_status_phase{phase=~"Failed|Pending", job="kube-state-metrics"} > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Kubernetes PersistentVolumeClaim pending ({{ $labels.namespace }}/{{ $labels.persistentvolumeclaim }})
    description: |-
        Persistent volume {{ $labels.persistentvolume }} is in bad state
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesPersistentvolumeError.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the KubernetesPersistentvolumeError alert:

## Meaning

The KubernetesPersistentvolumeError alert is triggered when a Persistent Volume (PV) in a Kubernetes cluster is in a failed or pending state. This can indicate a problem with the PV, the underlying storage system, or the Kubernetes cluster itself.

## Impact

A failed or pending PV can cause disruptions to applications and services that rely on it for storage. This can lead to data loss, application crashes, or even complete system failures. The impact can be severe, especially if the affected PV is critical to the operation of the cluster or application.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Kubernetes cluster logs for errors related to the PV and its underlying storage system.
2. Verify that the PV is correctly configured and provisioned.
3. Check the status of the PV using the `kubectl` command: `kubectl describe pv <pv-name> -n <namespace>`
4. Investigate any recent changes to the PV, such as deployments or upgrades, that may have caused the issue.
5. Review the Kubernetes cluster's storage capacity and utilization to ensure that there are no resource constraints.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and resolve the underlying cause of the PV failure or pending state.
2. If the PV is stuck in a pending state, try deleting and recreating the PV.
3. If the PV is failed, try repairing or replacing the underlying storage system.
4. Verify that the PV is correctly provisioned and configured after resolution.
5. Monitor the PV's status and the cluster's logs to ensure that the issue has been fully resolved.

Remember to also review and update any dependent applications or services to ensure they are functioning correctly after the PV issue is resolved.