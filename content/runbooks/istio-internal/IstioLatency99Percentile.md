---
title: IstioLatency99Percentile
description: Troubleshooting for alert IstioLatency99Percentile
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioLatency99Percentile

Istio 1% slowest requests are longer than 1000ms.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioLatency99Percentile" %}}

{{% comment %}}

```yaml
alert: IstioLatency99Percentile
expr: histogram_quantile(0.99, sum(rate(istio_request_duration_milliseconds_bucket[1m])) by (destination_canonical_service, destination_workload_namespace, source_canonical_service, source_workload_namespace, le)) > 1000
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio latency 99 percentile (instance {{ $labels.instance }})
    description: |-
        Istio 1% slowest requests are longer than 1000ms.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioLatency99Percentile.md

```

{{% /comment %}}

</details>


Here is a runbook for the IstioLatency99Percentile alert rule:

## Meaning

The IstioLatency99Percentile alert is triggered when the 99th percentile of Istio request duration exceeds 1000ms for a given service and namespace. This indicates that 1% of requests are experiencing high latency, which can impact the performance and responsiveness of the application.

## Impact

* Slow request processing can lead to a poor user experience, decreased system throughput, and increased error rates.
* High latency can also cause cascading failures, as dependent services may timeout or become unavailable.
* Prolonged latency issues can result in revenue loss, customer dissatisfaction, and damage to the organization's reputation.

## Diagnosis

To diagnose the root cause of high latency, follow these steps:

1. **Check request traffic patterns**: Analyze the request rate and distribution to identify any sudden changes or spikes.
2. **Investigate service dependencies**: Verify that dependent services are operating within expected latency bounds.
3. **Review Istio configuration**: Ensure that Istio is properly configured, and that service mesh metrics are being collected correctly.
4. **Examine pod and container logs**: Review logs for errors, warnings, or other indicators of issues that may be contributing to high latency.
5. **Check for resource constraints**: Verify that pods have sufficient resources (CPU, memory, etc.) to handle incoming requests.

## Mitigation

To mitigate high latency, take the following steps:

1. **Scale out services**: Temporarily increase the number of replicas for the affected service to handle the increased request volume.
2. **Optimize service configuration**: Review and optimize service configuration, such as timeouts, retries, and circuit breakers.
3. **Improve resource allocation**: Ensure that pods have sufficient resources to handle incoming requests.
4. **Apply caching or content compression**: Implement caching or content compression to reduce the load on services and improve response times.
5. **Investigate root cause**: Identify and address the underlying cause of high latency, which may require fixing application code, database queries, or other system components.