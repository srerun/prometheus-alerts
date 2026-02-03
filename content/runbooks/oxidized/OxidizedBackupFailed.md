---
title: OxidizedBackupFailed
description: Troubleshooting for alert OxidizedBackupFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OxidizedBackupFailed

Last backup attempt for device {{ $labels.name }} failed.

<details>
  <summary>Alert Rule</summary>

{{% rule "oxidized/oxidized.yml" "OxidizedBackupFailed" %}}

{{% comment %}}

```yaml
alert: OxidizedBackupFailed
expr: |
    oxidized_device_last_backup_status == 1
for: 15m
labels:
    severity: warning
annotations:
    summary: Oxidized backup failed for {{ $labels.name }} ({{ $labels.group }} / {{ $labels.model }})
    description: |
        Last backup attempt for device {{ $labels.name }} failed.
        Status = error (1)
        Last successful backup was {{ $value | humanizeDuration }} ago.
        Please check Oxidized logs for this device.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/oxidized/oxidizedbackupfailed/

```

{{% /comment %}}

</details>


## Meaning
The OxidizedBackupFailed alert is triggered when the last backup attempt for a device using Oxidized fails. Oxidized is a tool used for backing up network device configurations. The alert is raised when the `oxidized_device_last_backup_status` equals 1, indicating an error occurred during the backup process. This alert is critical as it may lead to configuration loss in case of device failure or replacement.

## Impact
The impact of this alert is significant, as a failed backup may result in the loss of the current device configuration. If a device fails or needs to be replaced, the lack of a valid backup can lead to extended downtime and additional workload to recreate the configuration from scratch or from outdated backups. This can also lead to potential security risks if the new configuration does not adhere to the current security policies.

## Diagnosis
To diagnose the issue, follow these steps:
1. **Check Oxidized Logs**: Inspect the Oxidized logs for the specific device to understand the cause of the backup failure. Common issues could include authentication problems, network connectivity issues, or device configuration errors that prevented the backup.
2. **Verify Device Connectivity**: Ensure that the network device is reachable and responding to management requests. This can be done through simple network management protocols (SNMP) checks or by attempting to log in to the device.
3. **Authentication Check**: Validate the credentials used by Oxidized to connect to the device. Incorrect or expired credentials can cause backup failures.
4. **Device Configuration Review**: Sometimes, changes in the device configuration can prevent backups. Review any recent configuration changes that might interfere with the backup process.

## Mitigation
To mitigate the OxidizedBackupFailed alert, perform the following steps:
1. **Resolve Underlying Issues**: Based on the diagnosis, address the root cause of the backup failure. This might involve fixing network connectivity issues, updating credentials, or adjusting device configurations to be compatible with Oxidized.
2. **Retry Backup**: Once the underlying issue is resolved, manually trigger a backup for the affected device to ensure the configuration is successfully captured.
3. **Monitor Future Backups**: Closely monitor subsequent backup attempts for the device to ensure the issue is fully resolved and does not recur.
4. **Review and Adjust Oxidized Configuration**: If necessary, review Oxidized's configuration and adjust settings such as retry policies, timeouts, or logging levels to better handle similar situations in the future or to improve diagnostics.