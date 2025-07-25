---
title: sysAlarmSvrtyChange
description: Troubleshooting for SNMP trap sysAlarmSvrtyChange
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-111-TRAPS-MIB::sysAlarmSvrtyChange 

Alarm severity is changed. 



Here is a sample runbook for the given SNMP trap description:

## Meaning
The sysAlarmSvrtyChange trap is generated when the severity of an alarm on the system changes. This could be an upgrade or downgrade of an existing alarm, or a change in the criticality of an issue. This trap is typically used to notify administrators of changes to the alarm state, allowing them to take necessary actions to investigate and resolve the issue.

## Impact
The impact of this trap depends on the specific alarm that has changed severity. If the alarm has increased in severity, it may indicate a critical issue that requires immediate attention. If the alarm has decreased in severity, it may indicate that a problem is resolving itself or has been resolved. In either case, failure to investigate and respond to the change in alarm severity could result in system downtime, data loss, or other negative consequences.

## Diagnosis
To diagnose the cause of this trap, follow these steps:

1. Identify the specific alarm that has changed severity using the trap information or system logs.
2. Investigate the cause of the alarm using system logs, performance metrics, and other monitoring tools.
3. Determine the severity of the alarm and its impact on system operations.
4. Check for any recent changes to the system or its configuration that may have contributed to the alarm severity change.

## Mitigation
To mitigate the effects of this trap, follow these steps:

1. Immediately investigate the cause of the alarm severity change and take necessary actions to resolve the underlying issue.
2. Notify relevant teams or stakeholders of the alarm severity change and its potential impact on system operations.
3. Update system logs and documentation to reflect the change in alarm severity and the actions taken to resolve the issue.
4. Consider implementing additional monitoring or logging to detect similar issues in the future and prevent them from occurring.
---




