---
title: KubernetesDaemonsetMisscheduled
description: Troubleshooting for alert KubernetesDaemonsetMisscheduled
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesDaemonsetMisscheduled

Some Pods of DaemonSet {{ $labels.namespace }}/{{ $labels.daemonset }} are running where they are not supposed to run

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesDaemonsetMisscheduled" %}}

{{% comment %}}

```yaml
alert: KubernetesDaemonsetMisscheduled
expr: kube_daemonset_status_number_misscheduled > 0
for: 1m
labels:
    severity: critical
annotations:
    summary: Kubernetes DaemonSet misscheduled ({{ $labels.namespace }}/{{ $labels.daemonset }})
    description: |-
        Some Pods of DaemonSet {{ $labels.namespace }}/{{ $labels.daemonset }} are running where they are not supposed to run
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesDaemonsetMisscheduled.md

```

{{% /comment %}}

</details>


Here is a runbook for the KubernetesDaemonsetMisscheduled alert rule:

## Meaning

The KubernetesDaemonsetMisscheduled alert is triggered when a DaemonSet has one or more Pods running on a node where they are not supposed to run. This can indicate a misconfiguration or a problem with the DaemonSet's scheduling.

## Impact

The impact of this alert can be significant, as it can lead to:

* Incorrect resource utilization and allocation
* Performance degradation
* Security vulnerabilities if Pods are running on nodes that lack necessary security configurations
* Unpredictable behavior and errors in the application

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the DaemonSet's configuration and ensure that it is correctly defined
2. Verify that the Pods are running on the correct nodes by checking the Pod's node affinity and tolerations
3. Check the node's resources and ensure that they meet the requirements of the DaemonSet's Pods
4. Verify that the DaemonSet's update strategy is correctly configured
5. Check the Kubernetes cluster's logs for any errors or warnings related to the DaemonSet or its Pods

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the misconfigured DaemonSet and update its configuration to ensure that it is correctly scheduling Pods
2. Check and update the node's resources to ensure that they meet the requirements of the DaemonSet's Pods
3. Verify that the DaemonSet's update strategy is correctly configured and apply the changes
4. Check and update the Pod's node affinity and tolerations to ensure that they are correctly configured
5. Consider restarting the affected Pods to ensure that they are running on the correct nodes
6. Monitor the DaemonSet's status and ensure that it is correctly scheduling Pods

Note: This runbook is a general guide and may need to be tailored to specific use cases and environments. It's also recommended to refer to the original runbook provided in the annotation `runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesDaemonsetMisscheduled.md` for more detailed and specific instructions.