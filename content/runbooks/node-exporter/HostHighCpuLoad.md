---
title: HostHighCpuLoad
description: Troubleshooting for alert HostHighCpuLoad
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostHighCpuLoad

CPU load is > 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostHighCpuLoad" %}}

{{% comment %}}

```yaml
alert: HostHighCpuLoad
expr: (sum by (instance) (avg by (mode, instance) (rate(node_cpu_seconds_total{mode!="idle"}[2m]))) > 0.8) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 10m
labels:
    severity: warning
annotations:
    summary: Host high CPU load (instance {{ $labels.instance }})
    description: |-
        CPU load is > 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostHighCpuLoad.md

```

{{% /comment %}}

</details>


## Meaning

The HostHighCpuLoad alert is triggered when a host's CPU load exceeds 80% for more than 10 minutes. This alert indicates that the host is experiencing high CPU utilization, which can lead to performance issues, slow response times, and even crashes.

## Impact

High CPU load can have a significant impact on the overall performance and reliability of the system. Some possible consequences include:

* Slow response times for users
* Increased latency for critical applications
* Decreased system throughput
* Increased risk of system crashes or downtime
* Potential impact on other dependent systems or services

## Diagnosis

To diagnose the root cause of the high CPU load, follow these steps:

1. Check the Prometheus graph for the `node_cpu_seconds_total` metric to identify the trend and pattern of CPU usage.
2. Investigate the top CPU-consuming processes using tools like `top`, `htop`, or `mpstat`.
3. Review system logs for any errors or warnings that may indicate the cause of the high CPU load.
4. Check for any recent changes or updates to the system configuration, software, or applications that may be contributing to the high CPU load.
5. Verify that the system is properly configured and optimized for the workload.

## Mitigation

To mitigate the high CPU load, follow these steps:

1. Identify and terminate any unnecessary or resource-intensive processes.
2. Optimize system configuration and settings to improve performance.
3. Consider upgrading or adding more CPU resources to the system.
4. Implement tuning and optimization techniques, such as adjusting system parameters or configuring caching.
5. Consider load balancing or distributing workloads across multiple systems to reduce the load on the affected host.

Remember to follow best practices and test any changes before implementing them in production. Additionally, consider creating a permanent fix to prevent similar issues from happening in the future.