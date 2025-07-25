---
title: aeTrapAlarm
description: Troubleshooting for SNMP trap aeTrapAlarm
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# AE-ALARM-TABLE-MIB::aeTrapAlarm 

aeTrapAlarm is generated whenever an alarm is raised or cleared. 


## Variables


  - aeTrapSequenceNo
  - aeTrapEquipmentType
  - aeTrapEquipmentInstance
  - aeTrapAlarmType
  - aeTrapAlarmStatus
  - aeTrapAlarmSeverityLevel
  - aeTrapServiceAffecting
  - aeTrapText
  - aeTrapTimeStamp
  - aeTrapTime
  - aeTrapFsanSerialNumber
  - aeTrapRegistrationID 

### Definitions 


**aeTrapSequenceNo** 
: Uniquely identifies each alarm trap that is transmitted by the ONT.
The value Increment for each alarm trap that is transmitted.
The first trap has a sequence number of one (1). 

**aeTrapEquipmentType** 
: The type of physical equipment the ONT Alarm is associated with. 

**aeTrapEquipmentInstance** 
: The instance of the physical equipment the ONT Alarm is assocaited with. 

**aeTrapAlarmType** 
: The type of the ONT Alarm. 

**aeTrapAlarmStatus** 
: The status of the ONT alarm - on/off. 

**aeTrapAlarmSeverityLevel** 
: The severity level of the ONT Alarm. 

**aeTrapServiceAffecting** 
: This value indicated whether the ONT Alarm is service affecting or not. 

**aeTrapText** 
: This object contains the brief description about the ONT Alarm. 

**aeTrapTimeStamp** 
: Local time string for the ONT Alarm. 

**aeTrapTime** 
: UTC time integer of the ONT Alarm. 

**aeTrapFsanSerialNumber** 
: The FSAN Serial Number of the ONT expressed as 4 charaters and 8 hex digits. 

**aeTrapRegistrationID** 
: The Registration ID of the ONT expressed as max 10 char numerical string. 


Here is a runbook for the SNMP trap AE-ALARM-TABLE-MIB::aeTrapAlarm:

## Meaning

The AE-ALARM-TABLE-MIB::aeTrapAlarm trap is generated whenever an alarm is raised or cleared on an Optical Network Terminal (ONT). This trap provides detailed information about the alarm, including the type, severity, and status.

## Impact

The impact of this trap depends on the specific alarm that is being raised or cleared. If the alarm is service-affecting, it may indicate a disruption to network services, which could affect users and customers. In some cases, the alarm may be a warning or notification of a potential issue, rather than an immediate problem. In any case, it is essential to investigate and diagnose the cause of the alarm to prevent or minimize service disruptions.

## Diagnosis

To diagnose the cause of the AE-ALARM-TABLE-MIB::aeTrapAlarm trap, follow these steps:

1. Review the trap variables to understand the specifics of the alarm:
	* aeTrapAlarmType: Determine the type of alarm that has been raised or cleared.
	* aeTrapAlarmStatus: Check if the alarm is active (on) or cleared (off).
	* aeTrapAlarmSeverityLevel: Assess the severity of the alarm to prioritize the response.
	* aeTrapServiceAffecting: Determine if the alarm is service-affecting or not.
2. Investigate the physical equipment associated with the alarm:
	* aeTrapEquipmentType: Identify the type of physical equipment involved.
	* aeTrapEquipmentInstance: Determine the specific instance of the equipment.
3. Check the alarm text and timestamp:
	* aeTrapText: Review the brief description of the alarm.
	* aeTrapTimeStamp: Note the local time of the alarm.
	* aeTrapTime: Check the UTC time of the alarm.

## Mitigation

To mitigate the impact of the AE-ALARM-TABLE-MIB::aeTrapAlarm trap, follow these steps:

1. Investigate the cause of the alarm: Based on the diagnosis, determine the root cause of the alarm and take corrective action.
2. Clear the alarm: If the alarm is no longer valid, clear it to prevent further notifications.
3. Update the equipment: If the alarm is related to a hardware or software issue, update or replace the equipment as necessary.
4. Notify stakeholders: Inform relevant stakeholders of the alarm and the actions taken to resolve it.
5. Document the incident: Log the incident and the actions taken to resolve it for future reference and improvement.
---




