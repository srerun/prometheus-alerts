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

{{% rule "host-and-hardware/node-exporter.yml" "HostCpuStealNoisyNeighbor" %}}

{{% comment %}}

```yaml
alert: HostCpuStealNoisyNeighbor
expr: (avg by(instance) (rate(node_cpu_seconds_total{mode="steal"}[5m])) * 100 > 10) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host CPU steal noisy neighbor (instance {{ $labels.instance }})
    description: |-
        CPU steal is > 10%. A noisy neighbor is killing VM performances or a spot instance may be out of credit.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostCpuStealNoisyNeighbor.md

```

{{% /comment %}}

</details>


## Meaning

The HostCpuStealNoisyNeighbor alert is triggered when the average CPU steal rate across all instances exceeds 10% over a 5-minute period. This indicates that one or more instances on a host are experiencing high CPU steal, which can negatively impact their performance. CPU steal occurs when a virtual machine (VM) or container is waiting for the hypervisor to allocate CPU resources, resulting in stolen CPU cycles.

## Impact

The impact of high CPU steal on instances can be significant, leading to:

* Increased latency and response times
* Decreased throughput and performance
* Increased risk of timeouts and errors
* Potential for instance crashes or failures
* In extreme cases, entire hosts or clusters may become unresponsive or crash

## Diagnosis

To diagnose the root cause of the HostCpuStealNoisyNeighbor alert, follow these steps:

1. Identify the affected instances: Check the `instance` label in the alert to determine which instances are experiencing high CPU steal.
2. Investigate the host: Use node_uname_info to identify the host on which the instances are running.
3. Check for noisy neighbors: Inspect the host's resource utilization to identify if there are any noisy neighbors (e.g., instances consuming excessive resources) that may be contributing to the high CPU steal.
4. Verify VM or spot instance configuration: Check the configuration of the affected instances to ensure they are not running in a spot instance or have incorrect VM settings that may be causing the high CPU steal.
5. Review system logs: Analyze system logs to identify any errors or warnings related to CPU steal or resource contention.

## Mitigation

To mitigate the HostCpuStealNoisyNeighbor alert, follow these steps:

1. Identify and terminate noisy neighbors: If a noisy neighbor is identified, terminate the instance to prevent further resource contention.
2. Adjust instance configuration: Check and adjust the instance configuration to ensure it is not running in a spot instance or has incorrect VM settings that may be causing the high CPU steal.
3. Implement resource constraints: Apply resource constraints (e.g., CPU limits) to prevent instances from consuming excessive resources.
4. Consider horizontal scaling: If the host is consistently experiencing high CPU steal, consider scaling out the host or adding more resources to improve performance.
5. Monitor and re-evaluate: Continuously monitor the affected instances and re-evaluate the alert threshold to ensure it is set appropriately for the environment.