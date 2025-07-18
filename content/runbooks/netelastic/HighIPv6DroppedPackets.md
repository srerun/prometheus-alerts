---
title: HighIPv6DroppedPackets
description: Troubleshooting for alert HighIPv6DroppedPackets
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HighIPv6DroppedPackets

IPv6 dropped packets (up or down) for user {{ $labels.vbng_userName }} (index: {{ $labels.vbng_userIndex }}) exceed 100 for more than 5 minutes (current value: {{ $value }}). This may indicate network issues.

<details>
  <summary>Alert Rule</summary>

{{% rule "netelastic/netelastic.yml" "HighIPv6DroppedPackets" %}}

{{% comment %}}

```yaml
alert: HighIPv6DroppedPackets
expr: vbng_ipv6_down_dropped_packets > 100 or vbng_ipv6_up_dropped_packets > 100
for: 5m
labels:
    severity: warning
annotations:
    summary: High IPv6 dropped packets
    description: 'IPv6 dropped packets (up or down) for user {{ $labels.vbng_userName }} (index: {{ $labels.vbng_userIndex }}) exceed 100 for more than 5 minutes (current value: {{ $value }}). This may indicate network issues.'
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/netelastic/highipv6droppedpackets/

```

{{% /comment %}}

</details>


Here is the runbook for the HighIPv6DroppedPackets alert:

## Meaning

The HighIPv6DroppedPackets alert indicates that the number of IPv6 packets dropped on the user's connection (either upstream or downstream) has exceeded 100 packets for more than 5 minutes. This may indicate network issues that need to be investigated and resolved.

## Impact

The impact of this alert is that the user may experience poor network performance, including slow data transfer rates, packet loss, and connectivity issues. This can lead to a poor user experience and may affect the user's ability to use critical applications.

## Diagnosis

To diagnose the root cause of this issue, follow these steps:

* Check the user's network configuration to ensure that it is properly configured for IPv6.
* Verify that there are no issues with the user's equipment or cables that could be causing packet loss.
* Check the network traffic and packet capture data to identify the source of the dropped packets.
* Investigate potential network congestion or bottleneck issues.
* Check the system logs for any error messages related to network connectivity or packet processing.

## Mitigation

To mitigate the HighIPv6DroppedPackets alert, follow these steps:

* Check the user's network configuration and make any necessary adjustments to ensure that it is properly configured for IPv6.
* Work with the user to troubleshoot and resolve any issues with their equipment or cables.
* Implement Quality of Service (QoS) policies to prioritize critical applications and reduce network congestion.
* Increase the buffer size or adjust the packet processing configuration to reduce packet loss.
* Consider implementing traffic shaping or policing to prevent network congestion.
* If the issue persists, consider escalating the issue to the network operations team for further investigation and resolution.