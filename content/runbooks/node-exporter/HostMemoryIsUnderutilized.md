---
title: HostMemoryIsUnderutilized
description: Troubleshooting for alert HostMemoryIsUnderutilized
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostMemoryIsUnderutilized

Node memory is < 20% for 1 week. Consider reducing memory space. (instance {{ $labels.instance }})

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostMemoryIsUnderutilized" %}}

{{% comment %}}

```yaml
alert: HostMemoryIsUnderutilized
expr: (100 - (avg_over_time(node_memory_MemAvailable_bytes[30m]) / node_memory_MemTotal_bytes * 100) < 20) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 1w
labels:
    severity: info
annotations:
    summary: Host Memory is underutilized (instance {{ $labels.instance }})
    description: |-
        Node memory is < 20% for 1 week. Consider reducing memory space. (instance {{ $labels.instance }})
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostMemoryIsUnderutilized.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule "HostMemoryIsUnderutilized":

## Meaning

The "HostMemoryIsUnderutilized" alert is triggered when a host's memory utilization has been below 20% for a period of one week. This alert is informational, indicating that the host's memory resources are not being fully utilized.

## Impact

The impact of underutilized memory is primarily related to resource efficiency and cost optimization. If a host's memory is consistently underutilized, it may be possible to:

* Reduce the host's memory allocation, resulting in cost savings
* Repurpose or reallocate the underutilized memory to other resources or applications that require it
* Improve overall resource utilization and efficiency

## Diagnosis

To diagnose the root cause of underutilized memory, follow these steps:

1. Review the host's resource utilization patterns over the past week to identify any trends or anomalies.
2. Check the host's running processes and applications to determine if any are memory-intensive.
3. Verify that the host's memory configuration is optimal for its workload.
4. Investigate whether any recent changes have been made to the host's configuration or workload that may be contributing to underutilization.

## Mitigation

To mitigate underutilized memory, consider the following steps:

1. Evaluate the host's workload and determine if it can be optimized to utilize more memory.
2. Consider reducing the host's memory allocation to a level that is more in line with its actual utilization.
3. Repurpose or reallocate the underutilized memory to other resources or applications that require it.
4. Monitor the host's memory utilization over time to ensure that the issue has been resolved and to identify any potential future issues.