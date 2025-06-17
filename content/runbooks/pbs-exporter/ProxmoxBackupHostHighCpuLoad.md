---
title: ProxmoxBackupHostHighCpuLoad
description: Troubleshooting for alert ProxmoxBackupHostHighCpuLoad
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupHostHighCpuLoad

CPU load is > 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupHostHighCpuLoad" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupHostHighCpuLoad
expr: avg_over_time(pbs_host_cpu_usage[2m]) > 0.8
for: 10m
labels:
    severity: warning
annotations:
    summary: Host high CPU load (id {{ $labels.id }})
    description: |-
        CPU load is > 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuphosthighcpuload/

```

{{% /comment %}}

</details>


## Meaning

The ProxmoxBackupHostHighCpuLoad alert is triggered when the average CPU usage of a Proxmox backup host over a 2-minute period exceeds 80%. This indicates that the system is experiencing high CPU load, which can lead to performance issues and potential disruptions to backup operations.

## Impact

The impact of this alert is high, as high CPU load can cause:

* Slow backup performance
* Increased risk of backup failures
* Potential data loss or corruption
* Inability to meet backup windows, leading to compliance issues

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox backup host's resource utilization, including CPU, memory, and disk usage.
2. Investigate any recent changes to the backup configuration, such as new backup jobs or increased data volume.
3. Review system logs for any error messages or warning signs of hardware failure.
4. Verify that the backup host's resources are sufficient to handle the current workload.

## Mitigation

To mitigate the issue, follow these steps:

1. Check for any resource-intensive processes or tasks running on the backup host and terminate or reschedule them as necessary.
2. Consider adding more resources (e.g., CPU, memory, or disk space) to the backup host to handle the increased workload.
3. Optimize backup jobs to reduce the impact on the backup host.
4. Consider implementing redundancy or load balancing to distribute the backup workload across multiple hosts.

 Additionally, refer to the [runbook](https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuphosthighcpuload/) for more detailed guidance and best practices for resolving this alert.