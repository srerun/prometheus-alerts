---
title: HostUnusualNetworkThroughputIn
description: Troubleshooting for alert HostUnusualNetworkThroughputIn
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostUnusualNetworkThroughputIn

Host network interfaces are probably receiving too much data (> 100 MB/s)

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostUnusualNetworkThroughputIn" %}}

{{% comment %}}

```yaml
alert: HostUnusualNetworkThroughputIn
expr: (sum by (instance) (rate(node_network_receive_bytes_total[2m])) / 1024 / 1024 > 100) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 5m
labels:
    severity: warning
annotations:
    summary: Host unusual network throughput in (instance {{ $labels.instance }})
    description: |-
        Host network interfaces are probably receiving too much data (> 100 MB/s)
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostUnusualNetworkThroughputIn.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

This alert is triggered when a host's network interfaces are receiving an unusually high amount of data, exceeding 100 MB/s over a 2-minute interval. This could indicate a potential issue with the host's network configuration, a denial-of-service (DoS) attack, or a misbehaving application.

## Impact

If left unaddressed, this issue could lead to:

* Network congestion, causing packet loss and slow network performance
* Increased CPU usage on the host, potentially leading to performance degradation
* Potential security risks if the high network traffic is related to malicious activity
* Downtime or instability of critical services running on the host

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the host experiencing high network throughput using the `instance` label in the alert.
2. Check the host's network configuration to ensure it is properly configured and not experiencing any hardware issues.
3. Investigate the applications running on the host to determine if any are causing the high network traffic.
4. Review system logs to identify any error messages or signs of malicious activity.
5. Use network monitoring tools to further analyze the network traffic patterns and identify potential sources of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and address any misconfigured network settings or hardware issues on the host.
2. Identify and stop any malicious applications or services causing the high network traffic.
3. Implement rate limiting or traffic shaping measures to prevent network congestion.
4. Consider increasing the host's network bandwidth or upgrading its network hardware if necessary.
5. Continuously monitor the host's network traffic to ensure the issue is resolved and does not recur.

Remember to update the alert annotations to reflect the resolving state once the issue is mitigated.