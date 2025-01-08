---
title: KubernetesNodeNotReady
description: Troubleshooting for alert KubernetesNodeNotReady
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeNotReady

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Node {{ $labels.node }} has been unready for a long time

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeNotReady" %}}

<!-- Rule when generated

```yaml
alert: KubernetesNodeNotReady
expr: kube_node_status_condition{condition="Ready",status="true"} == 0
for: 10m
labels:
    severity: critical
annotations:
    summary: Kubernetes Node ready (node {{ $labels.node }})
    description: |-
        Node {{ $labels.node }} has been unready for a long time
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeNotReady.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
