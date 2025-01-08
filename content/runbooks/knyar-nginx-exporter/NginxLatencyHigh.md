---
title: NginxLatencyHigh
description: Troubleshooting for alert NginxLatencyHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NginxLatencyHigh

Nginx p99 latency is higher than 3 seconds

<details>
  <summary>Alert Rule</summary>

{{% rule "nginx/knyar-nginx-exporter.yml" "NginxLatencyHigh" %}}

{{% comment %}}

```yaml
alert: NginxLatencyHigh
expr: histogram_quantile(0.99, sum(rate(nginx_http_request_duration_seconds_bucket[2m])) by (host, node, le)) > 3
for: 2m
labels:
    severity: warning
annotations:
    summary: Nginx latency high (instance {{ $labels.instance }})
    description: |-
        Nginx p99 latency is higher than 3 seconds
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/knyar-nginx-exporter/NginxLatencyHigh.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `NginxLatencyHigh`:

## Meaning

The `NginxLatencyHigh` alert is triggered when the 99th percentile of Nginx request latency exceeds 3 seconds over a 2-minute window. This indicates that a significant number of requests are experiencing high latency, which can impact the responsiveness and overall user experience of the application.

## Impact

High Nginx latency can have several negative impacts:

* Slow response times can lead to frustrated users and decreased engagement.
* Increased latency can cause requests to timeout, resulting in errors and failed transactions.
* High latency can also increase the load on Nginx and other upstream services, leading to potential resource exhaustion and further performance degradation.

## Diagnosis

To diagnose the root cause of high Nginx latency, follow these steps:

1. Check the Nginx error logs for any errors or warnings that may indicate the source of the latency.
2. Verify that the Nginx configuration is optimal and tuned for performance.
3. Investigate the performance of upstream services and databases to ensure they are not contributing to the latency.
4. Check the system metrics (e.g., CPU, memory, disk usage) to ensure that the Nginx server or underlying infrastructure is not resource-constrained.
5. Analyze the request patterns and traffic trends to identify any unusual or anomalous behavior.

## Mitigation

To mitigate high Nginx latency, follow these steps:

1. **Optimize Nginx configuration**: Review and optimize the Nginx configuration to ensure it is optimized for performance. Consider tuning parameters such as worker processes, keepalive timeouts, and buffer sizes.
2. **Scale Nginx horizontally**: If the Nginx server is resource-constrained, consider scaling out to additional instances to distribute the load.
3. **Optimize upstream services**: Verify that upstream services and databases are optimized for performance and can handle the request load.
4. **Implement caching**: Consider implementing caching mechanisms to reduce the load on Nginx and upstream services.
5. **Analyze and optimize request patterns**: Identify and optimize request patterns to reduce latency and improve overall performance.

Remember to investigate and address the root cause of the high latency to prevent similar issues in the future.