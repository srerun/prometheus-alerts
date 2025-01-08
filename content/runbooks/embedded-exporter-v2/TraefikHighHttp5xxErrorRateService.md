---
title: TraefikHighHttp5xxErrorRateService
description: Troubleshooting for alert TraefikHighHttp5xxErrorRateService
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikHighHttp5xxErrorRateService

Traefik service 5xx error rate is above 5%

<details>
  <summary>Alert Rule</summary>

{{% rule "traefik/embedded-exporter-v2.yml" "TraefikHighHttp5xxErrorRateService" %}}

{{% comment %}}

```yaml
alert: TraefikHighHttp5xxErrorRateService
expr: sum(rate(traefik_service_requests_total{code=~"5.*"}[3m])) by (service) / sum(rate(traefik_service_requests_total[3m])) by (service) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: Traefik high HTTP 5xx error rate service (instance {{ $labels.instance }})
    description: |-
        Traefik service 5xx error rate is above 5%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v2/TraefikHighHttp5xxErrorRateService.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `TraefikHighHttp5xxErrorRateService`:

## Meaning

This alert is triggered when the HTTP 5xx error rate for a Traefik service exceeds 5% over a 3-minute period. This indicates that there is a critical issue with the service, causing a high number of internal server errors.

## Impact

The impact of this alert includes:

* Downtime or degraded performance for users of the affected service
* Potential loss of data or transactions
* Increased latency or errors for dependent services
* Reduced confidence in the overall system's reliability and stability

## Diagnosis

To diagnose the cause of this alert, follow these steps:

1. Check the Traefik logs for errors or exceptions that may be contributing to the high 5xx error rate.
2. Verify that the service is properly configured and that there are no issues with the underlying infrastructure.
3. Review the service's metrics to identify any trends or patterns that may be contributing to the error rate.
4. Check for any recent changes or deployments that may be causing the issue.
5. Investigate any potential dependencies or upstream services that may be experiencing issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve the underlying cause of the high 5xx error rate.
2. Implement a rollback or revert to a previous version of the service if recent changes are suspected to be the cause.
3. Increase the resources or capacity of the service to handle the current load.
4. Implement temporary workarounds, such as load balancing or circuit breakers, to reduce the impact of the issue.
5. Notify stakeholders and users of the affected service and provide regular updates on the status of the issue.

Remember to update the alert's annotations and labels as necessary to ensure that the correct teams and stakeholders are notified and informed.