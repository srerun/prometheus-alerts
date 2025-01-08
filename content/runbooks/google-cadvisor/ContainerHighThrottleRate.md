---
title: ContainerHighThrottleRate
description: Troubleshooting for alert ContainerHighThrottleRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerHighThrottleRate

Container is being throttled

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerHighThrottleRate" %}}

{{% comment %}}

```yaml
alert: ContainerHighThrottleRate
expr: sum(increase(container_cpu_cfs_throttled_periods_total{container!=""}[5m])) by (container, pod, namespace) / sum(increase(container_cpu_cfs_periods_total[5m])) by (container, pod, namespace) > ( 25 / 100 )
for: 5m
labels:
    severity: warning
annotations:
    summary: Container high throttle rate (instance {{ $labels.instance }})
    description: |-
        Container is being throttled
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerHighThrottleRate.md

```

{{% /comment %}}

</details>


## Meaning

The ContainerHighThrottleRate alert is triggered when the CPU throttle rate of a container exceeds 25% over a 5-minute period. This means that the container is being throttled, resulting in reduced performance and potential impact on the overall system.

## Impact

The impact of a high throttle rate can be significant, leading to:

* Reduced container performance, resulting in slower response times and decreased throughput
* Increased latency and errors, potentially affecting user experience and business operations
* Resource waste, as the container is not utilizing available CPU resources efficiently
* Potential cascade effects on dependent services and applications

## Diagnosis

To diagnose the root cause of the high throttle rate, follow these steps:

1. Identify the affected container, pod, and namespace using the alert labels
2. Check the container's CPU usage and throttling period metrics using tools like `top` or `docker stats`
3. Investigate potential causes, such as:
	* Insufficient CPU resources allocated to the container
	* Resource-intensive tasks or workloads running in the container
	* Poor resource utilization or inefficient application design
4. Review system logs and container logs to identify any relevant errors or warnings
5. Consult with the application team and developers to understand the expected behavior and resource requirements of the container

## Mitigation

To mitigate the high throttle rate, follow these steps:

1. Investigate and address any underlying resource issues, such as:
	* Increasing CPU resources allocated to the container
	* Optimizing resource utilization through container tuning or rightsizing
	* Implementing load balancing or horizontal scaling to distribute workload
2. Optimize application performance and resource efficiency, such as:
	* Implementing caching or content delivery networks to reduce load
	* Optimizing database queries and transactions
	* Implementing efficient algorithms and data structures
3. Consider implementing throttling limits or rate limiting to prevent excessive resource usage
4. Monitor and analyze container performance and resource utilization to ensure the issue is resolved and to prevent future occurrences
5. Develop and implement a long-term plan to ensure sustainable resource allocation and efficient container performance.