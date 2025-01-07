---
title: KubernetesJobSlowCompletion
description: Troubleshooting for alert KubernetesJobSlowCompletion
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesJobSlowCompletion

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Kubernetes Job {{ $labels.namespace }}/{{ $labels.job_name }} did not complete in time.

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: KubernetesJobSlowCompletion
expr: kube_job_spec_completions - kube_job_status_succeeded - kube_job_status_failed > 0
for: 12h
labels:
    severity: critical
annotations:
    summary: Kubernetes job slow completion ({{ $labels.namespace }}/{{ $labels.job_name }})
    description: |-
        Kubernetes Job {{ $labels.namespace }}/{{ $labels.job_name }} did not complete in time.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/KubernetesJobSlowCompletion

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
