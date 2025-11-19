---
title: OPNsenseWireGuardPeerOffline
description: Troubleshooting for alert OPNsenseWireGuardPeerOffline
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseWireGuardPeerOffline

WireGuard peer {{ $labels.peer_name }} is offline

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseWireGuardPeerOffline" %}}

{{% comment %}}

```yaml
alert: OPNsenseWireGuardPeerOffline
expr: opnsense_wireguard_peer_status == 0 and on(device_name) opnsense_wireguard_interfaces_status == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: OPNsense WireGuard peer offline (instance {{ $labels.opnsense_instance }})
    description: |-
        WireGuard peer {{ $labels.peer_name }} is offline
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsensewireguardpeeroffline/

```

{{% /comment %}}

</details>


## Meaning
The `OPNsenseWireGuardPeerOffline` alert is triggered when a WireGuard peer connected to an OPNsense device is detected as being offline. This alert is critical, indicating a significant impact on the network connectivity and security. The alert is generated when the `opnsense_wireguard_peer_status` metric is 0 (indicating an offline state) and the `opnsense_wireguard_interfaces_status` metric is 1 (indicating that the WireGuard interface is up) for a period of 2 minutes.

## Impact
The impact of this alert is that the WireGuard peer is unable to establish or maintain a connection to the OPNsense device, potentially disrupting secure communication between the peer and the device. This could lead to a loss of connectivity, data breaches, or other security issues. The critical severity of this alert highlights the importance of promptly addressing the issue to minimize the impact on the network and its users.

## Diagnosis
To diagnose the issue, the following steps can be taken:
1. **Verify the peer configuration**: Check the WireGuard peer configuration on the OPNsense device to ensure that it is correctly set up and enabled.
2. **Check the peer status**: Use the OPNsense web interface or CLI to check the current status of the WireGuard peer and verify that it is indeed offline.
3. **Inspect the WireGuard logs**: Review the WireGuard logs on the OPNsense device to identify any error messages or issues that may be causing the peer to be offline.
4. **Verify network connectivity**: Check the network connectivity between the WireGuard peer and the OPNsense device to ensure that there are no issues with the underlying network infrastructure.

## Mitigation
To mitigate the issue, the following steps can be taken:
1. **Restart the WireGuard service**: Restart the WireGuard service on the OPNsense device to see if it resolves the issue.
2. **Verify the peer's public key**: Verify that the peer's public key is correctly configured on the OPNsense device and that it matches the key used by the peer.
3. **Check for firewall rules**: Check the firewall rules on the OPNsense device to ensure that they are not blocking the WireGuard traffic.
4. **Contact the peer administrator**: If the issue persists, contact the administrator of the WireGuard peer to verify their configuration and ensure that they are not experiencing any issues on their end.
5. **Refer to the OPNsense documentation**: Consult the OPNsense documentation and support resources for further troubleshooting and configuration guidance.