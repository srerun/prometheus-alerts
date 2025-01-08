---
title: IstioKubernetesGatewayAvailabilityDrop
description: Troubleshooting for alert IstioKubernetesGatewayAvailabilityDrop
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioKubernetesGatewayAvailabilityDrop

Gateway pods have dropped. Inbound traffic will likely be affected.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioKubernetesGatewayAvailabilityDrop" %}}

{{% comment %}}

```yaml
alert: IstioKubernetesGatewayAvailabilityDrop
expr: min(kube_deployment_status_replicas_available{deployment="istio-ingressgateway", namespace="istio-system"}) without (instance, pod) < 2
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio Kubernetes gateway availability drop (instance {{ $labels.instance }})
    description: |-
        Gateway pods have dropped. Inbound traffic will likely be affected.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioKubernetesGatewayAvailabilityDrop.md

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
