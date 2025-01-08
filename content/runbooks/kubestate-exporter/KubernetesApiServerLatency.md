---
title: KubernetesApiServerLatency
description: Troubleshooting for alert KubernetesApiServerLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# KubernetesApiServerLatency

Kubernetes API server has a 99th percentile latency of {{ $value }} seconds for {{ $labels.verb }} {{ $labels.resource }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "kubernetes/kubestate-exporter.yml" "KubernetesApiServerLatency" %}}

{{% comment %}}

```yaml
alert: KubernetesApiServerLatency
expr: histogram_quantile(0.99, sum(rate(apiserver_request_duration_seconds_bucket{verb!~"(?:CONNECT|WATCHLIST|WATCH|PROXY)"} [10m])) WITHOUT (subresource)) > 1
for: 2m
labels:
    severity: warning
annotations:
    summary: Kubernetes API server latency (instance {{ $labels.instance }})
    description: |-
        Kubernetes API server has a 99th percentile latency of {{ $value }} seconds for {{ $labels.verb }} {{ $labels.resource }}.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/kubestate-exporter/KubernetesApiServerLatency.md

```

{{% /comment %}}

</details>


Here is a runbook for the KubernetesApiServerLatency alert:

## Meaning

The KubernetesApiServerLatency alert is triggered when the 99th percentile latency of the Kubernetes API server exceeds 1 second over a 10-minute period. This means that a significant portion of API requests are taking longer than expected to process, which can impact the performance and responsiveness of the Kubernetes cluster.

## Impact

A high latency in the Kubernetes API server can have several impacts on the cluster and its users:

* Increased response times for kubectl commands and other API requests
* Delays in pod scheduling and deployment
* Impact on the overall performance and responsiveness of the cluster
* Frustration and decreased productivity for cluster users

## Diagnosis

To diagnose the root cause of the high latency, follow these steps:

1. Check the API server logs for errors or unusual patterns
2. Verify that the API server is not overloaded or under-provisioned
3. Check the latency of specific API requests using `kubectl get --raw /apis/<resource>`
4. Investigate any recent changes or deployments that may be causing the latency
5. Use tools like `kubectl top` or `kubectl describe` to monitor API server performance and resource utilization

## Mitigation

To mitigate the high latency, follow these steps:

1. **Investigate and resolve any underlying issues**: Address any errors, overload, or under-provisioning of the API server
2. **Optimize API server configuration**: Adjust API server settings to improve performance, such as increasing the number of workers or adjusting the request timeout
3. **Implement caching or latency-reducing mechanisms**: Consider implementing caching mechanisms, such as Cluster Autoscaler, to reduce the load on the API server
4. **Monitor and alert on API server performance**: Set up additional monitoring and alerting to quickly detect and respond to API server performance issues
5. **Consider scaling or upgrading the API server**: If the latency persists, consider scaling or upgrading the API server to improve performance and responsiveness.