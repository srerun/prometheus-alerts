---
title: KubernetesNodeNetworkUnavailable
description: Troubleshooting for alert KubernetesNodeNetworkUnavailable
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesNodeNetworkUnavailable

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Node {{ $labels.node }} has NetworkUnavailable condition

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesNodeNetworkUnavailable" %}}

<!-- Rule when generated

```yaml
alert: KubernetesNodeNetworkUnavailable
expr: kube_node_status_condition{condition="NetworkUnavailable",status="true"} == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Kubernetes Node network unavailable (instance {{ $labels.instance }})
    description: |-
        Node {{ $labels.node }} has NetworkUnavailable condition
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesNodeNetworkUnavailable.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
