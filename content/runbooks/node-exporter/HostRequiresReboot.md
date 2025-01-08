---
title: HostRequiresReboot
description: Troubleshooting for alert HostRequiresReboot
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostRequiresReboot

{{ $labels.instance }} requires a reboot.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostRequiresReboot" %}}

{{% comment %}}

```yaml
alert: HostRequiresReboot
expr: (node_reboot_required > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 4h
labels:
    severity: info
annotations:
    summary: Host requires reboot (instance {{ $labels.instance }})
    description: |-
        {{ $labels.instance }} requires a reboot.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostRequiresReboot.md

```

{{% /comment %}}

</details>


Here is a runbook for the `HostRequiresReboot` alert rule:

## Meaning

The `HostRequiresReboot` alert is triggered when a node requires a reboot, as indicated by the `node_reboot_required` metric. This metric is typically set by the node exporter when it detects that the node has pending reboots or updates that require a restart.

## Impact

If this alert is not addressed, the node may become unstable or experience performance issues, potentially affecting the overall health and availability of the system. Additionally, failing to reboot the node may lead to security vulnerabilities or missed updates, which can have serious consequences.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `node_reboot_required` metric to confirm that the node indeed requires a reboot.
2. Verify that the node is not currently undergoing maintenance or other scheduled reboots.
3. Review the system logs to identify any errors or warning messages related to the reboot requirement.
4. Check for any pending updates or patches that may be causing the reboot requirement.

## Mitigation

To mitigate the issue, follow these steps:

1. Schedule a maintenance window to reboot the node as soon as possible.
2. Apply any pending updates or patches that are causing the reboot requirement.
3. Verify that the `node_reboot_required` metric returns to 0 after the reboot.
4. Monitor the node for any signs of instability or performance issues after the reboot.

Note: Before taking any action, ensure you have followed your organization's change management procedures and have obtained necessary approvals.