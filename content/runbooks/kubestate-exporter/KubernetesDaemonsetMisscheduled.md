---
title: KubernetesDaemonsetMisscheduled
description: Troubleshooting for alert KubernetesDaemonsetMisscheduled
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesDaemonsetMisscheduled

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Some Pods of DaemonSet {{ $labels.namespace }}/{{ $labels.daemonset }} are running where they are not supposed to run

<details>
  <summary>Alert Rule</summary>

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
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/KubernetesDaemonsetMisscheduled

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
