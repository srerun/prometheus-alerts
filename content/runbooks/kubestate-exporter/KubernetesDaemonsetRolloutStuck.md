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

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Pods of DaemonSet {{ $labels.namespace }}/{{ $labels.daemonset }} are not scheduled or not ready

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesDaemonsetRolloutStuck" %}}

<!-- Rule when generated

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

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
