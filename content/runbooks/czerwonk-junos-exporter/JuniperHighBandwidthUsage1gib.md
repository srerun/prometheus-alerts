---
title: JuniperHighBandwidthUsage1gib
description: Troubleshooting for alert JuniperHighBandwidthUsage1gib
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# JuniperHighBandwidthUsage1gib

Interface is highly saturated. (> 0.90GiB/s)

<details>
  <summary>Alert Rule</summary>

{{% rule "juniper/czerwonk-junos-exporter.yml" "JuniperHighBandwidthUsage1gib" %}}

{{% comment %}}

```yaml
alert: JuniperHighBandwidthUsage1gib
expr: rate(junos_interface_transmit_bytes[1m]) * 8 > 1e+9 * 0.90
for: 1m
labels:
    severity: critical
annotations:
    summary: Juniper high Bandwidth Usage 1GiB (instance {{ $labels.instance }})
    description: |-
        Interface is highly saturated. (> 0.90GiB/s)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/czerwonk-junos-exporter/JuniperHighBandwidthUsage1gib.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule:

## Meaning

The **JuniperHighBandwidthUsage1gib** alert is triggered when the transmit bandwidth of a Juniper interface exceeds 90% of its 1GiB capacity. This alert indicates that the interface is highly saturated, which can lead to network congestion, packet loss, and poor performance.

## Impact

* Network congestion and packet loss
* Poor network performance and reliability
* Potential impact on critical business applications and services
* Increased risk of network downtime and outages

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check the interface usage**: Verify the transmit bandwidth usage of the affected interface using the `junos_interface_transmit_bytes` metric.
2. **Identify the traffic source**: Analyze the network traffic to determine the source of the high bandwidth usage (e.g., specific applications, hosts, or networks).
3. **Check for configuration issues**: Review the Juniper device configuration to ensure that there are no misconfigured settings or policies that could be contributing to the high bandwidth usage.
4. **Verify device resource utilization**: Check the CPU and memory utilization of the Juniper device to ensure that it is not overwhelmed.

## Mitigation

To mitigate the issue, follow these steps:

1. **Implement traffic shaping or policing**: Configure traffic shaping or policing policies on the Juniper device to limit the bandwidth usage of specific applications or hosts.
2. **Optimize network configuration**: Review and optimize the network configuration to ensure that it is aligned with best practices and is not contributing to the high bandwidth usage.
3. **Upgrade device resources**: Consider upgrading the Juniper device resources (e.g., CPU, memory) to ensure that it can handle the increased bandwidth demands.
4. **Monitor and analyze network traffic**: Continuously monitor and analyze network traffic to identify and address any potential issues before they become critical.