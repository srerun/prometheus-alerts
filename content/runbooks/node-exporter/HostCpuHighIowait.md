---
title: HostCpuHighIowait
description: Troubleshooting for alert HostCpuHighIowait
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostCpuHighIowait

CPU iowait > 10%. A high iowait means that you are disk or network bound.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostCpuHighIowait" %}}

{{% comment %}}

```yaml
alert: HostCpuHighIowait
expr: (avg by (instance) (rate(node_cpu_seconds_total{mode="iowait"}[5m])) * 100 > 10) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host CPU high iowait (instance {{ $labels.instance }})
    description: |-
        CPU iowait > 10%. A high iowait means that you are disk or network bound.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostCpuHighIowait.md

```

{{% /comment %}}

</details>


Here is a runbook for the HostCpuHighIowait alert rule:

## Meaning

The HostCpuHighIowait alert is triggered when the average CPU I/O wait time for a host exceeds 10% over a 5-minute period. This indicates that the host is experiencing high I/O wait times, which can be caused by disk or network bottlenecks.

## Impact

High I/O wait times can lead to performance degradation, increased latency, and reduced overall system throughput. This can impact the responsiveness and reliability of applications and services running on the affected host.

## Diagnosis

To diagnose the root cause of high I/O wait times, follow these steps:

1. **Check disk usage and performance**: Investigate disk usage and performance metrics (e.g., disk IOPS, disk latency) to determine if there are any disk bottlenecks.
2. **Inspect network performance**: Examine network performance metrics (e.g., network throughput, packet loss) to determine if there are any network bottlenecks.
3. **Check system logs**: Review system logs for any errors or warnings related to disk or network issues.
4. **Verify resource utilization**: Check resource utilization (e.g., CPU, memory, disk space) to ensure that the host is not experiencing resource starvation.

## Mitigation

To mitigate high I/O wait times, follow these steps:

1. **Optimize disk configuration**: Ensure that disks are properly configured, and consider optimizing disk layout, disk scheduling, and disk caching.
2. **Improve network performance**: Ensure that network interfaces are properly configured, and consider upgrading network hardware or optimizing network routing.
3. **Implement caching or buffering**: Consider implementing caching or buffering mechanisms to reduce the load on disks and networks.
4. **Scale up or out**: If resource utilization is high, consider scaling up or out to distribute the load across multiple hosts.

Remember to monitor the host's performance and adjust your mitigation strategy as needed to ensure that I/O wait times return to normal levels.