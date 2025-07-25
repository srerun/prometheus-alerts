---
title: sysAlarmSvrtyChange
description: Troubleshooting for SNMP trap sysAlarmSvrtyChange
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-110-TRAPS-MIB::sysAlarmSvrtyChange 

Alarm severity is changed. 



Here is a runbook for the SNMP trap `E5-110-TRAPS-MIB::sysAlarmSvrtyChange`:

## Meaning
The `sysAlarmSvrtyChange` trap indicates that the severity of an existing alarm has changed. This trap is sent when the severity of an alarm is updated, such as when an alarm changes from a warning to a critical severity.

## Impact
The impact of this trap is that it may indicate a change in the urgency or priority of an existing issue. Depending on the new severity level, this may require immediate attention from network administrators or operators to prevent service disruptions or data loss.

## Diagnosis
To diagnose the cause of this trap, follow these steps:

1. Identify the alarm that triggered the severity change.
2. Check the alarm description and details to understand the reason for the severity change.
3. Review system logs and event history to determine what events led to the severity change.
4. Verify that the alarm is valid and not a false positive.

## Mitigation
To mitigate the impact of this trap, follow these steps:

1. Immediately investigate the alarm and take necessary actions based on the new severity level.
2. Verify that all necessary teams and stakeholders are notified of the severity change.
3. Take corrective actions to resolve the underlying issue causing the alarm.
4. Update maintenance schedules or perform preventative maintenance to prevent similar issues from occurring in the future.
5. Review and update alarm configurations to ensure that severity levels are accurately reflecting the urgency of the issue.
---




