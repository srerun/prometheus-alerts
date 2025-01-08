---
title: KubernetesPodNotHealthy
description: Troubleshooting for alert KubernetesPodNotHealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesPodNotHealthy

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Pod {{ $labels.namespace }}/{{ $labels.pod }} has been in a non-running state for longer than 15 minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesPodNotHealthy" %}}

<!-- Rule when generated

```yaml
alert: KubernetesPodNotHealthy
expr: sum by (namespace, pod) (kube_pod_status_phase{phase=~"Pending|Unknown|Failed"}) > 0
for: 15m
labels:
    severity: critical
annotations:
    summary: Kubernetes Pod not healthy ({{ $labels.namespace }}/{{ $labels.pod }})
    description: |-
        Pod {{ $labels.namespace }}/{{ $labels.pod }} has been in a non-running state for longer than 15 minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesPodNotHealthy.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
