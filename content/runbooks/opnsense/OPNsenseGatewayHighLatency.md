---
title: OPNsenseGatewayHighLatency
description: Troubleshooting for alert OPNsenseGatewayHighLatency
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseGatewayHighLatency

Gateway {{ $labels.name }} has high latency (> 100ms)

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseGatewayHighLatency" %}}

{{% comment %}}

```yaml
alert: OPNsenseGatewayHighLatency
expr: opnsense_gateways_rtt_milliseconds > 100
for: 5m
labels:
    severity: warning
annotations:
    summary: OPNsense gateway high latency (instance {{ $labels.opnsense_instance }})
    description: |-
        Gateway {{ $labels.name }} has high latency (> 100ms)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsensegatewayhighlatency/

```

{{% /comment %}}

</details>


## Meaning
The OPNsenseGatewayHighLatency alert is triggered when the round-trip time (RTT) of an OPNsense gateway exceeds 100 milliseconds for a period of 5 minutes. This indicates a potential issue with the gateway's performance or connectivity, which may impact network availability and responsiveness.

## Impact
The high latency of the OPNsense gateway can have significant impacts on network operations, including:
* Slow data transfer rates
* Increased packet loss
* Poor performance of applications that rely on the gateway
* Potential downtime or disconnections for users
* Negative impact on overall network reliability and uptime

## Diagnosis
To diagnose the issue, follow these steps:
1. **Check gateway configuration**: Verify that the OPNsense gateway is properly configured and that there are no misconfigurations that could be causing the high latency.
2. **Monitor network traffic**: Analyze network traffic patterns to identify any potential bottlenecks or issues that could be contributing to the high latency.
3. **Check for hardware issues**: Verify that the gateway's hardware is functioning properly and that there are no issues with the network interfaces or other components.
4. **Verify internet connectivity**: Check that the gateway has a stable internet connection and that there are no issues with the upstream provider.
5. **Check system logs**: Review system logs to identify any error messages or warnings that could indicate the cause of the high latency.

## Mitigation
To mitigate the high latency issue, consider the following steps:
1. **Optimize gateway configuration**: Review and optimize the OPNsense gateway configuration to ensure that it is properly tuned for the network environment.
2. **Upgrade gateway hardware**: Consider upgrading the gateway's hardware to improve its performance and reduce latency.
3. **Implement Quality of Service (QoS) policies**: Implement QoS policies to prioritize critical traffic and ensure that it is not impacted by the high latency.
4. **Divert traffic to alternative gateways**: If possible, divert traffic to alternative gateways to reduce the load on the affected gateway and minimize the impact on network operations.
5. **Contact upstream provider**: If the issue is related to the upstream provider, contact their support team to report the issue and request assistance in resolving it.