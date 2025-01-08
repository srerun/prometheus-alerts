---
title: HaproxyHighHttp5xxErrorRateServer
description: Troubleshooting for alert HaproxyHighHttp5xxErrorRateServer
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHighHttp5xxErrorRateServer

Too many HTTP requests with status 5xx (> 5%) on server {{ $labels.server }}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyHighHttp5xxErrorRateServer" %}}

{{% comment %}}

```yaml
alert: HaproxyHighHttp5xxErrorRateServer
expr: sum by (server) (rate(haproxy_server_http_responses_total{code="5xx"}[1m]) * 100) / sum by (server) (rate(haproxy_server_http_responses_total[1m])) > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy high HTTP 5xx error rate server (instance {{ $labels.instance }})
    description: |-
        Too many HTTP requests with status 5xx (> 5%) on server {{ $labels.server }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyHighHttp5xxErrorRateServer.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `HaproxyHighHttp5xxErrorRateServer`:

## Meaning

The `HaproxyHighHttp5xxErrorRateServer` alert is triggered when the rate of HTTP requests with a 5xx status code on a specific HAProxy server exceeds 5% of the total HTTP requests on that server. This indicates a high error rate for the server, which can impact the overall availability and performance of the application.

## Impact

* High error rates can lead to a poor user experience, resulting in decreased customer satisfaction and potential revenue loss.
* The increased error rate may indicate underlying issues with the application or backend services, which can lead to further escalation and downtime if left unaddressed.
* The alert may trigger unnecessary escalations and notifications to teams, causing fatigue and decreasing the overall effectiveness of the monitoring system.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the HAProxy server logs for errors and exceptions related to the 5xx status code.
2. Investigate the backend services and application logs for errors and exceptions that may be contributing to the high error rate.
3. Verify that the HAProxy server is properly configured and that there are no issues with the load balancing or routing.
4. Check the system resources (e.g., CPU, memory, disk usage) to ensure that the server is not experiencing any resource constraints.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and address the root cause of the high error rate, whether it's an issue with the application, backend services, or HAProxy configuration.
2. Implement temporary mitigations, such as redirecting traffic to alternative servers or implementing rate limiting, to reduce the impact on users.
3. Perform a rolling restart of the HAProxy server to ensure that the issue is not related to a specific process or thread.
4. Consider implementing additional error handling and retry mechanisms to improve the overall resilience of the application.

If the issue persists, please refer to the [HAProxy documentation](https://www.haproxy.org/documentation/) and [ Prometheus documentation](https://prometheus.io/docs/alerting/rules/) for further guidance.