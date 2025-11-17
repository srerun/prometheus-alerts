---
title: OPNsenseGatewayPacketLoss
description: Troubleshooting for alert OPNsenseGatewayPacketLoss
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseGatewayPacketLoss

Gateway {{ $labels.name }} has packet loss (> 1%)

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseGatewayPacketLoss" %}}

{{% comment %}}

```yaml
alert: OPNsenseGatewayPacketLoss
expr: opnsense_gateways_loss_percentage > 1
for: 5m
labels:
    severity: warning
annotations:
    summary: OPNsense gateway packet loss (instance {{ $labels.opnsense_instance }})
    description: |-
        Gateway {{ $labels.name }} has packet loss (> 1%)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsensegatewaypacketloss/

```

{{% /comment %}}

</details>


## Meaning
The OPNsenseGatewayPacketLoss alert is triggered when the packet loss percentage of an OPNsense gateway exceeds 1% over a period of 5 minutes. This alert indicates that there is an issue with the gateway's connectivity, which may be causing network instability or disruption. The alert provides information about the affected gateway, including its name and instance, as well as the current packet loss percentage.

## Impact
The impact of this alert can vary depending on the specific use case and network configuration. However, in general, packet loss can cause:
* Network congestion and slow data transfer rates
* Disruption to critical services and applications
* Decreased overall network reliability and uptime
* Potential security risks if the packet loss is caused by a malicious actor

## Diagnosis
To diagnose the issue, follow these steps:
1. Check the OPNsense gateway logs for any error messages related to network connectivity or packet loss.
2. Verify that the gateway is properly configured and that there are no issues with the underlying network infrastructure.
3. Use network monitoring tools to check for any signs of congestion, packet loss, or other issues on the network.
4. Check for any firmware or software updates for the OPNsense gateway and apply them if necessary.
5. If the issue persists, consider performing a network capture to analyze the traffic and identify the root cause of the packet loss.

## Mitigation
To mitigate the issue, follow these steps:
1. Verify that the network infrastructure is properly configured and that there are no issues with the underlying network.
2. Check for any signs of congestion or packet loss on the network and take steps to address them, such as adjusting Quality of Service (QoS) settings or upgrading network hardware.
3. Consider implementing additional network monitoring and alerting tools to provide more detailed insights into network performance and detect issues earlier.
4. If the issue is caused by a configuration error, update the OPNsense gateway configuration to correct the issue.
5. If the issue is caused by a hardware or software problem, consider replacing or upgrading the affected component to resolve the issue.