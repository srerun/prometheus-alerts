---
title: ContainerLowCpuUtilization
description: Troubleshooting for alert ContainerLowCpuUtilization
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerLowCpuUtilization

Container CPU utilization is under 20% for 1 week. Consider reducing the allocated CPU.

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerLowCpuUtilization" %}}

{{% comment %}}

```yaml
alert: ContainerLowCpuUtilization
expr: (sum(rate(container_cpu_usage_seconds_total{container!=""}[5m])) by (pod, container) / sum(container_spec_cpu_quota{container!=""}/container_spec_cpu_period{container!=""}) by (pod, container) * 100) < 20
for: 7d
labels:
    severity: info
annotations:
    summary: Container Low CPU utilization (instance {{ $labels.instance }})
    description: |-
        Container CPU utilization is under 20% for 1 week. Consider reducing the allocated CPU.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerLowCpuUtilization.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule `ContainerLowCpuUtilization`:

## Meaning

The `ContainerLowCpuUtilization` alert is triggered when a container's CPU utilization falls below 20% for a period of 7 days. This alert is classified as informational and is not critical, but it may indicate inefficient resource allocation.

## Impact

Low CPU utilization can lead to wasted resources and inefficient use of cluster capacity. If left unchecked, it can result in:

* Underutilized nodes and clusters
* Inefficient resource allocation
* Higher costs due to unused resources
* Potential performance issues due to overprovisioning

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected container and pod using the `instance` label provided in the alert.
2. Check the container's CPU usage over the past 7 days using a tool like Prometheus or Grafana.
3. Verify that the container is not experiencing any performance issues or errors that could be contributing to the low CPU utilization.
4. Review the container's resource allocation and configuration to ensure it is properly sized for its workload.

## Mitigation

To mitigate the issue, consider the following steps:

1. Reduce the allocated CPU resources for the container to match its actual usage.
2. Right-size the container's resource allocation based on its workload and performance requirements.
3. Consider consolidating underutilized containers onto fewer nodes to optimize resource usage.
4. Implement automated resource scaling and optimization tools to detect and adjust resource allocation in real-time.

Note: This runbook provides general guidance and may need to be tailored to your specific use case and infrastructure.