---
title: KubernetesHpaMetricsUnavailability
description: Troubleshooting for alert KubernetesHpaMetricsUnavailability
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesHpaMetricsUnavailability

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is unable to collect metrics

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesHpaMetricsUnavailability" %}}

<!-- Rule when generated

```yaml
alert: KubernetesHpaMetricsUnavailability
expr: kube_horizontalpodautoscaler_status_condition{status="false", condition="ScalingActive"} == 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Kubernetes HPA metrics unavailability (instance {{ $labels.instance }})
    description: |-
        HPA {{ $labels.namespace }}/{{ $labels.horizontalpodautoscaler }} is unable to collect metrics
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesHpaMetricsUnavailability.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"
