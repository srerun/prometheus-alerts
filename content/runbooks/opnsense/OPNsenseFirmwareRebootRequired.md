---
title: OPNsenseFirmwareRebootRequired
description: Troubleshooting for alert OPNsenseFirmwareRebootRequired
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseFirmwareRebootRequired

OPNsense requires a reboot after firmware update

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseFirmwareRebootRequired" %}}

{{% comment %}}

```yaml
alert: OPNsenseFirmwareRebootRequired
expr: opnsense_firmware_needs_reboot == 1
for: 10m
labels:
    severity: info
annotations:
    summary: OPNsense firmware reboot required (instance {{ $labels.opnsense_instance }})
    description: |-
        OPNsense requires a reboot after firmware update
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsensefirmwarerebootrequired/

```

{{% /comment %}}

</details>


## Meaning
The `OPNsenseFirmwareRebootRequired` alert is triggered when the OPNsense firmware indicates that a reboot is required, typically after a firmware update. This alert is raised to ensure that the system is properly restarted to apply the updated firmware and maintain the stability and security of the OPNsense installation.

## Impact
The impact of not addressing this alert is that the OPNsense system may not be running with the latest firmware updates, potentially leaving it vulnerable to known security issues or missing out on newer features and improvements. Failure to reboot the system as required can lead to unpredictable behavior, reduced performance, or even complete system failure. It is crucial to perform the reboot as soon as possible to ensure the system's integrity and security.

## Diagnosis
To diagnose the issue, follow these steps:
1. **Verify the Alert Details**: Check the alert description and labels for specific information about the OPNsense instance that requires a reboot.
2. **Login to OPNsense**: Access the OPNsense web interface to review the system's current status and firmware version.
3. **Check for Firmware Updates**: Look for any pending firmware updates or notifications that a reboot is required.
4. **System Logs Review**: Examine the system logs for any error messages or indicators of system instability that could be related to the firmware update.

## Mitigation
To mitigate the issue and resolve the alert:
1. **Plan Maintenance Window**: Schedule a maintenance window to minimize the impact of the reboot on ongoing operations.
2. **Backup Configuration (Optional)**: Although not always necessary, consider backing up your current OPNsense configuration before proceeding, especially if you have made recent changes.
3. **Reboot OPNsense**: Perform the reboot from the OPNsense web interface or through the command line, depending on your access and preference.
4. **Verify System Stability**: After the reboot, verify that the OPNsense system is stable and functioning correctly, including checking for any errors in the system logs.
5. **Confirm Alert Resolution**: Once the system is confirmed to be stable and functioning as expected, verify that the alert is cleared in your monitoring system.