---
title: e5AlarmNotify
description: Troubleshooting for SNMP trap e5AlarmNotify
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-110-TRAPS-MIB::e5AlarmNotify 

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


Here is a sample runbook for the SNMP trap description:

## Meaning

The E5-110-TRAPS-MIB::e5AlarmNotify trap is generated by an E5 device to notify of an alarm condition. This trap provides detailed information about the alarm, including its type, severity, timestamp, and impact on services.

## Impact

The impact of this trap depends on the specifics of the alarm condition. Potential effects include:

* Disruption of services or network connectivity
* Equipment failure or malfunction
* Performance degradation
* Security breaches
* Notification of a critical event that requires immediate attention

## Diagnosis

To diagnose the cause of the alarm, follow these steps:

1. Review the trap details, specifically the `e5AlarmObjectClass`, `e5AlarmObjectInstance*`, and `e5AlarmType` variables, to understand the type and scope of the alarm.
2. Check the `e5AlarmSeverity` variable to determine the urgency of the alarm.
3. Investigate the `e5AlarmTimeStamp` variable to determine when the alarm was triggered.
4. Examine the `e5AlarmServiceAffecting` variable to determine if the alarm is service-affecting or not.
5. Analyze the `e5AlarmText` variable to gather more information about the alarm.
6. Verify the `e5AlarmTime` variable to ensure the timestamp is correct.

## Mitigation

To mitigate the effects of the alarm, follow these steps:

1. Notify the relevant teams and stakeholders of the alarm condition.
2. Investigate the root cause of the alarm and take corrective action.
3. Implement temporary fixes or workarounds to minimize the impact of the alarm.
4. Schedule maintenance or repairs to address the underlying issue.
5. Monitor the device and alarm conditions to ensure the issue is resolved.
6. Update documentation and procedures to prevent similar alarms in the future.
---




