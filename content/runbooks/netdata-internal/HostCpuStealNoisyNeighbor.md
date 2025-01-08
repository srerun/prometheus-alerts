---
title: HostCpuStealNoisyNeighbor
description: Troubleshooting for alert HostCpuStealNoisyNeighbor
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostCpuStealNoisyNeighbor

CPU steal is > 10%. A noisy neighbor is killing VM performances or a spot instance may be out of credit.

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "HostCpuStealNoisyNeighbor" %}}

{{% comment %}}

```yaml
alert: HostCpuStealNoisyNeighbor
expr: rate(netdata_cpu_cpu_percentage_average{dimension="steal"}[1m]) > 10
for: 5m
labels:
    severity: warning
annotations:
    summary: Host CPU steal noisy neighbor (instance {{ $labels.instance }})
    description: |-
        CPU steal is > 10%. A noisy neighbor is killing VM performances or a spot instance may be out of credit.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/HostCpuStealNoisyNeighbor.md

```

{{% /comment %}}

</details>


## Meaning

The HostCpuStealNoisyNeighbor alert indicates that the average CPU steal percentage on a host has exceeded 10% over the past 1 minute, and this condition has persisted for at least 5 minutes. CPU steal occurs when a virtual machine (VM) is waiting for the hypervisor to allocate CPU resources, which can lead to performance issues.

## Impact

This alert can have a significant impact on the performance of applications running on the affected host. Noisy neighbors, such as other VMs on the same host, can consume excessive CPU resources, leading to:

* Increased latency
* Decreased throughput
* Unresponsive applications
* Potential crashes or failures of critical services

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

1. Check the Netdata dashboard for the affected host to identify the current CPU steal percentage and trend.
2. Investigate the instance labels `$labels` to determine the specific VM or instance affected.
3. Review system logs and monitoring data to identify any other performance issues or anomalies on the host.
4. Check the hypervisor logs to identify if there are any issues with resource allocation or scheduling.

## Mitigation

To mitigate the impact of this issue, follow these steps:

1. Isolate the noisy neighbor: Identify the specific VM or instance causing the CPU steal and consider migrating it to a different host or adjusting its resource allocation.
2. Adjust instance sizing: Review the instance sizes and resource allocations to ensure they are adequate for the workload.
3. Optimize hypervisor settings: Check the hypervisor settings to ensure they are optimized for the current workload and resource utilization.
4. Monitor and adjust: Continuously monitor the CPU steal percentage and adjust the instance sizing, resource allocation, and hypervisor settings as needed to maintain optimal performance.

Remember to consult the [Netdata documentation](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/HostCpuStealNoisyNeighbor.md) for more detailed guidance on diagnosing and mitigating this issue.