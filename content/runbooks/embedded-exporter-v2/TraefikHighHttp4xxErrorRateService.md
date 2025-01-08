---
title: TraefikHighHttp4xxErrorRateService
description: Troubleshooting for alert TraefikHighHttp4xxErrorRateService
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikHighHttp4xxErrorRateService

Traefik service 4xx error rate is above 5%

<details>
  <summary>Alert Rule</summary>

{{% rule "traefik/embedded-exporter-v2.yml" "TraefikHighHttp4xxErrorRateService" %}}

{{% comment %}}

```yaml
alert: TraefikHighHttp4xxErrorRateService
expr: sum(rate(traefik_service_requests_total{code=~"4.*"}[3m])) by (service) / sum(rate(traefik_service_requests_total[3m])) by (service) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: Traefik high HTTP 4xx error rate service (instance {{ $labels.instance }})
    description: |-
        Traefik service 4xx error rate is above 5%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/TraefikHighHttp4xxErrorRateService.md

```

{{% /comment %}}

</details>


Here is a runbook for the TraefikHighHttp4xxErrorRateService alert:

## Meaning

The TraefikHighHttp4xxErrorRateService alert is triggered when the rate of HTTP 4xx errors for a Traefik service exceeds 5% of the total request rate over a 3-minute period. This alert indicates that there is a high error rate for a specific service, which may impact the availability and reliability of the service.

## Impact

A high HTTP 4xx error rate can have a significant impact on the service:

* Users may experience errors and failures when accessing the service
* The service may become unavailable or unstable
* The high error rate may indicate a underlying issue with the service, such as a configuration problem or a resource constraint

## Diagnosis

To diagnose the root cause of the high HTTP 4xx error rate, follow these steps:

1. Check the Traefik logs for errors and exceptions related to the service
2. Verify the service configuration and check for any recent changes
3. Check the resource utilization of the service, such as CPU and memory usage
4. Check for any network connectivity issues or firewall rules blocking traffic to the service
5. Use tools like curl or a web browser to test the service and reproduce the error

## Mitigation

To mitigate the high HTTP 4xx error rate, follow these steps:

1. Review and update the service configuration to ensure it is correct and optimal
2. Investigate and resolve any underlying issues causing the errors, such as resource constraints or network connectivity problems
3. Implement retry mechanisms or circuit breakers to handle temporary failures
4. Consider implementing rate limiting or traffic shaping to prevent overwhelming the service
5. Monitor the service closely to ensure the error rate returns to a normal level