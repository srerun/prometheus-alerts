---
title: HostUnusualNetworkThroughputOut
description: Troubleshooting for alert HostUnusualNetworkThroughputOut
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualNetworkThroughputOut

Host network interfaces are probably sending too much data (> 100 MB/s)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostUnusualNetworkThroughputOut" %}}

{{% comment %}}

```yaml
alert: HostUnusualNetworkThroughputOut
expr: (sum by (instance) (rate(node_network_transmit_bytes_total[2m])) / 1024 / 1024 > 100) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 5m
labels:
    severity: warning
annotations:
    summary: Host unusual network throughput out (instance {{ $labels.instance }})
    description: |-
        Host network interfaces are probably sending too much data (> 100 MB/s)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostUnusualNetworkThroughputOut.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "HostUnusualNetworkThroughputOut":

## Meaning

The "HostUnusualNetworkThroughputOut" alert is triggered when a host's network interface is transmitting data at an unusually high rate, exceeding 100 MB/s. This could indicate a potential issue with the host or its network configuration.

## Impact

A high network throughput can have several consequences, including:

* Network congestion: High network traffic can lead to congestion, causing delays and packet loss.
* Increased CPU usage: Handling large amounts of network traffic can increase CPU usage, potentially impacting system performance.
* Security risks: Unusual network activity can be a sign of malicious activity, such as data exfiltration or denial-of-service attacks.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected host: Check the `instance` label in the alert to determine which host is experiencing the unusual network throughput.
2. Check network interface usage: Use tools like `top` or `htop` to monitor network interface usage on the affected host.
3. Investigate processes using high network bandwidth: Use tools like `netstat` or `ss` to identify processes that are using high amounts of network bandwidth.
4. Review system logs: Check system logs for any errors or signs of malicious activity.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and stop any malicious processes: If you identify any malicious processes, stop them immediately.
2. Optimize system configuration: Review system configuration to ensure it is optimized for network performance.
3. Implement network traffic filtering: Consider implementing network traffic filtering rules to block or rate-limit suspicious traffic.
4. Monitor network activity: Continuously monitor network activity to ensure the issue is resolved and to detect any future anomalies.

Remember to update the runbook with specific steps and tools relevant to your environment.