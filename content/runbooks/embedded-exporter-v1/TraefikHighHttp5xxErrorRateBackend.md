---
title: TraefikHighHttp5xxErrorRateBackend
description: Troubleshooting for alert TraefikHighHttp5xxErrorRateBackend
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# TraefikHighHttp5xxErrorRateBackend

Traefik backend 5xx error rate is above 5%

<details>
  <summary>Alert Rule</summary>

{{% rule "traefik/embedded-exporter-v1.yml" "TraefikHighHttp5xxErrorRateBackend" %}}

{{% comment %}}

```yaml
alert: TraefikHighHttp5xxErrorRateBackend
expr: sum(rate(traefik_backend_requests_total{code=~"5.*"}[3m])) by (backend) / sum(rate(traefik_backend_requests_total[3m])) by (backend) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: Traefik high HTTP 5xx error rate backend (instance {{ $labels.instance }})
    description: |-
        Traefik backend 5xx error rate is above 5%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/embedded-exporter-v1/TraefikHighHttp5xxErrorRateBackend.md

```

{{% /comment %}}

</details>


## Meaning

The TraefikHighHttp5xxErrorRateBackend alert is triggered when the 5xx error rate for a Traefik backend exceeds 5% over a 3-minute period. This indicates that the backend is experiencing a high rate of internal server errors, which can impact the availability and performance of the application.

## Impact

* High 5xx error rates can lead to:
	+ Increased latency and response times
	+ Decreased application availability
	+ Frustrated users and potential loss of revenue
	+ Increased load on the backend, leading to resource exhaustion
* If left unaddressed, this issue can cause a cascading failure of the application and related services.

## Diagnosis

* Check the Traefik dashboard and logs for errors and exceptions related to the backend
* Investigate the backend service for issues, such as:
	+ High CPU or memory usage
	+ Slow database queries or connectivity issues
	+ Configuration errors or misconfigurations
	+ Insufficient resources or capacity
* Review the Traefik configuration for any misconfigurations or issues with the backend routing
* Verify that the backend service is properly deployed and configured

## Mitigation

* Immediately investigate and address the root cause of the 5xx errors in the backend service
* Implement short-term mitigations, such as:
	+ Temporarily reducing the load on the backend by load-shedding or rate-limiting requests
	+ Increasing the capacity of the backend service by adding more resources or instances
	+ Enabling caching or buffering to reduce the load on the backend
* Work on long-term solutions, such as:
	+ Optimizing database queries and database performance
	+ Improving the architecture and design of the backend service
	+ Implementing monitoring and alerting to detect issues earlier
	+ Conducting regular performance testing and reviews to identify bottlenecks and areas for improvement