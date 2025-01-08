---
title: HostNetworkInterfaceSaturated
description: Troubleshooting for alert HostNetworkInterfaceSaturated
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNetworkInterfaceSaturated

The network interface "{{ $labels.device }}" on "{{ $labels.instance }}" is getting overloaded.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostNetworkInterfaceSaturated" %}}

{{% comment %}}

```yaml
alert: HostNetworkInterfaceSaturated
expr: ((rate(node_network_receive_bytes_total{device!~"^tap.*|^vnet.*|^veth.*|^tun.*"}[1m]) + rate(node_network_transmit_bytes_total{device!~"^tap.*|^vnet.*|^veth.*|^tun.*"}[1m])) / node_network_speed_bytes{device!~"^tap.*|^vnet.*|^veth.*|^tun.*"} > 0.8 < 10000) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 1m
labels:
    severity: warning
annotations:
    summary: Host Network Interface Saturated (instance {{ $labels.instance }})
    description: |-
        The network interface "{{ $labels.device }}" on "{{ $labels.instance }}" is getting overloaded.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostNetworkInterfaceSaturated.md

```

{{% /comment %}}

</details>


Here is the runbook for the `HostNetworkInterfaceSaturated` alert:

## Meaning

This alert is triggered when the network interface on a host is experiencing high utilization, indicating potential network congestion and performance issues. The alert is based on the rate of received and transmitted bytes over a 1-minute period, compared to the network interface's speed.

## Impact

* Network congestion and packet loss
* Slow application performance and response times
* Increased latency and timeouts
* Potential impact on critical business applications and services

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check network interface utilization**: Verify the current network utilization using tools like `top`, `htop`, or `mpstat`.
2. **Investigate network traffic patterns**: Use tools like `tcpdump` or `Wireshark` to analyze network traffic patterns and identify potential causes of high utilization.
3. **Check system resource utilization**: Verify CPU, memory, and disk utilization to ensure they are not contributing to the high network utilization.
4. **Review application logs**: Check application logs for any errors or warnings related to network connectivity or performance issues.

## Mitigation

To mitigate the issue, follow these steps:

1. **Investigate and resolve underlying causes**: Identify and address the root cause of high network utilization, such as:
	* High traffic volumes from a specific application or service
	* Network configuration issues
	* Hardware or firmware issues with the network interface
2. **Implement traffic shaping or rate limiting**: Consider implementing traffic shaping or rate limiting to prevent overutilization of the network interface.
3. **Upgrade network infrastructure**: If necessary, consider upgrading network infrastructure to increase bandwidth and reduce congestion.
4. **Monitor and adjust**: Continuously monitor network utilization and adjust mitigation strategies as needed to ensure optimal network performance.