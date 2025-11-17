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
expr: (time() - opnsense_wireguard_peer_last_handshake_seconds) > 300
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
The OPNsenseWireGuardPeerStaleHandshake alert is triggered when a WireGuard peer has not performed a handshake in over 5 minutes. This indicates a potential issue with the WireGuard connection, which could be caused by a variety of factors such as network connectivity problems, configuration errors, or peer downtime. The alert is classified as a warning, suggesting that the issue is not critical but should be investigated and resolved to prevent potential disruptions to the network.

## Impact
The impact of this alert is that the WireGuard peer is not able to establish a secure connection, which could lead to a range of problems including:
* Loss of connectivity to the peer
* Inability to access resources or services behind the peer
* Potential security vulnerabilities if the peer is not able to authenticate or encrypt traffic
* Downtime or disruptions to critical services or applications that rely on the WireGuard connection

## Diagnosis
To diagnose the issue, the following steps can be taken:
1. Check the WireGuard peer status and configuration to ensure that it is correctly set up and enabled.
2. Verify network connectivity to the peer and check for any firewalls or network devices that may be blocking the connection.
3. Review the WireGuard logs to check for any error messages or hints about the cause of the issue.
4. Check the system clocks on both the OPNsense device and the WireGuard peer to ensure that they are synchronized, as a clock skew can cause handshake issues.
5. Attempt to manually initiate a handshake or restart the WireGuard service to see if the issue resolves itself.

## Mitigation
To mitigate the issue, the following steps can be taken:
1. Investigate and resolve any underlying network connectivity issues that may be preventing the handshake.
2. Verify and update the WireGuard configuration to ensure that it is correct and up-to-date.
3. Restart the WireGuard service or reboot the peer device to reset the connection.
4. Monitor the WireGuard logs and system metrics to detect any potential issues or errors that may be contributing to the problem.
5. Consider implementing additional monitoring or alerting to detect similar issues in the future and prevent downtime or disruptions.