---
title: ThanosQueryHttpRequestQueryRangeErrorRateHigh
description: Troubleshooting for alert ThanosQueryHttpRequestQueryRangeErrorRateHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryHttpRequestQueryRangeErrorRateHigh

Thanos Query {{$labels.job}} is failing to handle {{$value | humanize}}% of "query_range" requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryHttpRequestQueryRangeErrorRateHigh" %}}

{{% comment %}}

```yaml
alert: ThanosQueryHttpRequestQueryRangeErrorRateHigh
expr: (sum by (job) (rate(http_requests_total{code=~"5..", job=~".*thanos-query.*", handler="query_range"}[5m]))/  sum by (job) (rate(http_requests_total{job=~".*thanos-query.*", handler="query_range"}[5m]))) * 100 > 5
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Query Http Request Query Range Error Rate High (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} is failing to handle {{$value | humanize}}% of "query_range" requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryHttpRequestQueryRangeErrorRateHigh.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ThanosQueryHttpRequestQueryRangeErrorRateHigh` using level 2 headers:

## Meaning

This alert is triggered when the error rate for "query_range" requests to Thanos Query exceeds 5% over a 5-minute period. The error rate is calculated as the ratio of failed requests to total requests, aggregated by job.

## Impact

A high error rate for "query_range" requests can indicate a problem with Thanos Query, which may lead to:

* Incomplete or inaccurate query results
* Increased latency or timeouts for query requests
* Degraded performance or unavailability of dependent systems

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Query logs for errors related to "query_range" requests
2. Verify that the Thanos Query instance is running with the correct configuration and resources (e.g., CPU, memory)
3. Investigate any recent changes to the Thanos Query configuration, infrastructure, or dependent systems
4. Check the system metrics (e.g., CPU usage, memory usage, disk usage) to identify any resource bottlenecks
5. Verify that the network connectivity and firewall rules are allowing traffic to/from the Thanos Query instance

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Thanos Query instance to reset the query_range request handler
2. Investigate and resolve any underlying issues identified during diagnosis (e.g., configuration errors, resource bottlenecks)
3. Implement retry mechanisms or circuits breakers to handle temporary failures in the query_range request handler
4. Consider scaling up or out the Thanos Query instance to handle increased load or traffic
5. Monitor the error rate and system metrics closely to ensure the issue is resolved and does not recur.