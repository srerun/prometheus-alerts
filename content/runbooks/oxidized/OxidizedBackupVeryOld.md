---
title: OxidizedBackupVeryOld
description: Troubleshooting for alert OxidizedBackupVeryOld
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OxidizedBackupVeryOld

Last successful backup of {{ $labels.name }} is older than 2 days.

<details>
  <summary>Alert Rule</summary>

{{% rule "oxidized/oxidized.yml" "OxidizedBackupVeryOld" %}}

{{% comment %}}

```yaml
alert: OxidizedBackupVeryOld
expr: |
    time() - oxidized_device_last_backup_end > 86400 * 2
    and on(name) oxidized_device_status == 2
    and on(name) oxidized_device_last_backup_status == 2
for: 60m
labels:
    severity: warning
annotations:
    summary: Very old backup for {{ $labels.name }}
    description: |
        Last successful backup of {{ $labels.name }} is older than 2 days.
        Age: {{ $value | humanizeDuration }}
        Device is still marked as success, but backup is stale.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/oxidized/oxidizedbackupveryold/

```

{{% /comment %}}

</details>


## Meaning
The OxidizedBackupVeryOld alert is triggered when the last successful backup of a device using Oxidized is older than 2 days. Oxidized is a tool used for backing up network device configurations. This alert indicates that the device's configuration has not been backed up recently, which could lead to potential issues in case of a device failure or configuration loss.

## Impact
The impact of this alert is that the device's configuration is not being backed up regularly, which could result in:
* Loss of configuration data in case of a device failure
* Increased downtime and recovery time in case of a device failure
* Potential security risks if the device's configuration is not properly backed up and restored
* Compliance issues if regulatory requirements for configuration backups are not met

## Diagnosis
To diagnose the issue, the following steps can be taken:
* Check the Oxidized logs to see if there are any errors or issues that are preventing the backups from running successfully
* Verify that the device is configured correctly in Oxidized and that the backup job is scheduled to run regularly
* Check the device's configuration to ensure that it is not blocking the backup process
* Review the Prometheus metrics for the device to see if there are any other issues that could be related to the backup failure

## Mitigation
To mitigate the issue, the following steps can be taken:
* Check the Oxidized configuration and ensure that the backup job is scheduled to run regularly
* Run a manual backup of the device to ensure that the backup process is working correctly
* Verify that the device's configuration is not blocking the backup process
* Consider increasing the frequency of backups or implementing a more robust backup solution to prevent similar issues in the future
* Update the device's status in Oxidized to reflect the current backup status and ensure that the device is properly configured for backups.