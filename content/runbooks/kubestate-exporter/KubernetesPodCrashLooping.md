---
title: KubernetesPodCrashLooping
description: Troubleshooting for alert KubernetesPodCrashLooping
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesPodCrashLooping

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Pod {{ $labels.namespace }}/{{ $labels.pod }} is crash looping

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesPodCrashLooping" %}}

<!-- Rule when generated

```yaml
alert: KubernetesPodCrashLooping
expr: increase(kube_pod_container_status_restarts_total[1m]) > 3
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes pod crash looping ({{ $labels.namespace }}/{{ $labels.pod }})
    description: |-
        Pod {{ $labels.namespace }}/{{ $labels.pod }} is crash looping
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesPodCrashLooping.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
