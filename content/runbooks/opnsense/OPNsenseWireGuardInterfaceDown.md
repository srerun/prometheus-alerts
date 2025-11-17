---
title: OPNsenseWireGuardInterfaceDown
description: Troubleshooting for alert OPNsenseWireGuardInterfaceDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseWireGuardInterfaceDown

WireGuard interface {{ $labels.device_name }} is down

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseWireGuardInterfaceDown" %}}

{{% comment %}}

```yaml
alert: OPNsenseWireGuardInterfaceDown
expr: opnsense_wireguard_interfaces_status == 0
for: 2m
labels:
    severity: critical
annotations:
    summary: OPNsense WireGuard interface down (instance {{ $labels.opnsense_instance }})
    description: |-
        WireGuard interface {{ $labels.device_name }} is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsensewireguardinterfacedown/

```

{{% /comment %}}

</details>


## Meaning
The `OPNsenseWireGuardInterfaceDown` alert is triggered when a WireGuard interface on an OPNsense device is down. This alert is critical and indicates that the WireGuard interface is not functioning as expected, which can lead to disruptions in network connectivity and security.

## Impact
The impact of this alert is significant, as a down WireGuard interface can cause:
* Loss of secure network connectivity
* Disruption of critical network services
* Exposure to security risks due to the lack of encryption and authentication
* Potential data breaches or unauthorized access

## Diagnosis
To diagnose the issue, follow these steps:
1. Check the OPNsense device's system logs for any error messages related to WireGuard.
2. Verify the WireGuard interface configuration and ensure that it is correctly set up.
3. Check the device's network connections and ensure that the WireGuard interface is properly connected.
4. Verify that the WireGuard peer is reachable and responding.
5. Check for any firewall or network policies that may be blocking the WireGuard traffic.

## Mitigation
To mitigate the issue, follow these steps:
1. Check the WireGuard interface configuration and correct any errors or misconfigurations.
2. Restart the WireGuard service on the OPNsense device.
3. Verify that the WireGuard peer is reachable and responding, and restart the peer's WireGuard service if necessary.
4. Check for any firewall or network policies that may be blocking the WireGuard traffic and update them accordingly.
5. If the issue persists, consider restarting the OPNsense device or seeking further assistance from a network administrator or OPNsense support specialist.