---
title: KubernetesCronjobTooLong
description: Troubleshooting for alert KubernetesCronjobTooLong
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesCronjobTooLong

CronJob {{ $labels.namespace }}/{{ $labels.cronjob }} is taking more than 1h to complete.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesCronjobTooLong" %}}

{{% comment %}}

```yaml
alert: KubernetesCronjobTooLong
expr: time() - kube_cronjob_next_schedule_time > 3600
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes CronJob too long ({{ $labels.namespace }}/{{ $labels.cronjob }})
    description: |-
        CronJob {{ $labels.namespace }}/{{ $labels.cronjob }} is taking more than 1h to complete.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesCronjobTooLong.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
