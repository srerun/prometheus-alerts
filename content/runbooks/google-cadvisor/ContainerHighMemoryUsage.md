---
title: ContainerHighMemoryUsage
description: Troubleshooting for alert ContainerHighMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerHighMemoryUsage

Container Memory usage is above 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerHighMemoryUsage" %}}

{{% comment %}}

```yaml
alert: ContainerHighMemoryUsage
expr: (sum(container_memory_working_set_bytes{name!=""}) BY (instance, name) / sum(container_spec_memory_limit_bytes > 0) BY (instance, name) * 100) > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: Container High Memory usage (instance {{ $labels.instance }})
    description: |-
        Container Memory usage is above 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerHighMemoryUsage.md

```

{{% /comment %}}

</details>


## Meaning

The ContainerHighMemoryUsage alert is triggered when the average memory usage of a container exceeds 80% of its memory limit. This alert is designed to detect containers that are consuming excessive memory, which can lead to performance issues, slow response times, and even crashes.

## Impact

* Containers consuming high memory may lead to:
	+ Slow response times and poor application performance
	+ Increased latency and timeouts
	+ Decreased system stability and increased crash likelihood
	+ Overuse of resources, leading to increased costs
* If left unattended, high memory usage can cause:
	+ Cascading failures of dependent services
	+ Increased difficulty in troubleshooting and diagnosing issues
	+ Decreased overall system reliability and availability

## Diagnosis

1. **Identify the affected container**: Check the alert annotations for the instance and name of the affected container.
2. **Check container memory usage**: Use tools like `docker stats` or `kubectl top` to verify the current memory usage of the container.
3. **Review container configuration**: Examine the container's resource allocation, such as the memory limit and request, to ensure they are properly configured.
4. **Investigate recent changes**: Look for recent changes to the container's configuration, code, or dependencies that may be contributing to the high memory usage.
5. **Check for memory leaks**: Use tools like `pmap` or `heapdump` to identify potential memory leaks in the application.

## Mitigation

1. **Investigate and address root cause**: Identify and resolve the underlying issue causing high memory usage, such as a memory leak or inefficient coding practices.
2. **Adjust container resource allocation**: Increase the memory limit or request for the container, if necessary, to ensure it has sufficient resources to operate efficiently.
3. **Implement memory efficient practices**: Optimize the application's memory usage by implementing efficient data structures, caching, and garbage collection.
4. **Monitor container memory usage**: Regularly monitor the container's memory usage to detect potential issues before they cause significant problems.
5. **Consider horizontal pod autoscaling**: If the container is part of a Kubernetes deployment, consider implementing horizontal pod autoscaling to dynamically adjust the number of replicas based on resource utilization.