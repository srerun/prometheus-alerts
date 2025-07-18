---
title: HighIPv4DroppedPackets
description: Troubleshooting for alert HighIPv4DroppedPackets
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HighIPv4DroppedPackets

IPv4 dropped packets (up or down) for user {{ $labels.vbng_userName }} (index: {{ $labels.vbng_userIndex }}) exceed 100 for more than 5 minutes (current value: {{ $value }}). This may indicate network congestion or configuration issues.

<details>
  <summary>Alert Rule</summary>

{{% rule "netelastic/netelastic.yml" "HighIPv4DroppedPackets" %}}

{{% comment %}}

```yaml
alert: HighIPv4DroppedPackets
expr: vbng_ipv4_down_dropped_packets > 100 or vbng_ipv4_up_dropped_packets > 100
for: 5m
labels:
    severity: warning
annotations:
    summary: High IPv4 dropped packets
    description: 'IPv4 dropped packets (up or down) for user {{ $labels.vbng_userName }} (index: {{ $labels.vbng_userIndex }}) exceed 100 for more than 5 minutes (current value: {{ $value }}). This may indicate network congestion or configuration issues.'
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/netelastic/highipv4droppedpackets/

```

{{% /comment %}}

</details>


Here is the runbook for the HighIPv4DroppedPackets alert:

## Meaning

This runbook is intended to guide the troubleshooting and resolution of the HighIPv4DroppedPackets alert, which is triggered when the number of IPv4 dropped packets (either upstream or downstream) exceeds 100 for more than 5 minutes.

## Impact

The impact of this issue is potential network congestion or configuration problems, which can lead to:

* Poor network performance
* Packet loss
* Degraded user experience
* Increased latency
* Decreased throughput

## Diagnosis

To diagnose the root cause of the issue, follow these steps:

1. **Check network utilization**: Verify if the network is experiencing high utilization, which could be causing packet drops.
2. **Verify routing and forwarding**: Check if there are any issues with routing or forwarding that could be causing packets to be dropped.
3. **Investigate configuration issues**: Review the network configuration to ensure it is correct and not causing packet drops.
4. **Check for firmware or software issues**: Verify if there are any known firmware or software issues that could be contributing to the packet drops.
5. **Gather more data**: Collect additional data using tools like `tcpdump` or `Wireshark` to gather more insight into the packet drops.

## Mitigation

To mitigate the issue, follow these steps:

1. **Reduce network utilization**: Implement traffic shaping or policing to reduce network utilization and prevent packet drops.
2. **Adjust routing and forwarding**: Adjust routing and forwarding configurations to prevent packet drops.
3. **Update firmware or software**: Apply firmware or software updates to resolve any known issues.
4. **Implement quality of service (QoS)**: Configure QoS policies to prioritize critical traffic and prevent packet drops.
5. **Contact network operations team**: If the issue cannot be resolved, contact the network operations team for further assistance.

Note: The specific mitigation steps will vary depending on the root cause of the issue and the network architecture.