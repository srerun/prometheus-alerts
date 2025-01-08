---
title: IstioHigh4xxErrorRate
description: Troubleshooting for alert IstioHigh4xxErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioHigh4xxErrorRate

High percentage of HTTP 5xx responses in Istio (> 5%).

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioHigh4xxErrorRate" %}}

{{% comment %}}

```yaml
alert: IstioHigh4xxErrorRate
expr: sum(rate(istio_requests_total{reporter="destination", response_code=~"4.*"}[5m])) / sum(rate(istio_requests_total{reporter="destination"}[5m])) * 100 > 5
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio high 4xx error rate (instance {{ $labels.instance }})
    description: |-
        High percentage of HTTP 5xx responses in Istio (> 5%).
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioHigh4xxErrorRate.md

```

{{% /comment %}}

</details>


Here is the runbook for the IstioHigh4xxErrorRate alert:

## Meaning

The IstioHigh4xxErrorRate alert is triggered when the percentage of HTTP 4xx error responses in Istio exceeds 5% over a 5-minute period. This alert indicates that there is an issue with the Istio configuration or the underlying application that is causing a high rate of client-side errors.

## Impact

A high 4xx error rate can have a significant impact on the user experience and the overall reliability of the application. It can lead to:

* Increased latency and timeouts
* Decreased throughput and performance
* Frustrated users and potential revenue loss
* Increased load on upstream services and resources

## Diagnosis

To diagnose the root cause of the high 4xx error rate, follow these steps:

1. Check the Istio logs for any errors or issues related to the affected service.
2. Verify that the service is correctly configured and deployed.
3. Check the application logs for any errors or exceptions that may be causing the 4xx errors.
4. Use tools like `istioctl` or `kubectl` to inspect the Istio configuration and verify that it is correct.
5. Check for any known issues or bugs in the Istio version being used.

## Mitigation

To mitigate the high 4xx error rate, follow these steps:

1. Investigate and resolve any underlying issues with the service or application.
2. Verify that the Istio configuration is correct and up-to-date.
3. Check for any misconfigured or incorrect routing rules.
4. Consider implementing retries or circuit breakers to handle transient errors.
5. Monitor the error rate and performance metrics to ensure that the issue is resolved.

Additional resources:

* Istio documentation: [Troubleshooting](https://istio.io/latest/docs/ops/troubleshooting/)
* Istio documentation: [Configuring routing rules](https://istio.io/latest/docs/reference/config/networking/virtual-service/)