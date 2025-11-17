---
title: OPNsenseGatewayOffline
description: Troubleshooting for alert OPNsenseGatewayOffline
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseGatewayOffline

Gateway {{ $labels.name }} ({{ $labels.address }}) is offline

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseGatewayOffline" %}}

{{% comment %}}

```yaml
alert: OPNsenseGatewayOffline
expr: opnsense_gateways_status != 1
for: 2m
labels:
    severity: critical
annotations:
    summary: OPNsense gateway offline (instance {{ $labels.opnsense_instance }})
    description: |-
        Gateway {{ $labels.name }} ({{ $labels.address }}) is offline
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsensegatewayoffline/

```

{{% /comment %}}

</details>


## Meaning
The OPNsenseGatewayOffline alert is triggered when the status of an OPNsense gateway is not equal to 1, indicating that the gateway is offline. This alert is critical, as it may cause network connectivity issues and impact the availability of services relying on this gateway.

## Impact
The impact of this alert is significant, as it may lead to:
* Loss of network connectivity for devices and services relying on the affected gateway
* Inability to access external resources, such as the internet
* Potential security risks, as the offline gateway may not be able to filter or block malicious traffic
* Degraded performance, as traffic may be rerouted through alternative gateways, increasing latency and decreasing overall network efficiency

## Diagnosis
To diagnose the issue, follow these steps:
1. Check the OPNsense instance logs for any error messages related to the gateway
2. Verify the gateway's configuration and ensure it is correctly set up
3. Check the physical connectivity of the gateway, including power and network cables
4. Use network monitoring tools to verify the gateway's status and identify any issues
5. Investigate any recent changes or updates that may have caused the gateway to go offline

## Mitigation
To mitigate the issue, follow these steps:
1. Check the OPNsense instance for any configuration issues or errors that may be causing the gateway to be offline
2. Restart the gateway or the OPNsense instance to see if it resolves the issue
3. If the issue persists, investigate alternative gateways or network paths to maintain network connectivity
4. Verify that the gateway is properly configured and connected to the network
5. Consider implementing redundant gateways or high-availability solutions to minimize the impact of a single gateway failure
6. Refer to the OPNsense documentation and community resources for troubleshooting and configuration guidance.