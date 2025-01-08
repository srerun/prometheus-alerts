---
title: ThanosQueryHttpRequestQueryErrorRateHigh
description: Troubleshooting for alert ThanosQueryHttpRequestQueryErrorRateHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosQueryHttpRequestQueryErrorRateHigh

Thanos Query {{$labels.job}} is failing to handle {{$value | humanize}}% of "query" requests.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-query.yml" "ThanosQueryHttpRequestQueryErrorRateHigh" %}}

{{% comment %}}

```yaml
alert: ThanosQueryHttpRequestQueryErrorRateHigh
expr: (sum by (job) (rate(http_requests_total{code=~"5..", job=~".*thanos-query.*", handler="query"}[5m]))/  sum by (job) (rate(http_requests_total{job=~".*thanos-query.*", handler="query"}[5m]))) * 100 > 5
for: 5m
labels:
    severity: critical
annotations:
    summary: Thanos Query Http Request Query Error Rate High (instance {{ $labels.instance }})
    description: |-
        Thanos Query {{$labels.job}} is failing to handle {{$value | humanize}}% of "query" requests.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-query/ThanosQueryHttpRequestQueryErrorRateHigh.md

```

{{% /comment %}}

</details>


## Meaning

The ThanosQueryHttpRequestQueryErrorRateHigh alert is triggered when the error rate for "query" requests in Thanos Query exceeds 5% over a 5-minute period. This indicates that Thanos Query is experiencing issues handling queries, which can lead to.data loss, delayed processing, or incomplete results.

## Impact

* Data inconsistencies or loss due to failed queries
* Delayed processing or incomplete results
* Potential impact on dependent services or applications
* Increased error rates can lead to decreased system performance and reliability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Query logs for error messages related to query handling.
2. Verify that the Thanos Query instance is running and configured correctly.
3. Check the system resources (CPU, memory, disk space) to ensure they are not causing the issue.
4. Verify that the query load is within expected limits and that there are no sudden spikes.
5. Check the networking and connectivity to ensure there are no issues.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Thanos Query configuration and adjust it if necessary to handle the query load.
2. Scale up the Thanos Query instance to handle the increased query load.
3. Implement query optimization techniques to reduce the load on the system.
4. Check for any software or configuration updates that may resolve the issue.
5. Consider implementing a load balancer or queueing mechanism to handle sudden spikes in query load.

Additional resources:

* Refer to the Thanos Query documentation for configuration and optimization guidelines.
* Consult with the development team to identify any application-level issues that may be contributing to the error rate.
* Review system monitoring and logging to identify any other potential issues.