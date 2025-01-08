---
title: ContainerHighCpuUtilization
description: Troubleshooting for alert ContainerHighCpuUtilization
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerHighCpuUtilization

Container CPU utilization is above 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerHighCpuUtilization" %}}

{{% comment %}}

```yaml
alert: ContainerHighCpuUtilization
expr: (sum(rate(container_cpu_usage_seconds_total{container!=""}[5m])) by (pod, container) / sum(container_spec_cpu_quota{container!=""}/container_spec_cpu_period{container!=""}) by (pod, container) * 100) > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: Container High CPU utilization (instance {{ $labels.instance }})
    description: |-
        Container CPU utilization is above 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerHighCpuUtilization.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ContainerHighCpuUtilization`:

## Meaning
This alert is triggered when a container's CPU utilization exceeds 80% over a 5-minute period. This can indicate that the container is experiencing high CPU usage, which may lead to performance issues, slower response times, or even crashes.

## Impact
High CPU utilization can have several consequences, including:

* Slower response times: High CPU usage can cause the container to slow down, leading to delayed responses to requests.
* Increased latency: High CPU utilization can increase the time it takes for the container to process requests, leading to increased latency.
* Resource starvation: High CPU utilization can starve other containers or processes of CPU resources, leading to performance issues or failures.
* Increased risk of crashes: Prolonged high CPU utilization can cause the container to crash or become unresponsive.

## Diagnosis
To diagnose the root cause of high CPU utilization, follow these steps:

1. Identify the affected container: Check the `pod` and `container` labels in the alert to determine which container is experiencing high CPU utilization.
2. Check container logs: Review the container logs to see if there are any error messages or indications of high CPU usage.
3. Check container resource usage: Use tools like `kubectl top` or `docker stats` to check the container's CPU usage, memory usage, and other resource utilization.
4. Analyze container configuration: Review the container configuration to ensure that it is properly configured and optimized for performance.
5. Check system resource usage: Verify that the underlying system has sufficient resources (e.g., CPU, memory) to support the container's workload.

## Mitigation
To mitigate high CPU utilization, follow these steps:

1. Optimize container configuration: Review and optimize the container configuration to ensure it is properly configured for performance.
2. Increase container resources: If necessary, increase the container's CPU or memory resources to ensure it has sufficient resources to handle the workload.
3. Implement rate limiting: Implement rate limiting or throttling to prevent excessive requests that can cause high CPU utilization.
4. Scale container instances: If the container is experiencing high CPU utilization due to high traffic, consider scaling the number of container instances to distribute the workload.
5. Investigate underlying system issues: If high CPU utilization persists, investigate underlying system issues, such as hardware failures or system configuration problems.

Remember to take a proactive approach to monitoring and addressing high CPU utilization to prevent performance issues and ensure the reliability of your system.