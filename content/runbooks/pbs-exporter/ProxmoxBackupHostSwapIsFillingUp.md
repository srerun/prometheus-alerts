---
title: ProxmoxBackupHostSwapIsFillingUp
description: Troubleshooting for alert ProxmoxBackupHostSwapIsFillingUp
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxBackupHostSwapIsFillingUp

Swap is filling up (> 80%)

<details>
  <summary>Alert Rule</summary>

{{% rule "pbs-exporter/pbs-exporter.yml" "ProxmoxBackupHostSwapIsFillingUp" %}}

{{% comment %}}

```yaml
alert: ProxmoxBackupHostSwapIsFillingUp
expr: pbs_host_swap_used / (pbs_host_swap_used + pbs_host_swap_free) * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: Proxmox host swap is filling up  (instance {{ $labels.instance }})
    description: |-
        Swap is filling up (> 80%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/pbs-exporter/proxmoxbackuphostswapisfillingup/

```

{{% /comment %}}

</details>


Here is a runbook with the specified sections for the Prometheus alert rule:

## Meaning

This alert indicates that the Proxmox host's swap space is filling up. The swap space is a critical system resource that is used when the system runs low on physical RAM. If the swap space is filling up, it may indicate a memory leak or a configuration issue.

## Impact

If left unattended, a full swap space can cause the system to slow down, crash, or become unresponsive. This can lead to:

* Downtime of critical services
* Data loss or corruption
* Increased security risks due to system instability
* Decreased overall system performance and reliability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Prometheus dashboard for any memory-related metrics that may indicate a memory leak or spike.
2. Investigate the system's memory usage and identify any processes or applications that are consuming excessive memory.
3. Verify that the system's swap configuration is correct and that the swap space is not too small.
4. Check the system logs for any error messages related to memory or swap space.

## Mitigation

To mitigate the issue, follow these steps:

1. Increase the system's RAM to reduce the likelihood of the system using swap space.
2. Identify and terminate any memory-leaking processes or applications.
3. Optimize system configurations to reduce memory usage.
4. Consider implementing memory monitoring and alerting to detect potential issues earlier.
5. If the issue persists, consider increasing the swap space or implementing more robust memory management strategies.

Remember to update this runbook with any specific details relevant to your environment and to keep it up-to-date with any changes to your systems or processes.