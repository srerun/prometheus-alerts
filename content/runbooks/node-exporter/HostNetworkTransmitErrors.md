---
title: HostNetworkTransmitErrors
description: Troubleshooting for alert HostNetworkTransmitErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNetworkTransmitErrors

Host {{ $labels.instance }} interface {{ $labels.device }} has encountered {{ printf "%.0f" $value }} transmit errors in the last two minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostNetworkTransmitErrors" %}}

{{% comment %}}

```yaml
alert: HostNetworkTransmitErrors
expr: (rate(node_network_transmit_errs_total[2m]) / rate(node_network_transmit_packets_total[2m]) > 0.01) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host Network Transmit Errors (instance {{ $labels.instance }})
    description: |-
        Host {{ $labels.instance }} interface {{ $labels.device }} has encountered {{ printf "%.0f" $value }} transmit errors in the last two minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostNetworkTransmitErrors.md

```

{{% /comment %}}

</details>


Here is a runbook for the HostNetworkTransmitErrors alert rule:

## Meaning

This alert is triggered when the rate of network transmit errors on a host exceeds 1% of the total packets transmitted over a 2-minute period. This indicates that there is an issue with the host's network interface or the network itself, leading to a significant number of packet transmission errors.

## Impact

The impact of this issue can be significant, as it can cause:

* Packet loss and retransmission, leading to increased latency and decreased network performance
* Increased CPU usage on the host as it attempts to retransmit packets
* Potential application errors or crashes due to network connectivity issues

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the host's network interface configuration and status using tools such as `ip addr show` or `ethtool`.
2. Verify that the host's network cables are securely connected and that there are no issues with the physical network infrastructure.
3. Check the system logs for any errors or warnings related to the network interface or device driver.
4. Use tools such as `tcpdump` or `Wireshark` to capture and analyze network traffic to identify any patterns or issues.
5. Check the Node Exporter metrics to see if there are any other network-related issues, such as high levels of packet drops or errors.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the network interface or the entire host to reset the network configuration and clear any temporary issues.
2. Check and update the network interface drivers to the latest version.
3. Adjust the network interface configuration to optimize performance and reduce errors.
4. Implement network Quality of Service (QoS) policies to prioritize critical network traffic and reduce congestion.
5. Consider replacing or repairing the network interface or physical network infrastructure if hardware issues are suspected.

Remember to investigate and address the root cause of the issue to prevent it from recurring.