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


Here is a sample runbook for the IstioKubernetesGatewayAvailabilityDrop alert:

## Meaning

The IstioKubernetesGatewayAvailabilityDrop alert is triggered when the number of available replicas for the Istio ingress gateway deployment in the istio-system namespace falls below 2. This indicates a potential issue with the gateway's availability, which may impact inbound traffic.

## Impact

* Inbound traffic to the cluster may be affected, leading to potential service disruptions or errors.
* Applications relying on the Istio ingress gateway may experience connectivity issues or timeouts.

## Diagnosis

* Check the Istio ingress gateway deployment in the istio-system namespace for any errors or issues.
* Verify that the replicas are running and available by checking the `kube_deployment_status_replicas_available` metric.
* Review the Kubernetes pod logs for any errors or issues that may be related to the gateway pods.
* Check the Istio gateway's configuration and ensure it is correct and up-to-date.

## Mitigation

* Investigate and resolve any issues with the Istio ingress gateway deployment, such as pod crashes or configuration errors.
* Scale the Istio ingress gateway deployment to ensure at least 2 replicas are running and available.
* Verify that the gateway's configuration is correct and up-to-date.
* Consider implementing automating scaling or self-healing mechanisms for the Istio ingress gateway deployment to prevent similar issues in the future.

Note: This is just a sample runbook, and you may need to adapt it to your specific environment and requirements.