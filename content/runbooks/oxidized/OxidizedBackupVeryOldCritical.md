---
title: OxidizedBackupVeryOldCritical
description: Troubleshooting for alert OxidizedBackupVeryOldCritical
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OxidizedBackupVeryOldCritical

Device {{ $labels.name }} has not been backed up in over 7 days!

<details>
  <summary>Alert Rule</summary>

{{% rule "oxidized/oxidized.yml" "OxidizedBackupVeryOldCritical" %}}

{{% comment %}}

```yaml
alert: OxidizedBackupVeryOldCritical
expr: |
    time() - oxidized_device_last_backup_end > 86400 * 7
    and on(name) oxidized_device_status == 2
for: 60m
labels:
    severity: critical
annotations:
    summary: 'CRITICAL: No recent backup for {{ $labels.name }}'
    description: |
        Device {{ $labels.name }} has not been backed up in over 7 days!
        Last backup: {{ $value | humanizeDuration }} ago
        Immediate attention required.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/oxidized/oxidizedbackupveryoldcritical/

```

{{% /comment %}}

</details>


## Meaning
The OxidizedBackupVeryOldCritical alert is triggered when a device has not been backed up in over 7 days. Oxidized is a tool used for backing up network device configurations. This alert indicates that the backup process for a specific device has failed or been skipped for an extended period, which could lead to potential configuration loss in case of a device failure or misconfiguration.

## Impact
The impact of this alert is significant, as it indicates that the configuration of a critical network device has not been backed up recently. If the device were to fail or become misconfigured, the lack of a recent backup could lead to extended downtime and potential data loss. This could have serious consequences for business operations, network reliability, and security.

## Diagnosis
To diagnose the issue, follow these steps:
1. **Check Oxidized logs**: Review the Oxidized logs to determine the cause of the backup failure. Common issues include connectivity problems, authentication failures, or device configuration errors.
2. **Verify device connectivity**: Ensure that the device is reachable and responding to Oxidized requests.
3. **Confirm device configuration**: Verify that the device configuration has not changed in a way that would prevent backups from occurring.
4. **Check Oxidized configuration**: Review the Oxidized configuration to ensure that it is correctly set up to back up the device.

## Mitigation
To mitigate the issue, follow these steps:
1. **Manually trigger a backup**: Use Oxidized to manually trigger a backup of the device to ensure that the configuration is captured.
2. **Investigate and resolve the root cause**: Identify the root cause of the backup failure and take corrective action to prevent future failures.
3. **Schedule regular backups**: Ensure that regular backups are scheduled to prevent this issue from occurring in the future.
4. **Implement monitoring and alerting**: Implement additional monitoring and alerting to detect backup failures and notify teams in a timely manner.
5. **Review and update runbooks**: Review and update runbooks to ensure that they are current and effective in resolving backup-related issues.