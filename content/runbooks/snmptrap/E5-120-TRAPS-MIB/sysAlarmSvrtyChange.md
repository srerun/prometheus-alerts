---
title: sysAlarmSvrtyChange
description: Troubleshooting for SNMP trap sysAlarmSvrtyChange
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-120-TRAPS-MIB::sysAlarmSvrtyChange 

Alarm severity is changed. 



## Meaning

The `E5-120-TRAPS-MIB::sysAlarmSvrtyChange` SNMP trap indicates that the severity of an alarm has changed. This trap is generated when the severity of an existing alarm is updated, and it serves as a notification to monitoring systems and administrators that the alarm's severity has been modified.

## Impact

The impact of this trap is moderate to high, depending on the specific alarm and the systems it affects. If the severity of an alarm is changed, it may indicate a change in the criticality of the underlying issue. For example, if the severity of an alarm is escalated from warning to critical, it may require immediate attention from administrators to prevent system downtime or data loss.

## Diagnosis

To diagnose the cause of this trap, follow these steps:

1. Identify the specific alarm that triggered the trap: Review the SNMP trap message to determine the OID and description of the alarm that triggered the trap.
2. Check the alarm severity: Verify the current severity of the alarm and compare it to its previous severity to determine the nature of the change.
3. Investigate the underlying cause: Analyze system logs and monitoring data to determine the root cause of the alarm severity change.
4. Check for system configuration changes: Verify that no recent configuration changes have been made to the system that could have triggered the alarm severity change.

## Mitigation

To mitigate the impact of this trap, follow these steps:

1. Acknowledge and investigate the alarm: Immediately acknowledge the alarm and investigate the underlying cause to prevent further escalation.
2. Update monitoring systems: Update monitoring systems to reflect the new severity of the alarm to ensure accurate reporting and alerting.
3. Take corrective action: Based on the diagnosis, take corrective action to address the underlying issue and prevent further alarm escalations.
4. Review and update system configuration: Review system configuration to ensure it is aligned with the current alarm severity and make updates as necessary to prevent future alarm escalations.
---




