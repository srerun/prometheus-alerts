---
title: ProxmoxNodeTargetLost
description: Troubleshooting for alert ProxmoxNodeTargetLost
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProxmoxNodeTargetLost

Something wrong with the exporter, the Proxmox API server(s) it is configured to make requests to, or the server the exporter is running on

<details>
  <summary>Alert Rule</summary>

{{% rule "proxmox/proxmox-exporter.yml" "ProxmoxNodeTargetLost" %}}

{{% comment %}}

```yaml
alert: ProxmoxNodeTargetLost
expr: |
    absent_over_time(proxmox_node_up[1h])
for: 1m
labels:
    severity: critical
annotations:
    summary: Proxmox node up metric absent for {{ printf "{{ $labels.node }}" }}
    description: Something wrong with the exporter, the Proxmox API server(s) it is configured to make requests to, or the server the exporter is running on
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/proxmox-exporter/proxmoxnodetargetlost/

```

{{% /comment %}}

</details>


## Meaning

The ProxmoxNodeTargetLost alert is triggered when the `proxmox_node_up` metric is absent for more than 1 hour. This metric is expected to be present and updated regularly by the Proxmox exporter, which collects data from the Proxmox API server(s). The absence of this metric indicates a problem with the exporter, the Proxmox API server(s), or the server the exporter is running on.

## Impact

The impact of this alert is high, as it indicates a loss of visibility into the Proxmox node's status. Without this metric, it is not possible to monitor the node's uptime, which can lead to:

* Un detected downtime or performance issues
* Inability to perform automated failovers or take corrective actions
* Increased Mean Time To Detect (MTTD) and Mean Time To Resolve (MTTR) for issues affecting the Proxmox node

## Diagnosis

To diagnose the cause of the alert, follow these steps:

1. Check the Proxmox exporter logs for errors or warnings related to connectivity issues with the Proxmox API server(s)
2. Verify that the Proxmox API server(s) are reachable and responding correctly
3. Check the server running the Proxmox exporter for any issues, such as high CPU or memory usage, disk space issues, or network connectivity problems
4. Verify that the `proxmox_node_up` metric is being collected and reported correctly by the exporter
5. Check the Prometheus configuration and rules to ensure that the alert is not being suppressed or misconfigured

## Mitigation

To mitigate the alert, follow these steps:

1. Investigate and resolve any connectivity issues between the Proxmox exporter and the Proxmox API server(s)
2. Restart the Proxmox exporter service or the server it is running on, if necessary
3. Verify that the `proxmox_node_up` metric is being collected and reported correctly by the exporter
4. Consider implementing additional monitoring and logging to detect issues with the Proxmox exporter or API server(s) earlier
5. Consider implementing automated failovers or corrective actions to minimize the impact of Proxmox node downtime