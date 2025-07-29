---
title: ProxmoxGuestDown
description: Troubleshooting for alert ProxmoxGuestDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxGuestDown

Guest {{ printf "{{ $labels.name }}" }} of type {{ printf "{{ $labels.type }}" }} on node {{ printf "{{ $labels.node }}" }} is down

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxGuestDown" %}}

{{% comment %}}

```yaml
alert: ProxmoxGuestDown
expr: |
    proxmox_guest_up == 0
for: 1m
labels:
    severity: critical
annotations:
    summary: Proxmox guest {{ printf "{{ $labels.name }}" }} is down
    description: Guest {{ printf "{{ $labels.name }}" }} of type {{ printf "{{ $labels.type }}" }} on node {{ printf "{{ $labels.node }}" }} is down
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxguestdown/

```

{{% /comment %}}

</details>


## Meaning

This runbook is for the ProxmoxGuestDown alert, which is triggered when a Proxmox guest virtual machine is down. This alert is critical and indicates that a guest machine is not functioning, which can impact services and applications that rely on it.

## Impact

* Unavailability of services or applications running on the affected guest machine
* Potential data loss or corruption if the guest machine was not properly shut down
* Impact on business operations and productivity if the guest machine is critical to daily activities
* Increased risk of security breaches if the guest machine is not properly secured

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Proxmox web interface or command-line interface to confirm the guest machine is down
2. Verify the guest machine's configuration and settings to ensure they are correct
3. Review the system logs to identify any errors or issues that may have caused the guest machine to go down
4. Check for any hardware or software issues on the node that the guest machine is running on
5. Verify that the Proxmox node is properly configured and running correctly

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the guest machine from the Proxmox web interface or command-line interface
2. If the guest machine does not restart, attempt to migrate it to another node or reboot the node it is running on
3. If the issue persists, restore the guest machine from a backup or rebuild it from scratch
4. Verify that all services and applications running on the guest machine are functioning correctly
5. Take preventative measures to avoid similar issues in the future, such as:
	* Regularly backing up guest machines
	* Implementing high availability and redundancy for critical guest machines
	* Monitoring guest machine performance and logs for signs of issues
	* Performing regular maintenance and updates on the Proxmox node and guest machines