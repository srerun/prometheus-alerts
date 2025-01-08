---
title: ThanosQueryOverload
description: Troubleshooting for alert ThanosQueryOverload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryOverload

Thanos Query {{$labels.job}} has been overloaded for more than 15 minutes. This may be a symptom of excessive simultanous complex requests, low performance of the Prometheus API, or failures within these components. Assess the health of the Thanos query instances, the connnected Prometheus instances, look for potential senders of these requests and then contact support.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryOverload" %}}

{{% comment %}}

```yaml
alert: ThanosQueryOverload
expr: (max_over_time(thanos_query_concurrent_gate_queries_max[5m]) - avg_over_time(thanos_query_concurrent_gate_queries_in_flight[5m]) < 1)
for: 15m
labels:
    severity: warning
annotations:
    summary: Thanos Query Overload (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} has been overloaded for more than 15 minutes. This may be a symptom of excessive simultanous complex requests, low performance of the Prometheus API, or failures within these components. Assess the health of the Thanos query instances, the connnected Prometheus instances, look for potential senders of these requests and then contact support.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryOverload.md

```

{{% /comment %}}

</details>


## Meaning

The ThanosQueryOverload alert is triggered when the number of concurrent Thanos query gate requests exceeds the average number of in-flight requests by a significant margin (less than 1) for more than 15 minutes. This indicates that the Thanos query instances are overwhelmed, which can lead to performance issues, slow query responses, or even failures.

## Impact

The impact of this alert can be significant, as it may lead to:

* Slow or failed queries, affecting the reliability and accuracy of monitoring data
* Increased latency and decreased performance of the Prometheus API
* Potential cascading failures in connected components
* Delays in identifying and resolving underlying issues due to slow query responses

## Diagnosis

To diagnose the root cause of the ThanosQueryOverload alert, follow these steps:

1. Investigate the health of the Thanos query instances:
	* Check the instance's resource utilization, such as CPU, memory, and disk usage
	* Verify that the instance is not experiencing any failures or restarts
2. Examine the connected Prometheus instances:
	* Check the Prometheus instances' resource utilization and health
	* Verify that the Prometheus instances are not experiencing any issues with scraping, storing, or querying data
3. Identify potential senders of excessive requests:
	* Analyze the query patterns and identify any unusual or excessive query requests
	* Check the dashboards and Alertmanager configurations for any misconfigured or abusive queries
4. Review the Thanos query logs for any errors or warnings:
	* Check the logs for any errors related to query execution, connection issues, or timeouts

## Mitigation

To mitigate the ThanosQueryOverload alert, follow these steps:

1. Scale the Thanos query instances:
	* Increase the instance count or resource allocation to handle the increased query load
	* Verify that the instances are properly configured and healthy
2. Optimize Prometheus instances:
	* Ensure that the Prometheus instances are properly configured and optimized for performance
	* Consider implementing query caching or other performance optimizations
3. Identify and address excessive request senders:
	* Work with the teams responsible for the dashboards and Alertmanager configurations to identify and fix any misconfigured or abusive queries
	* Implement rate limiting or other measures to prevent excessive requests
4. Continue monitoring and troubleshooting:
	* Closely monitor the Thanos query instances and Prometheus instances for any signs of overload or failure
	* Continuously troubleshoot and address any underlying issues to prevent future occurrences of this alert