---
title: e5AlarmNotify
description: Troubleshooting for SNMP trap e5AlarmNotify
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-111-TRAPS-MIB::e5AlarmNotify 

E5 device alarm notification. 


## Variables


  - e5AlarmObjectClass
  - e5AlarmObjectInstance1
  - e5AlarmObjectInstance2
  - e5AlarmObjectInstance3
  - e5AlarmObjectInstance4
  - e5AlarmType
  - e5AlarmSeverity
  - e5AlarmTimeStamp
  - e5AlarmServiceAffecting
  - e5AlarmText
  - e5AlarmTime 

### Definitions 


**e5AlarmObjectClass** 
: Object Class for an alarm 

**e5AlarmObjectInstance1** 
: Object instance for an alarm, level 1 

**e5AlarmObjectInstance2** 
: Object instance for an alarm, level 2 

**e5AlarmObjectInstance3** 
: Object instance for an alarm, level 3 

**e5AlarmObjectInstance4** 
: Object instance for an alarm, level 4 

**e5AlarmType** 
: Unique type for an alarm 

**e5AlarmSeverity** 
: Severity of the alarm 

**e5AlarmTimeStamp** 
: Timestamp indicating the set/clear time of the alarm 

**e5AlarmServiceAffecting** 
: Indicated the nature of the alarm i.e. service
affecting or not 

**e5AlarmText** 
: Alarm description 

**e5AlarmTime** 
: UTC time 


Here is a runbook for the SNMP trap description:

## Meaning

The E5-111-TRAPS-MIB::e5AlarmNotify trap indicates that an alarm has been triggered on an E5 device. This trap provides detailed information about the alarm, including its type, severity, timestamp, and impact on services.

## Impact

The impact of this trap varies depending on the severity and type of alarm. However, potential impacts include:

* Service disruption or degradation
* Network availability issues
* Performance degradation
* Security vulnerabilities

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `e5AlarmType` variable to determine the specific type of alarm that has been triggered.
2. Examine the `e5AlarmSeverity` variable to determine the severity of the alarm.
3. Review the `e5AlarmText` variable for a detailed description of the alarm.
4. Check the `e5AlarmServiceAffecting` variable to determine if the alarm is service-affecting.
5. Verify the `e5AlarmTimeStamp` variable to determine when the alarm was triggered.
6. Investigate the E5 device logs and monitoring systems to gather more information about the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Take immediate action to address the alarm, depending on its severity and type.
2. If the alarm is service-affecting, prioritize restoring service as quickly as possible.
3. Perform troubleshooting and root cause analysis to determine the cause of the alarm.
4. Implement corrective actions to prevent the alarm from recurring.
5. Update monitoring systems and logs to reflect the status of the alarm.
6. Notify relevant stakeholders and teams of the alarm and its resolution.
---




