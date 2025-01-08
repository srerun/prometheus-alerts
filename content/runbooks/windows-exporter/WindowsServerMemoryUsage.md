---
title: WindowsServerMemoryUsage
description: Troubleshooting for alert WindowsServerMemoryUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerMemoryUsage

Memory usage is more than 90%

<details>
  <summary>Alert Rule</summary>

{{% rule "windows-server/windows-exporter.yml" "WindowsServerMemoryUsage" %}}

{{% comment %}}

```yaml
alert: WindowsServerMemoryUsage
expr: 100 - ((windows_os_physical_memory_free_bytes / windows_cs_physical_memory_bytes) * 100) > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Windows Server memory Usage (instance {{ $labels.instance }})
    description: |-
        Memory usage is more than 90%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/windows-exporter/WindowsServerMemoryUsage.md

```

{{% /comment %}}

</details>


Here is a runbook for the WindowsServerMemoryUsage alert rule:

## Meaning

The WindowsServerMemoryUsage alert indicates that the memory usage on a Windows server has exceeded 90%. This can lead to performance issues, slow response times, and even crashes if left unchecked.

## Impact

High memory usage can significantly impact the performance and reliability of the Windows server, leading to:

* Slow application response times
* Increased CPU usage
* Reduced system stability
* Potential crashes or freezes
* Impaired user experience

## Diagnosis

To diagnose the issue, follow these steps:

1. **Verify the alert**: Check the Prometheus graph to confirm that the memory usage is indeed above 90%.
2. **Check system logs**: Review system logs to identify any error messages or warnings related to memory issues.
3. **Investigate running processes**: Use tools like Task Manager or Process Explorer to identify which processes are consuming the most memory.
4. **Check for memory leaks**: Look for any signs of memory leaks or abnormal memory usage patterns.
5. **Verify system configuration**: Check the system configuration to ensure that memory settings are adequate for the server's workload.

## Mitigation

To mitigate the issue, follow these steps:

1. **Identify and terminate unnecessary processes**: Terminate any unnecessary processes consuming excessive memory.
2. **Optimize system configuration**: Adjust system configuration settings to optimize memory usage for the server's workload.
3. **Implement memory monitoring**: Set up regular memory usage monitoring to detect potential issues before they become critical.
4. **Consider capacity planning**: Review capacity planning to ensure the server has sufficient resources to handle its workload.
5. **Apply OS and software updates**: Ensure the operating system and software are up-to-date, as newer versions may include memory-related performance optimizations.