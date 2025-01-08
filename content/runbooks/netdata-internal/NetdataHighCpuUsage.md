---
title: NetdataHighCpuUsage
description: Troubleshooting for alert NetdataHighCpuUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataHighCpuUsage

Netdata high CPU usage (> 80%)

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataHighCpuUsage" %}}

{{% comment %}}

```yaml
alert: NetdataHighCpuUsage
expr: rate(netdata_cpu_cpu_percentage_average{dimension="idle"}[1m]) > 80
for: 5m
labels:
    severity: warning
annotations:
    summary: Netdata high cpu usage (instance {{ $labels.instance }})
    description: |-
        Netdata high CPU usage (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataHighCpuUsage.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `NetdataHighCpuUsage`:

## Meaning

The `NetdataHighCpuUsage` alert is triggered when the average CPU usage of the Netdata instance exceeds 80% over a 1-minute period. This alert indicates that the Netdata instance is experiencing high CPU utilization, which may impact its performance and ability to collect and process metrics.

## Impact

High CPU usage on the Netdata instance can lead to:

* Delays or loss of metric data
* Increased latency in data processing and visualization
* Potential crashes or instability of the Netdata instance
* Impaired ability to monitor and troubleshoot system performance

## Diagnosis

To diagnose the root cause of high CPU usage on the Netdata instance:

1. Check the Netdata instance logs for errors or warnings related to CPU usage.
2. Verify that the system resources (CPU, memory, etc.) are sufficient to support the Netdata instance.
3. Investigate any recent changes to the system configuration, plugins, or metric collection settings that may be contributing to the high CPU usage.
4. Use Netdata's built-in dashboards and charts to identify any specific metrics or plugins that are consuming excessive CPU resources.

## Mitigation

To mitigate high CPU usage on the Netdata instance:

1. Optimize system resources: Ensure that the system has sufficient CPU and memory resources to support the Netdata instance.
2. Adjust plugin settings: Review and adjust plugin settings to reduce CPU usage, such as reducing the frequency of metric collection or disabling unnecessary plugins.
3. Implement caching: Consider implementing caching mechanisms to reduce the load on the Netdata instance.
4. Upgrade Netdata: Ensure that the Netdata instance is running the latest version, as newer versions may include performance optimizations.
5. Consider load balancing: If the Netdata instance is handling a high volume of metrics, consider load balancing the traffic across multiple instances to reduce the CPU load on individual instances.