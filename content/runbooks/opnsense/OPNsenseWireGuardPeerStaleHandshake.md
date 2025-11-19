---
title: OPNsenseWireGuardPeerStaleHandshake
description: Troubleshooting for alert OPNsenseWireGuardPeerStaleHandshake
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseWireGuardPeerStaleHandshake

WireGuard peer {{ $labels.peer_name }} has not handshaked in > 5 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseWireGuardPeerStaleHandshake" %}}

{{% comment %}}

```yaml
alert: OPNsenseWireGuardPeerStaleHandshake
expr: (time() - opnsense_wireguard_peer_last_handshake_seconds) > 300 and on(device_name) opnsense_wireguard_interfaces_status == 1
for: 5m
labels:
    severity: warning
annotations:
    summary: OPNsense WireGuard peer stale handshake (instance {{ $labels.opnsense_instance }})
    description: |-
        WireGuard peer {{ $labels.peer_name }} has not handshaked in > 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsensewireguardpeerstalehandshake/

```

{{% /comment %}}

</details>


## Meaning
The OPNsenseWireGuardPeerStaleHandshake alert is triggered when a WireGuard peer has not performed a handshake in over 5 minutes. This indicates a potential issue with the WireGuard connection, such as a network connectivity problem or a configuration error. The alert is raised when the time since the last handshake exceeds 300 seconds (5 minutes) and the WireGuard interface is up.

## Impact
The impact of this alert is that the WireGuard peer is not able to establish a secure connection, which can lead to a loss of connectivity and potential security risks. If the peer is not able to handshake, it may not be able to send or receive data, which can cause disruptions to critical services or applications. Additionally, a stale handshake can indicate a more serious underlying issue, such as a network connectivity problem or a configuration error, which can have broader implications for the overall network.

## Diagnosis
To diagnose the issue, the following steps can be taken:
1. Check the WireGuard peer configuration to ensure that it is correctly configured and enabled.
2. Verify that the WireGuard interface is up and running.
3. Check the network connectivity between the WireGuard peer and the OPNsense device.
4. Review the WireGuard logs to see if there are any error messages or indications of issues.
5. Check the system resources (e.g., CPU, memory, disk space) to ensure that they are not causing issues.

## Mitigation
To mitigate the issue, the following steps can be taken:
1. Restart the WireGuard peer or the OPNsense device to see if it resolves the issue.
2. Check and update the WireGuard configuration to ensure that it is correct and up-to-date.
3. Verify that the network connectivity is stable and not causing issues.
4. Increase the logging level for WireGuard to gather more detailed information about the issue.
5. Consider implementing additional monitoring or alerting to detect similar issues in the future.
By taking these steps, the issue can be diagnosed and mitigated, and the WireGuard peer can be restored to a healthy state.