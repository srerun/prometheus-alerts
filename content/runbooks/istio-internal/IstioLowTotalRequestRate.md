---
title: IstioLowTotalRequestRate
description: Troubleshooting for alert IstioLowTotalRequestRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioLowTotalRequestRate

Global request rate in the service mesh is unusually low.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioLowTotalRequestRate" %}}

{{% comment %}}

```yaml
alert: IstioLowTotalRequestRate
expr: sum(rate(istio_requests_total{reporter="destination"}[5m])) < 100
for: 2m
labels:
    severity: warning
annotations:
    summary: Istio low total request rate (instance {{ $labels.instance }})
    description: |-
        Global request rate in the service mesh is unusually low.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioLowTotalRequestRate.md

```

{{% /comment %}}

</details>


Here is a runbook for the IstioLowTotalRequestRate alert:

## Meaning

The IstioLowTotalRequestRate alert is triggered when the total request rate in the service mesh falls below 100 requests per 5 minutes. This indicates a significant decrease in traffic flowing through the service mesh, which may be a sign of a problem with the application, infrastructure, or Istio configuration.

## Impact

A low total request rate can have several impacts on the system:

* Reduced application performance or availability
* Impact on business KPIs, such as revenue or user engagement
* Increased latency or errors for users
* Potential security risks if the decrease in traffic is due to a misconfigured firewall or network policy

## Diagnosis

To diagnose the cause of the low total request rate, follow these steps:

1. Check the Istio dashboard for any errors or warnings related to request processing.
2. Verify that the application and its dependencies are functioning correctly.
3. Review the networking and firewall configurations to ensure they are not blocking traffic.
4. Check the load balancer and ingress gateway configurations to ensure they are functioning correctly.
5. Analyze the request traffic patterns to identify any unusual changes or trends.

## Mitigation

To mitigate the impact of the low total request rate, follow these steps:

1. Investigate and resolve any errors or warnings related to request processing in Istio.
2. Verify that the application and its dependencies are properly configured and scaled.
3. Adjust the networking and firewall configurations to allow traffic to flow properly.
4. Review and adjust the load balancer and ingress gateway configurations as needed.
5. Consider scaling up or out to increase capacity and handle increased traffic.

Additional resources:

* For more information on troubleshooting Istio, refer to the Istio documentation: <https://istio.io/latest/docs/ops/troubleshooting/>
* For more information on Prometheus alerting, refer to the Prometheus documentation: <https://prometheus.io/docs/alerting/latest/>