---
title: sysAlarmSvrtyChange
description: Troubleshooting for SNMP trap sysAlarmSvrtyChange
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-121-TRAPS-MIB::sysAlarmSvrtyChange 

Alarm severity is changed. 



## Meaning

The `E5-121-TRAPS-MIB::sysAlarmSvrtyChange` SNMP trap indicates that the severity of an existing alarm has changed. This trap is generated when the device detects a change in the severity of an alarm, which can affect the overall system's performance and operation. The severity of an alarm can range from minor to critical, and this trap is used to notify the network management system of the change.

## Impact

The impact of this trap can vary depending on the specific alarm and its new severity level. However, in general, a change in alarm severity can have the following effects:

* Increased or decreased urgency for technical support teams to respond to the alarm
* Changes to the prioritization of alarms and troubleshooting efforts
* Potential disruption to normal system operations or service delivery
* In some cases, a change in alarm severity can indicate a worsening or improving condition that requires immediate attention

## Diagnosis

To diagnose the cause of the `E5-121-TRAPS-MIB::sysAlarmSvrtyChange` trap, follow these steps:

1. Identify the specific alarm that triggered the trap by checking the alarm ID and description.
2. Review the alarm's history to determine the previous severity level and when the change occurred.
3. Check the system logs and event history to identify any related events or errors that may have contributed to the severity change.
4. Verify the current system configuration and settings to ensure they are correct and up-to-date.
5. Consult the device's documentation and technical support resources for guidance on troubleshooting the specific alarm and severity change.

## Mitigation

To mitigate the effects of the `E5-121-TRAPS-MIB::sysAlarmSvrtyChange` trap, follow these steps:

1. Assess the new severity level of the alarm and prioritize response efforts accordingly.
2. Take immediate action to address the alarm, if necessary, to prevent further system disruption or service degradation.
3. Update the network management system and technical support teams on the alarm's new severity level and any required response efforts.
4. Perform a thorough root cause analysis to identify the underlying cause of the severity change and implement corrective actions to prevent similar events in the future.
5. Verify that the system is operating within normal parameters and that the alarm's new severity level is accurate and reflect the current system state.
---




