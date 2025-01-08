---
title: HostNetworkBondDegraded
description: Troubleshooting for alert HostNetworkBondDegraded
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNetworkBondDegraded

Bond "{{ $labels.device }}" degraded on "{{ $labels.instance }}".

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostNetworkBondDegraded" %}}

{{% comment %}}

```yaml
alert: HostNetworkBondDegraded
expr: ((node_bonding_active - node_bonding_slaves) != 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host Network Bond Degraded (instance {{ $labels.instance }})
    description: |-
        Bond "{{ $labels.device }}" degraded on "{{ $labels.instance }}".
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostNetworkBondDegraded.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the `HostNetworkBondDegraded` alert:

## Meaning

The `HostNetworkBondDegraded` alert is triggered when the number of active bonding interfaces on a host does not match the number of slave interfaces. This indicates that the network bonding configuration on the host is in a degraded state, which can lead to network connectivity issues or failures.

## Impact

A degraded network bond can cause:

* Network connectivity issues
* Packet loss or corruption
* Slower network performance
* Uneven traffic distribution across bonding interfaces

This can impact critical workloads and services running on the host, leading to downtime, data loss, or other business disruptions.

## Diagnosis

To diagnose the issue:

1. Check the bonding configuration on the affected host using commands like `cat /proc/net/bonding/bondX` (where X is the bond interface number).
2. Verify that the number of active and slave interfaces matches the expected configuration.
3. Check system logs for any errors or warnings related to bonding or network interfaces.
4. Use tools like `ip link show` or `ip -d link show` to inspect the bonding interface configuration and status.

## Mitigation

To mitigate the issue:

1. Review the bonding configuration and adjust it according to the expected setup.
2. Restart the bonding service or reload the bonding configuration to apply changes.
3. Verify that the number of active and slave interfaces matches the expected configuration.
4. Monitor the host's network performance and adjust the bonding configuration as needed to ensure optimal performance and redundancy.
5. Consider implementing redundancy and failover mechanisms for critical network interfaces to minimize downtime and data loss.

Remember to follow your organization's change management procedures and testing protocols before making any changes to the bonding configuration.