---
title: NginxHighHttp5xxErrorRate
description: Troubleshooting for alert NginxHighHttp5xxErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NginxHighHttp5xxErrorRate

Too many HTTP requests with status 5xx (> 5%)

<details>
  <summary>Alert Rule</summary>

{{% rule "nginx/knyar-nginx-exporter.yml" "NginxHighHttp5xxErrorRate" %}}

{{% comment %}}

```yaml
alert: NginxHighHttp5xxErrorRate
expr: sum(rate(nginx_http_requests_total{status=~"^5.."}[1m])) / sum(rate(nginx_http_requests_total[1m])) * 100 > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: Nginx high HTTP 5xx error rate (instance {{ $labels.instance }})
    description: |-
        Too many HTTP requests with status 5xx (> 5%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/knyar-nginx-exporter/NginxHighHttp5xxErrorRate.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "NginxHighHttp5xxErrorRate":

## Meaning

This alert is triggered when the rate of HTTP requests with a 5xx status code (indicating an error on the server side) exceeds 5% of the total HTTP requests over a 1-minute period. This may indicate a problem with the Nginx server or the application it is serving.

## Impact

* High error rates can lead to a poor user experience, as users may encounter errors when trying to access the application.
* If left unaddressed, this issue can result in a loss of user trust and revenue.
* It can also indicate a potential security issue or misconfiguration of the Nginx server.

## Diagnosis

* Check the Nginx error logs to identify the specific errors and their causes.
* Verify that the application is functioning correctly and not returning errors.
* Investigate recent changes to the Nginx configuration or the application code that may have caused the issue.
* Check the system resources (CPU, memory, disk space) to ensure they are not overwhelmed.
* Review the access logs to identify any patterns or trends in the errors.

## Mitigation

* Immediately investigate and address the root cause of the errors.
* Implement a temporary fix, such as increasing the resources available to the Nginx server or load balancing the traffic.
* Perform a rolling restart of the Nginx servers to ensure that any stuck processes are terminated.
* Consider implementing retries or circuit breakers in the application to reduce the impact of errors on users.
* Review and refine the Nginx configuration to prevent similar issues in the future.
* Monitor the error rate closely to ensure that it returns to a normal level.