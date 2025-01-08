---
title: HadoopResourceManagerMemoryHigh
description: Troubleshooting for alert HadoopResourceManagerMemoryHigh
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HadoopResourceManagerMemoryHigh

The Hadoop ResourceManager is approaching its memory limit.

<details>
  <summary>Alert Rule</summary>

{{% rule "hadoop/jmx_exporter.yml" "HadoopResourceManagerMemoryHigh" %}}

{{% comment %}}

```yaml
alert: HadoopResourceManagerMemoryHigh
expr: hadoop_resourcemanager_memory_bytes / hadoop_resourcemanager_memory_max_bytes > 0.8
for: 15m
labels:
    severity: warning
annotations:
    summary: Hadoop Resource Manager Memory High (instance {{ $labels.instance }})
    description: |-
        The Hadoop ResourceManager is approaching its memory limit.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/jmx_exporter/HadoopResourceManagerMemoryHigh.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `HadoopResourceManagerMemoryHigh`:

## Meaning

The `HadoopResourceManagerMemoryHigh` alert is triggered when the memory usage of the Hadoop ResourceManager exceeds 80% of its maximum allowed memory. This indicates that the ResourceManager is approaching its memory limit, which can lead to performance degradation, slowdowns, and potentially even crashes.

## Impact

The impact of this alert is significant, as a ResourceManager with high memory usage can lead to:

* Slower job scheduling and execution
* Increased latency for users submitting jobs
* Potential crashes or restarts of the ResourceManager, leading to cluster instability
* In extreme cases, complete cluster unavailability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ResourceManager's log files for any error messages or warnings related to memory usage.
2. Verify that the ResourceManager's memory settings are configured correctly and are not too low.
3. Investigate any recent changes to the cluster configuration or application workload that may be contributing to the increased memory usage.
4. Check the Hadoop cluster's overall health and performance using metrics such as CPU usage, disk usage, and job queue length.
5. Consider increasing the ResourceManager's memory allocation or adding more nodes to the cluster to alleviate pressure on the ResourceManager.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the ResourceManager's memory allocation to give it more headroom to operate.
2. Optimize job scheduling and resource allocation to reduce the load on the ResourceManager.
3. Consider implementing memory-saving techniques such as compressing data or using more efficient data structures.
4. Monitor the ResourceManager's memory usage closely to catch any future issues early.
5. Consider implementing automated alerting and remediation actions to respond to high memory usage more quickly in the future.