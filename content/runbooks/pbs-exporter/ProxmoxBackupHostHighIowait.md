---
title: ProxmoxBackupHostHighIowait
description: Troubleshooting for alert ProxmoxBackupHostHighIowait
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupHostHighIowait

CPU iowait > 10%. A high iowait means that you are disk or network bound.

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupHostHighIowait" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupHostHighIowait
expr: avg by (instance) (rate(pbs_host_io_wait[5m]) * 100) > 10
for: 0m
labels:
    severity: warning
annotations:
    summary: Host CPU high iowait (instance {{ $labels.instance }})
    description: |-
        CPU iowait > 10%. A high iowait means that you are disk or network bound.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuphosthighiowait/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The ProxmoxBackupHostHighIowait alert is triggered when the average CPU iowait percentage exceeds 10% over a 5-minute period for a specific host. CPU iowait measures the percentage of CPU time spent waiting for disk or network I/O operations to complete. High CPU iowait indicates that the host is disk or network-bound, leading to performance issues and potential slowdowns.

## Impact

The impact of high CPU iowait can be significant, leading to:

* Slow response times for users
* Decreased system performance
* Potential crashes or freezes
* Inefficient use of system resources

## Diagnosis

To diagnose the root cause of high CPU iowait, follow these steps:

1. **Check disk usage**: Verify that disk usage is within acceptable limits. High disk usage can cause high iowait.
2. **Check network usage**: Verify that network usage is within acceptable limits. High network usage can cause high iowait.
3. **Check system resource utilization**: Verify that system resources such as CPU, memory, and swap are not overutilized.
4. **Check disk performance**: Verify that disk performance is adequate to handle the workload.
5. **Check for resource-intensive processes**: Identify and troubleshoot resource-intensive processes that may be contributing to high CPU iowait.

## Mitigation

To mitigate high CPU iowait, perform the following steps:

1. **Optimize disk usage**: Implement disk usage optimization techniques, such as compressing data, removing unnecessary files, and upgrading to faster storage.
2. **Optimize network usage**: Implement network usage optimization techniques, such as upgrading network infrastructure, optimizing network protocols, and reducing network traffic.
3. **Upgrade system resources**: Upgrade system resources such as CPU, memory, and swap to handle increased workloads.
4. **Optimize system configuration**: Optimize system configuration to ensure efficient use of system resources.
6. **Schedule resource-intensive tasks**: Schedule resource-intensive tasks during off-peak hours to minimize the impact on the system.

Remember to monitor the system closely after implementing these steps to ensure that the issue is resolved. If the issue persists, consider escalating to a subject matter expert for further assistance.