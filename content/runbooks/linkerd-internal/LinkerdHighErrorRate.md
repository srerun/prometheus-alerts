---
title: LinkerdHighErrorRate
description: Troubleshooting for alert LinkerdHighErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LinkerdHighErrorRate

Linkerd error rate for {{ $labels.deployment | $labels.statefulset | $labels.daemonset }} is over 10%

<details>
  <summary>Alert Rule</summary>

{{% rule "linkerd/linkerd-internal.yml" "LinkerdHighErrorRate" %}}

{{% comment %}}

```yaml
alert: LinkerdHighErrorRate
expr: sum(rate(request_errors_total[1m])) by (deployment, statefulset, daemonset) / sum(rate(request_total[1m])) by (deployment, statefulset, daemonset) * 100 > 10
for: 1m
labels:
    severity: warning
annotations:
    summary: Linkerd high error rate (instance {{ $labels.instance }})
    description: |-
        Linkerd error rate for {{ $labels.deployment | $labels.statefulset | $labels.daemonset }} is over 10%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/linkerd-internal/LinkerdHighErrorRate.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "LinkerdHighErrorRate":

## Meaning

The LinkerdHighErrorRate alert is triggered when the error rate of requests processed by Linkerd exceeds 10% over a 1-minute window. This indicates that there is an issue with the requests being processed by Linkerd, which may be caused by a variety of factors such as misconfigured services, network issues, or high latency.

## Impact

The impact of a high error rate in Linkerd can be significant, as it may lead to:

* Service disruptions or downtime
* Increased latency and response times
* decreased user satisfaction and revenue loss
* Increased load on error-handling mechanisms, leading to further performance degradation

## Diagnosis

To diagnose the root cause of the high error rate, follow these steps:

1. Check the Linkerd logs for errors and exceptions to identify the specific services or requests causing the errors.
2. Verify that the services and deployments involved are properly configured and running correctly.
3. Check the network and infrastructure for any issues or congestion that may be contributing to the errors.
4. Review the request patterns and distributions to identify any anomalies or trends that may be causing the errors.
5. Consult with the development and operations teams to gather more information about recent changes or deployments that may be related to the issue.

## Mitigation

To mitigate the high error rate, take the following steps:

1.Restart or redeploy the affected services to ensure they are running correctly.
2. Investigate and address any underlying issues with the services, deployments, or infrastructure.
3. Implement retries or circuit breakers to handle temporary errors and prevent cascading failures.
4. Consider implementing additional logging or monitoring to improve visibility into the request processing pipeline.
5. Perform a post-mortem analysis to identify the root cause of the issue and implement preventive measures to avoid similar issues in the future.