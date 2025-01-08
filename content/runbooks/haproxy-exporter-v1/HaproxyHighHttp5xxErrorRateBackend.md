---
title: HaproxyHighHttp5xxErrorRateBackend
description: Troubleshooting for alert HaproxyHighHttp5xxErrorRateBackend
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHighHttp5xxErrorRateBackend

Too many HTTP requests with status 5xx (> 5%) on backend {{ $labels.fqdn }}/{{ $labels.backend }}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyHighHttp5xxErrorRateBackend" %}}

{{% comment %}}

```yaml
alert: HaproxyHighHttp5xxErrorRateBackend
expr: sum by (backend) (rate(haproxy_server_http_responses_total{code="5xx"}[1m])) / sum by (backend) (rate(haproxy_server_http_responses_total[1m])) > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy high HTTP 5xx error rate backend (instance {{ $labels.instance }})
    description: |-
        Too many HTTP requests with status 5xx (> 5%) on backend {{ $labels.fqdn }}/{{ $labels.backend }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyHighHttp5xxErrorRateBackend.md

```

{{% /comment %}}

</details>


## Meaning

The `HaproxyHighHttp5xxErrorRateBackend` alert is triggered when the rate of HTTP requests with a 5xx status code on a specific backend exceeds 5% of the total requests to that backend over a 1-minute period. This indicates a significant error rate on the backend, which can impact the overall availability and performance of the service.

## Impact

* High error rates can lead to a degradation of service quality, causing frustration for users and potential revenue loss.
* Increased error rates can also put additional load on the system, leading to performance issues and potentially even crashes.
* If left unaddressed, high error rates can lead to a loss of customer trust and reputation damage.

## Diagnosis

1. Check the HAProxy logs to identify the specific backend and instance affected by the high error rate.
2. Investigate the root cause of the 5xx errors using the HAProxy logs, server logs, and application logs.
3. Verify that the backend is correctly configured and that there are no issues with the underlying infrastructure.
4. Check for any recent changes or deployments that may have introduced the error.
5. Use tools like `curl` or `wget` to test the backend and verify that the 5xx errors are still occurring.

## Mitigation

1. Identify and address the root cause of the 5xx errors, which may involve:
	* Restarting or reloading the backend service.
	* Adjusting server or application configurations.
	* Implementing retries or circuit breakers to reduce the impact of errors.
2. Implement a temporary fix to reduce the error rate, such as:
	* Decreasing the traffic to the affected backend.
	* Implementing a load balancing strategy to distribute traffic to other backends.
3. Perform a thorough investigation and create a permanent fix to prevent similar issues in the future.
4. Monitor the error rate and adjust the alert threshold as necessary to ensure it is effective in detecting issues.
5. Update the runbook and documentation to reflect the root cause and mitigation steps taken to address the issue.