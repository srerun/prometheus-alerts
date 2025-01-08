---
title: IstioHighRequestLatency
description: Troubleshooting for alert IstioHighRequestLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IstioHighRequestLatency

Istio average requests execution is longer than 100ms.

<details>
  <summary>Alert Rule</summary>

{{% rule "istio/istio-internal.yml" "IstioHighRequestLatency" %}}

{{% comment %}}

```yaml
alert: IstioHighRequestLatency
expr: rate(istio_request_duration_milliseconds_sum{reporter="destination"}[1m]) / rate(istio_request_duration_milliseconds_count{reporter="destination"}[1m]) > 100
for: 1m
labels:
    severity: warning
annotations:
    summary: Istio high request latency (instance {{ $labels.instance }})
    description: |-
        Istio average requests execution is longer than 100ms.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/istio-internal/IstioHighRequestLatency.md

```

{{% /comment %}}

</details>


Here is a runbook for the IstioHighRequestLatency alert rule:

## Meaning

The IstioHighRequestLatency alert is triggered when the average request latency for an Istio instance exceeds 100ms. This alert is critical because high request latency can negatively impact the performance and responsiveness of the application.

## Impact

High request latency can lead to:

* Slow user experience
* Increased error rates
* Decreased system throughput
* Negative business impact due to poor performance

## Diagnosis

To diagnose the root cause of the high request latency, follow these steps:

1. Check the Istio dashboard for any signs of network congestion or packet loss.
2. Verify that the application is properly configured and that there are no misconfigured routes or policies.
3. Review the Istio request logs to identify any pattern of slow requests.
4. Check the resource utilization of the pods and nodes in the cluster to ensure that they are not overutilized.
5. Verify that the underlying infrastructure is healthy and not experiencing any issues.

## Mitigation

To mitigate the high request latency, follow these steps:

1. Identify and optimize any slow-performing services or deployments.
2. Verify that the Istio configuration is optimal for the current workload.
3. Consider increasing the resources allocated to the pods or nodes in the cluster.
4. Implement caching or other performance optimization techniques to reduce the load on the system.
5. Work with the development team to optimize the application code to reduce latency.

Remember to also review the Istio request logs and dashboards to ensure that the mitigation steps have resolved the issue and that the average request latency has returned to a normal level.