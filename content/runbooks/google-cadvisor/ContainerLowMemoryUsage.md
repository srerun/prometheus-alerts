---
title: ContainerLowMemoryUsage
description: Troubleshooting for alert ContainerLowMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerLowMemoryUsage

Container Memory usage is under 20% for 1 week. Consider reducing the allocated memory.

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerLowMemoryUsage" %}}

{{% comment %}}

```yaml
alert: ContainerLowMemoryUsage
expr: (sum(container_memory_working_set_bytes{name!=""}) BY (instance, name) / sum(container_spec_memory_limit_bytes > 0) BY (instance, name) * 100) < 20
for: 7d
labels:
    severity: info
annotations:
    summary: Container Low Memory usage (instance {{ $labels.instance }})
    description: |-
        Container Memory usage is under 20% for 1 week. Consider reducing the allocated memory.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerLowMemoryUsage.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "ContainerLowMemoryUsage":

## Meaning

The "ContainerLowMemoryUsage" alert is triggered when a container's memory usage falls below 20% of its allocated memory limit for an extended period of time (7 days). This alert is informational in nature and is intended to notify operators of potential opportunities to optimize memory allocation.

## Impact

This alert has a relatively low impact on the system, as it is primarily an optimization opportunity. However, if left unaddressed, underutilized memory can lead to inefficient resource allocation and potentially impact the overall performance of the system.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected container: Check the alert labels to determine which container is experiencing low memory usage.
2. Verify the memory usage: Use tools like `kubectl top` or `docker stats` to verify the container's memory usage and confirm that it is indeed below 20% of its allocated limit.
3. Review container configuration: Check the container's configuration files (e.g. Dockerfile, Kubernetes deployment YAML) to understand the memory allocation settings.

## Mitigation

To mitigate this issue, follow these steps:

1. Review and adjust memory allocation: Evaluate whether the allocated memory limit is reasonable for the container's workload. Consider reducing the memory limit to a more suitable value.
2. Monitor container performance: Continue to monitor the container's performance and adjust the memory allocation as needed to ensure optimal resource utilization.
3. Implement container resource optimization: Consider implementing container resource optimization techniques, such as vertical pod autoscaling, to ensure efficient use of resources.

Remember to refer to the [alert documentation](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerLowMemoryUsage.md) for more information and guidance on handling this alert.