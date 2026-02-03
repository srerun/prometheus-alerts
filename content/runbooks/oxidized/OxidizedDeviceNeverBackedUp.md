---
title: OxidizedDeviceNeverBackedUp
description: Troubleshooting for alert OxidizedDeviceNeverBackedUp
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OxidizedDeviceNeverBackedUp

Oxidized has never successfully backed up device {{ $labels.name }}.

<details>
  <summary>Alert Rule</summary>

{{% rule "oxidized/oxidized.yml" "OxidizedDeviceNeverBackedUp" %}}

{{% comment %}}

```yaml
alert: OxidizedDeviceNeverBackedUp
expr: |
    oxidized_device_status == 1
    and on(name) (oxidized_device_last_backup_end == 0 or absent(oxidized_device_last_backup_end))
for: 24h
labels:
    severity: warning
annotations:
    summary: Device {{ $labels.name }} has never been backed up
    description: |
        Oxidized has never successfully backed up device {{ $labels.name }}.
        Status = never backed up (1)
        This device may be newly added or permanently failing.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/oxidized/oxidizeddeviceneverbackedup/

```

{{% /comment %}}

</details>


## Meaning
The OxidizedDeviceNeverBackedUp alert is triggered when a device has never been successfully backed up by Oxidized. This is indicated by the `oxidized_device_status` being equal to 1 and the `oxidized_device_last_backup_end` being either 0 or absent. This alert is raised after 24 hours of the device being in this state, suggesting that the device may be newly added or is permanently failing to back up.

## Impact
The impact of this alert is that the configuration of the device is not being backed up, which can lead to loss of configuration data in case the device fails or is replaced. This can result in downtime and additional effort to rebuild the configuration. Additionally, this can also lead to compliance issues if configuration backups are required by regulatory or organizational policies.

## Diagnosis
To diagnose the issue, the following steps can be taken:
1. Check the Oxidized logs for errors related to the device.
2. Verify that the device is reachable and responding to Oxidized.
3. Check the device configuration to ensure that it is correctly configured for backup.
4. Verify that the Oxidized service is running and configured correctly.
5. Check the `oxidized_device_status` and `oxidized_device_last_backup_end` metrics for the device to confirm the alert condition.

## Mitigation
To mitigate the issue, the following steps can be taken:
1. Check the device configuration and ensure that it is correctly configured for backup.
2. Verify that the Oxidized service is running and configured correctly.
3. Restart the Oxidized service if it is not running.
4. Check the Oxidized logs for errors and resolve any issues found.
5. Manually trigger a backup of the device using Oxidized to verify that the backup process is working correctly.
6. If the issue persists, check the device for any configuration issues that may be preventing the backup.
7. Consider adding additional logging or monitoring to detect and alert on backup failures in the future.