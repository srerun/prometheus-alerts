---
title: KubernetesPersistentvolumeclaimPending
description: Troubleshooting for alert KubernetesPersistentvolumeclaimPending
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesPersistentvolumeclaimPending

PersistentVolumeClaim {{ $labels.namespace }}/{{ $labels.persistentvolumeclaim }} is pending

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesPersistentvolumeclaimPending" %}}

{{% comment %}}

```yaml
alert: KubernetesPersistentvolumeclaimPending
expr: kube_persistentvolumeclaim_status_phase{phase="Pending"} == 1
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes PersistentVolumeClaim pending ({{ $labels.namespace }}/{{ $labels.persistentvolumeclaim }})
    description: |-
        PersistentVolumeClaim {{ $labels.namespace }}/{{ $labels.persistentvolumeclaim }} is pending
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesPersistentvolumeclaimPending.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the KubernetesPersistentvolumeclaimPending alert:

## Meaning

The KubernetesPersistentvolumeclaimPending alert is triggered when a Persistent Volume Claim (PVC) in a Kubernetes cluster is in a pending state. This means that the PVC has been created and is waiting for a matching Persistent Volume (PV) to be provisioned and bound to it. This alert is warning-level, indicating that the issue needs to be addressed to ensure normal operation of the application.

## Impact

If a PVC remains in a pending state for an extended period, it can cause several issues:

* Applications that rely on the PVC may not function correctly or may experience errors.
* The cluster's storage capacity may be unavailable, leading to performance issues or even outages.
* Other PVCs or Pods may be blocked from using the same storage resources, causing a ripple effect in the cluster.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the PVC's status using `kubectl describe pvc <pvc-name> -n <namespace>`.
2. Verify that the PVC's storage class and access mode match the available PVs in the cluster.
3. Check the cluster's storage capacity and available resources using `kubectl get pv -o wide`.
4. Review the cluster's events and logs to identify any errors or issues related to the PVC or PV provisioning.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify that the cluster has sufficient storage capacity and available resources.
2. Check if there are any issues with the storage class or access mode configuration.
3. If necessary, create a new PV that matches the PVC's requirements.
4. Bind the PV to the PVC using `kubectl patch pvc <pvc-name> -n <namespace> -p '{"spec":{"volumeName":"<pv-name>"}}'`.
5. Verify that the PVC is now in a bound state using `kubectl get pvc <pvc-name> -n <namespace>`.

Additional resources:

* Kubernetes documentation: [Persistent Volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)
* Kubernetes documentation: [Persistent Volume Claims](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims)