---
title: OxidizedDeviceNoConnection
description: Troubleshooting for alert OxidizedDeviceNoConnection
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OxidizedDeviceNoConnection

Oxidized is unable to connect to device {{ $labels.name }} ({{ $labels.group }}).

<details>
  <summary>Alert Rule</summary>

{{% rule "oxidized/oxidized.yml" "OxidizedDeviceNoConnection" %}}

{{% comment %}}

```yaml
alert: OxidizedDeviceNoConnection
expr: |
    oxidized_device_status == 0
for: 30m
labels:
    severity: warning
annotations:
    summary: Oxidized cannot connect to {{ $labels.name }}
    description: |
        Oxidized is unable to connect to device {{ $labels.name }} ({{ $labels.group }}).
        Status = no connection (0)
        This has been the case for more than 30 minutes.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/oxidized/oxidizeddevicenoconnection/

```

{{% /comment %}}

</details>


## Meaning
The OxidizedDeviceNoConnection alert is triggered when the Prometheus monitoring system detects that Oxidized, a tool used for automating network device configuration backups, is unable to connect to a network device. This alert is raised when the device status remains at 0 (no connection) for more than 30 minutes. The alert provides information about the affected device, including its name and group.

## Impact
The impact of this alert is that the network device configurations are not being backed up, which can lead to potential configuration loss in case of a device failure or misconfiguration. This can result in extended downtime, increased risk of human error during the recovery process, and potential security vulnerabilities. Additionally, the lack of up-to-date configuration backups can hinder troubleshooting and auditing efforts.

## Diagnosis
To diagnose the issue, the following steps can be taken:
1. **Verify device connectivity**: Check if the device is reachable over the network and if the necessary ports are open.
2. **Check Oxidized configuration**: Review the Oxidized configuration files to ensure that the device is properly configured and that the connection settings are correct.
3. **Inspect device logs**: Examine the device logs to identify any potential issues that may be preventing the connection.
4. **Test Oxidized connection**: Attempt to manually connect to the device using Oxidized to isolate the issue.

## Mitigation
To mitigate the issue, the following steps can be taken:
1. **Restore device connectivity**: Resolve any network connectivity issues or port configuration problems to enable Oxidized to connect to the device.
2. **Update Oxidized configuration**: Correct any configuration issues or inaccuracies in the Oxidized configuration files.
3. **Restart Oxidized service**: Restart the Oxidized service to apply any configuration changes and re-establish the connection to the device.
4. **Verify backup configuration**: Confirm that the device configuration is being backed up correctly and that the backup schedule is intact.
5. **Monitor device connection**: Closely monitor the device connection to ensure that the issue does not recur and that the backup configuration remains intact.