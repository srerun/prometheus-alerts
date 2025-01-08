---
title: WindowsServerCpuUsage
description: Troubleshooting for alert WindowsServerCpuUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerCpuUsage

CPU Usage is more than 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "windows-server/windows-exporter.yml" "WindowsServerCpuUsage" %}}

{{% comment %}}

```yaml
alert: WindowsServerCpuUsage
expr: 100 - (avg by (instance) (rate(windows_cpu_time_total{mode="idle"}[2m])) * 100) > 80
for: 0m
labels:
    severity: warning
annotations:
    summary: Windows Server CPU Usage (instance {{ $labels.instance }})
    description: |-
        CPU Usage is more than 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/windows-exporter/WindowsServerCpuUsage.md

```

{{% /comment %}}

</details>


Here is a runbook for the WindowsServerCpuUsage alert rule:

## Meaning

The WindowsServerCpuUsage alert rule is triggered when the CPU usage of a Windows server exceeds 80%. This alert is indicative of potential performance issues on the server, which can lead to slower response times, increased latency, and decreased overall system performance.

## Impact

High CPU usage on a Windows server can have several negative impacts on the system and its users. These include:

* Slower application response times
* Increased latency
* Decreased system performance
* Increased risk of system crashes or freezes
* Decreased overall user experience

## Diagnosis

To diagnose the root cause of high CPU usage on a Windows server, follow these steps:

1. Check the system's Task Manager to identify which processes are consuming the most CPU resources.
2. Review system logs to identify any errors or warnings related to CPU usage.
3. Check for any resource-intensive applications or services that may be causing the high CPU usage.
4. Verify that the server's hardware is sufficient to handle the current workload.
5. Check for any signs of malware or virus infections.

## Mitigation

To mitigate high CPU usage on a Windows server, follow these steps:

1. Identify and terminate any unnecessary or resource-intensive processes.
2. Optimize system settings to improve performance, such as adjusting power settings or disabling unnecessary features.
3. Consider upgrading the server's hardware to improve performance.
4. Implement load balancing or distribute workloads across multiple servers to reduce the load on individual servers.
5. Consider implementing CPU throttling or limiting CPU usage for specific applications or services.

Remember to always monitor system performance and adjust mitigation strategies as needed to ensure optimal system performance.