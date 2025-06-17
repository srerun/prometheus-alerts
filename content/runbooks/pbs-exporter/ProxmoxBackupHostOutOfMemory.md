---
title: ProxmoxBackupHostOutOfMemory
description: Troubleshooting for alert ProxmoxBackupHostOutOfMemory
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupHostOutOfMemory

Node memory is filling up (< 10% left)

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupHostOutOfMemory" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupHostOutOfMemory
expr: pbs_host_memory_used / (pbs_host_memory_used + pbs_host_memory_free) * 100 > 90
for: 2m
labels:
    severity: warning
annotations:
    summary: Proxmox host out of memory (instance {{ $labels.instance }})
    description: |-
        Node memory is filling up (< 10% left)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuphostoutofmemory/

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `ProxmoxBackupHostOutOfMemory`:

## Meaning

The `ProxmoxBackupHostOutOfMemory` alert indicates that a Proxmox host is experiencing low memory availability. Specifically, the alert is triggered when the percentage of used memory exceeds 90% for 2 minutes or more. This can lead to performance issues, slow backups, and even crashes.

## Impact

The impact of this alert can be significant, as it may:

* Cause backups to fail or slow down, resulting in data loss or inconsistencies
* Lead to performance issues on the Proxmox host, and its dependent services
* Potentially crash the Proxmox host, leading to downtime and service unavailability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox host's resource utilization (CPU, memory, disk usage) to identify the cause of the memory shortage.
2. Verify that there are no resource-intensive processes or applications running on the host.
3. Check the backup configuration and adjust it if necessary to reduce memory usage.
4. Consider adding more memory to the Proxmox host or migrating to a host with more resources.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately investigate and address the cause of the memory shortage.
2. Implement a memory allocation strategy to ensure that critical services receive sufficient memory.
3. Consider setting up memory reservations or limits for critical applications.
4. Monitor the Proxmox host's memory usage closely to prevent similar issues in the future.

Note: This runbook is meant to serve as a starting point and might need to be adapted to your specific environment and setup.