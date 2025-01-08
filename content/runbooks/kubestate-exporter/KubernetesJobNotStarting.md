---
title: KubernetesJobNotStarting
description: Troubleshooting for alert KubernetesJobNotStarting
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesJobNotStarting

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Job {{ $labels.namespace }}/{{ $labels.job_name }} did not start for 10 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesJobNotStarting" %}}

{{% comment %}}

```yaml
alert: KubernetesJobNotStarting
expr: kube_job_status_active == 0 and kube_job_status_failed == 0 and kube_job_status_succeeded == 0 and (time() - kube_job_status_start_time) > 600
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes Job not starting ({{ $labels.namespace }}/{{ $labels.job_name }})
    description: |-
        Job {{ $labels.namespace }}/{{ $labels.job_name }} did not start for 10 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesJobNotStarting.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
