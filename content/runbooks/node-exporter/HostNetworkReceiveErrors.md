---
title: HostNetworkReceiveErrors
description: Troubleshooting for alert HostNetworkReceiveErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostNetworkReceiveErrors

Host {{ $labels.instance }} interface {{ $labels.device }} has encountered {{ printf "%.0f" $value }} receive errors in the last two minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostNetworkReceiveErrors" %}}

{{% comment %}}

```yaml
alert: HostNetworkReceiveErrors
expr: (rate(node_network_receive_errs_total[2m]) / rate(node_network_receive_packets_total[2m]) > 0.01) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 2m
labels:
    severity: warning
annotations:
    summary: Host Network Receive Errors (instance {{ $labels.instance }})
    description: |-
        Host {{ $labels.instance }} interface {{ $labels.device }} has encountered {{ printf "%.0f" $value }} receive errors in the last two minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostNetworkReceiveErrors.md

```

{{% /comment %}}

</details>


## Meaning

The `HostNetworkReceiveErrors` alert is triggered when the rate of receive errors on a host's network interface exceeds 1% of the total receive packets over a 2-minute period. This indicates a potential issue with the host's network connectivity or configuration.

## Impact

* Receive errors can cause packet loss, leading to decreased network performance and potential application failures.
* High receive error rates can indicate issues with the network interface, cable, or switch, which can impact the entire cluster or application.
* Unaddressed receive errors can lead to increased latency, reduced throughput, and decreased overall system reliability.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected host and interface using the `instance` and `device` labels.
2. Check the system logs for any error messages related to the network interface or driver.
3. Verify the network cable and connection to ensure they are secure and functioning properly.
4. Use tools like `ethtool` or `ip link` to check the interface configuration and stats.
5. Check for any firmware or driver updates for the network interface.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the network interface using `ip link set <interface> down && ip link set <interface> up`.
2. Check and update the network interface firmware or driver to the latest version.
3. Inspect and clean the network cable and connection to ensure they are secure and functioning properly.
4. Consider replacing the network cable or interface if it's faulty.
5. If the issue persists, consider escalating to a network administrator or infrastructure team for further assistance.

Remember to reference the node exporter documentation and official runbooks for more detailed troubleshooting and mitigation steps.