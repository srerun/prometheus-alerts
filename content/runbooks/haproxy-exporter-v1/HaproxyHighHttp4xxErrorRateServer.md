---
title: HaproxyHighHttp4xxErrorRateServer
description: Troubleshooting for alert HaproxyHighHttp4xxErrorRateServer
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyHighHttp4xxErrorRateServer

Too many HTTP requests with status 4xx (> 5%) on server {{ $labels.server }}

<details>
  <summary>Alert Rule</summary>

{{% rule "haproxy/haproxy-exporter-v1.yml" "HaproxyHighHttp4xxErrorRateServer" %}}

{{% comment %}}

```yaml
alert: HaproxyHighHttp4xxErrorRateServer
expr: sum by (server) (rate(haproxy_server_http_responses_total{code="4xx"}[1m]) * 100) / sum by (server) (rate(haproxy_server_http_responses_total[1m])) > 5
for: 1m
labels:
    severity: critical
annotations:
    summary: HAProxy high HTTP 4xx error rate server (instance {{ $labels.instance }})
    description: |-
        Too many HTTP requests with status 4xx (> 5%) on server {{ $labels.server }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/haproxy-exporter-v1/HaproxyHighHttp4xxErrorRateServer.md

```

{{% /comment %}}

</details>


Here is a runbook for the HaproxyHighHttp4xxErrorRateServer alert:

## Meaning

The HaproxyHighHttp4xxErrorRateServer alert is triggered when the rate of HTTP requests with a 4xx status code (e.g. 404, 403, etc.) exceeds 5% of the total HTTP requests on a specific HAProxy server. This indicates that there may be a problem with the server or the upstream services it is proxying to, resulting in a high error rate.

## Impact

A high HTTP 4xx error rate can have several negative impacts:

* Users may experience errors or hangs when accessing the proxied services, leading to a poor user experience.
* The error rate may indicate a problem with the server or upstream services, which can lead to instability or downtime.
* In extreme cases, a high error rate can lead to a denial-of-service (DoS) scenario, where the server becomes unresponsive or crashes.

## Diagnosis

To diagnose the root cause of the high HTTP 4xx error rate, follow these steps:

1. Check the HAProxy server logs to identify the specific URL or services that are returning 4xx errors.
2. Verify that the upstream services are functioning correctly and responding as expected.
3. Check for any network connectivity issues or firewall rules that may be blocking requests.
4. Analyze the HAProxy configuration to ensure that it is correctly routing requests to the upstream services.
5. Review the system resource utilization (e.g. CPU, memory, disk space) to ensure that the server is not experiencing any resource constraints.

## Mitigation

To mitigate the high HTTP 4xx error rate, follow these steps:

1. Identify and fix any issues with the upstream services or network connectivity.
2. Adjust the HAProxy configuration to route requests to alternative upstream services or implement load balancing to distribute the load.
3. Implement rate limiting or traffic shaping to prevent excessive requests from overwhelming the server.
4. Consider implementing caching or content delivery networks (CDNs) to reduce the load on the server.
5. Monitor the error rate and system resource utilization to ensure that the mitigation steps are effective.