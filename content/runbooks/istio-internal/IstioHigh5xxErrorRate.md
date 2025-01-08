---
title: IstioHigh5xxErrorRate
description: Troubleshooting for alert IstioHigh5xxErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioHigh5xxErrorRate

High percentage of HTTP 5xx responses in Istio (> 5%).

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioHigh5xxErrorRate" %}}

{{% comment %}}

```yaml
alert: IstioHigh5xxErrorRate
expr: sum(rate(istio_requests_total{reporter="destination", response_code=~"5.*"}[5m])) / sum(rate(istio_requests_total{reporter="destination"}[5m])) * 100 > 5
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio high 5xx error rate (instance {{ $labels.instance }})
    description: |-
        High percentage of HTTP 5xx responses in Istio (> 5%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioHigh5xxErrorRate.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `IstioHigh5xxErrorRate`:

## Meaning

The `IstioHigh5xxErrorRate` alert is triggered when the percentage of HTTP 5xx responses in Istio exceeds 5% over a 5-minute period. This alert indicates that there is a high rate of internal server errors within the Istio service mesh, which can impact the overall reliability and performance of the system.

## Impact

The impact of this alert can be significant, as it may cause:

* Increased error rates and degraded user experience
* Decreased system reliability and availability
* Difficulty troubleshooting and debugging issues
* Potential revenue loss or business impact due to system downtime

## Diagnosis

To diagnose the root cause of this alert, follow these steps:

1. Check the Istio dashboard and logs for any signs of component failures or errors.
2. Verify that the Istio control plane is healthy and functional.
3. Inspect the `istio_requests_total` metrics to identify the specific services or pods that are experiencing high error rates.
4. Check the application logs for any errors or exceptions that may be contributing to the high error rate.
5. Verify that the underlying infrastructure is healthy and has sufficient resources (e.g., CPU, memory, disk space).

## Mitigation

To mitigate the impact of this alert, follow these steps:

1. Identify and address any underlying component failures or errors in the Istio control plane or data plane.
2. Check for any configuration issues or misconfigurations in the Istio setup.
3. Implement circuit breakers or retries to handle temporary errors and prevent cascading failures.
4. Scale out the Istio control plane or data plane components to handle increased traffic or load.
5. Consider implementing additional monitoring and logging to gain more visibility into the system and detect issues earlier.

Remember to investigate and resolve the underlying cause of the alert to prevent it from happening again in the future.