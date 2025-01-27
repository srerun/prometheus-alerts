---
title: HostSystemdServiceCrashed
description: Troubleshooting for alert HostSystemdServiceCrashed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostSystemdServiceCrashed

systemd service crashed

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostSystemdServiceCrashed" %}}

{{% comment %}}

```yaml
alert: HostSystemdServiceCrashed
expr: (node_systemd_unit_state{state="failed"} == 1) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host systemd service crashed (instance {{ $labels.instance }})
    description: |-
        systemd service crashed
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostSystemdServiceCrashed.md

```

{{% /comment %}}

</details>


## Meaning

The `HostSystemdServiceCrashed` alert is triggered when a systemd service on a host crashes. This indicates that a critical system service has failed, which can impact the overall functionality and stability of the host. This alert is classified as a warning, indicating that immediate attention is required to prevent potential disruptions.

## Impact

* The failed systemd service may cause dependent services to also fail or become unavailable.
* The host may experience instability, slowdowns, or even become unresponsive.
* Critical system functions, such as logging, networking, or storage, may be affected.
* The availability and reliability of the host may be compromised, leading to potential outages or data loss.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the systemd service status using the `systemctl` command:
```
systemctl status <service_name>
```
Replace `<service_name>` with the name of the failed service.

2. Review the system logs for errors related to the failed service:
```
journalctl -u <service_name>
```
3. Verify the service configuration and dependencies using:
```
systemctl show <service_name>
```
4. Check the node's system resource utilization (CPU, memory, disk space) to identify potential resource constraints.

## Mitigation

To mitigate the issue, follow these steps:

1. Attempt to restart the failed systemd service:
```
systemctl restart <service_name>
```
2. If the service fails to restart, investigate the underlying cause and resolve any issues found.
3. Check for any system updates or patches that may be applicable to the failed service.
4. Consider implementing fail-safe mechanisms, such as service redundancy or automatic restart policies, to minimize the impact of future service crashes.
5. Monitor the host's system resources and adjust resource allocation as needed to prevent resource constraints.

Remember to consult the node's system documentation and configuration files for specific guidance on troubleshooting and resolving systemd service crashes.

If the service is no longer relevant on the affected node:

1. Disable the service:

```
systemctl disable <service_name>
```

2. Clear the failed state:

```
systemctl reset-failed <service_name>
```