---
title: HostCpuIsUnderutilized
description: Troubleshooting for alert HostCpuIsUnderutilized
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostCpuIsUnderutilized

CPU load is < 20% for 1 week. Consider reducing the number of CPUs.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostCpuIsUnderutilized" %}}

{{% comment %}}

```yaml
alert: HostCpuIsUnderutilized
expr: (100 - (rate(node_cpu_seconds_total{mode="idle"}[30m]) * 100) < 20) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 1w
labels:
    severity: info
annotations:
    summary: Host CPU is underutilized (instance {{ $labels.instance }})
    description: |-
        CPU load is < 20% for 1 week. Consider reducing the number of CPUs.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostCpuIsUnderutilized.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the `HostCpuIsUnderutilized` alert:

## Meaning

The `HostCpuIsUnderutilized` alert is triggered when the CPU load of a host is less than 20% for a sustained period of 1 week. This indicates that the host is not utilizing its CPU resources efficiently, and there may be opportunities to optimize resource allocation.

## Impact

The impact of underutilized CPU resources can be significant. It can lead to:

* Inefficient use of resources, resulting in unnecessary costs
* Inability to handle increased workloads, potentially leading to performance issues
* Difficulty in scaling applications and services, as underutilized resources may not be able to handle increased demand

## Diagnosis

To diagnose the root cause of the underutilized CPU, follow these steps:

1. Check the node's CPU utilization graph in Prometheus to identify the pattern of underutilization.
2. Review the node's workload and application logs to identify any potential bottlenecks or inefficiencies.
3. Verify that the node's configuration and resource allocation are optimal for the current workload.
4. Check for any resource-intensive processes or applications that may be consuming CPU resources inefficiently.

## Mitigation

To mitigate the issue of underutilized CPU, follow these steps:

1. Review the node's configuration and resource allocation to identify opportunities for optimization.
2. Consider reducing the number of CPUs allocated to the node, if possible, to align with the current workload demands.
3. Optimize resource-intensive processes or applications to improve their CPU efficiency.
4. Consider right-sizing the node's resources to match the current workload demands.
5. Monitor the node's CPU utilization regularly to ensure that the mitigation steps have been effective.