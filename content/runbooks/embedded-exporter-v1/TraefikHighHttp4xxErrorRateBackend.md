---
title: TraefikHighHttp4xxErrorRateBackend
description: Troubleshooting for alert TraefikHighHttp4xxErrorRateBackend
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikHighHttp4xxErrorRateBackend

Traefik backend 4xx error rate is above 5%

<details>
  <summary>Alert Rule</summary>

{{% rule "traefik/embedded-exporter-v1.yml" "TraefikHighHttp4xxErrorRateBackend" %}}

{{% comment %}}

```yaml
alert: TraefikHighHttp4xxErrorRateBackend
expr: sum(rate(traefik_backend_requests_total{code=~"4.*"}[3m])) by (backend) / sum(rate(traefik_backend_requests_total[3m])) by (backend) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: Traefik high HTTP 4xx error rate backend (instance {{ $labels.instance }})
    description: |-
        Traefik backend 4xx error rate is above 5%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v1/TraefikHighHttp4xxErrorRateBackend.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `TraefikHighHttp4xxErrorRateBackend`:

## Meaning

This alert is triggered when the rate of HTTP 4xx errors for a Traefik backend exceeds 5% of the total requests over a 3-minute window. This indicates that a significant number of requests to the backend are failing, which can have a negative impact on the user experience and application performance.

## Impact

* Users may experience failed requests or errors when interacting with the application.
* High error rates can lead to increased latency, decreased system performance, and potential cascading failures.
* If left unaddressed, this issue can result in revenue loss, damage to reputation, and decreased customer satisfaction.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected Traefik backend using the `backend` label in the alert.
2. Check the Traefik logs for errors related to the identified backend.
3. Verify that the backend service is operational and responding correctly.
4. Review recent changes to the application, backend, or Traefik configuration that may be causing the error rate to spike.
5. Use tools like `curl` or a web debugging proxy to simulate requests to the backend and reproduce the error.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and address the root cause of the high error rate, such as:
	* Fixing issues with the backend service or application.
	* Adjusting Traefik configuration or routing rules.
	* Implementing retries or circuit breakers to handle transient errors.
2. Implement a temporary fix to reduce the error rate, such as:
	* Load shedding or rate limiting to reduce the load on the backend.
	* Routing traffic to a different backend or instance.
3. Monitor the error rate and adjust the mitigation strategies as needed to ensure the issue is fully resolved.
4. Consider implementing additional monitoring and alerting to catch similar issues earlier in the future.