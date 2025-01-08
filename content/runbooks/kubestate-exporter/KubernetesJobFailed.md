---
title: KubernetesJobFailed
description: Troubleshooting for alert KubernetesJobFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesJobFailed

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Job {{ $labels.namespace }}/{{ $labels.job_name }} failed to complete

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesJobFailed" %}}

{{% comment %}}

```yaml
alert: KubernetesJobFailed
expr: kube_job_status_failed > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes Job failed ({{ $labels.namespace }}/{{ $labels.job_name }})
    description: |-
        Job {{ $labels.namespace }}/{{ $labels.job_name }} failed to complete
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesJobFailed.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
