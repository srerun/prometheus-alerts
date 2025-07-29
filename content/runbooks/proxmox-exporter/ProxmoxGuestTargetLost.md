---
title: ProxmoxGuestTargetLost
description: Troubleshooting for alert ProxmoxGuestTargetLost
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxGuestTargetLost

Guest {{ printf "{{ $labels.name }}" }} of type {{ printf "{{ $labels.type }}" }} on node {{ printf "{{ $labels.node }}" }} may be down

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxGuestTargetLost" %}}

{{% comment %}}

```yaml
alert: ProxmoxGuestTargetLost
expr: |
    absent_over_time(proxmox_guest_up[1h])
for: 1m
labels:
    severity: critical
annotations:
    summary: Proxmox guest up metric absent for {{ printf "{{ $labels.name }}" }}
    description: Guest {{ printf "{{ $labels.name }}" }} of type {{ printf "{{ $labels.type }}" }} on node {{ printf "{{ $labels.node }}" }} may be down
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxguesttargetlost/

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule `ProxmoxGuestTargetLost`:

## Meaning

The `ProxmoxGuestTargetLost` alert is triggered when the `proxmox_guest_up` metric is absent for 1 hour, indicating that a Proxmox guest virtual machine is potentially down. This alert is critical and requires immediate attention to prevent downtime and data loss.

## Impact

The impact of this alert is high, as it indicates that a guest virtual machine is unavailable, which can result in:

* Downtime and loss of productivity for users relying on the virtual machine
* Potential data loss or corruption if the virtual machine was not shut down cleanly
* Increased workload for IT teams to troubleshoot and resolve the issue

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox web interface to verify the status of the guest virtual machine.
2. Review the Proxmox logs to identify any errors or issues related to the guest virtual machine.
3. Check the node's system logs to identify any underlying hardware or software issues.
4. Verify that the Proxmox node is reachable and responding to requests.
5. Check the network connectivity between the Proxmox node and the guest virtual machine.

## Mitigation

To mitigate the issue, follow these steps:

1. Attempt to restart the guest virtual machine from the Proxmox web interface.
2. If the guest virtual machine cannot be restarted, try to migrate it to a different node or host.
3. If migration is not possible, restore the guest virtual machine from a backup (if available).
4. Investigate and resolve any underlying issues causing the guest virtual machine to become unavailable.
5. Implement proactive monitoring and alerting to prevent similar issues in the future.

Remember to update the runbook with specific steps and procedures relevant to your organization's infrastructure and policies.