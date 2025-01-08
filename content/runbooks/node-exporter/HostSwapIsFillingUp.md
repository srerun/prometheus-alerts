---
title: HostSwapIsFillingUp
description: Troubleshooting for alert HostSwapIsFillingUp
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostSwapIsFillingUp

Swap is filling up (>80%)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostSwapIsFillingUp" %}}

{{% comment %}}

```yaml
alert: HostSwapIsFillingUp
expr: ((1 - (node_memory_SwapFree_bytes / node_memory_SwapTotal_bytes)) * 100 > 80) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host swap is filling up (instance {{ $labels.instance }})
    description: |-
        Swap is filling up (>80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostSwapIsFillingUp.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule "HostSwapIsFillingUp":

## Meaning

The "HostSwapIsFillingUp" alert is triggered when the swap usage on a host exceeds 80% for more than 2 minutes. This alert is raised by the Node Exporter, which monitors system metrics, including memory and swap usage.

## Impact

A high swap usage can lead to performance degradation, slow response times, and even system crashes. When the system runs low on physical memory, it uses the swap space to store data temporarily. However, excessive swap usage can cause the system to slow down, leading to:

* Slow application responses
* Increased latency
* Decreased system responsiveness
* Potential system crashes or freezes

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the system's memory usage using tools like `top`, `htop`, or `free -h`.
2. Identify any resource-intensive processes or applications consuming excessive memory.
3. Verify that the system's swap partition has sufficient space.
4. Investigate if there are any memory leaks or inefficient memory allocation in applications.

## Mitigation

To mitigate the issue, follow these steps:

1. **Identify and terminate memory-intensive processes**: Use `pkill` or `kill` commands to stop processes consuming excessive memory.
2. **Increase the system's physical memory**: Consider adding more RAM to the system to reduce the reliance on swap space.
3. **Optimize system configuration**: Adjust system settings to optimize memory usage, such as tuning the `vm.swappiness` kernel parameter.
4. **Implement memory-efficient applications**: Ensure applications are optimized for memory usage and fix any memory leaks.
5. **Monitor system resources**: Continuously monitor system resources, including memory and swap usage, to catch potential issues before they become critical.

Remember to investigate the root cause of the issue and address it accordingly to prevent future occurrences.