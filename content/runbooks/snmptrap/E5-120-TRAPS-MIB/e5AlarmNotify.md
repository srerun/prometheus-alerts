---
title: e5AlarmNotify
description: Troubleshooting for SNMP trap e5AlarmNotify
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-120-TRAPS-MIB::e5AlarmNotify 

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


## Meaning

The E5-120-TRAPS-MIB::e5AlarmNotify trap is generated by an E5 device to notify of an alarm condition. This trap provides detailed information about the alarm, including its type, severity, timestamp, and impact on services.

## Impact

The impact of this trap is significant, as it indicates a problem with the E5 device that requires immediate attention. The alarm may be service-affecting, which means it could be causing disruptions to critical services or impacting network performance. Ignoring this trap could lead to prolonged outages, data loss, or security breaches.

## Diagnosis

To diagnose the root cause of the alarm, follow these steps:

1. Check the `e5AlarmType` variable to determine the specific type of alarm.
2. Review the `e5AlarmSeverity` variable to understand the severity of the alarm.
3. Examine the `e5AlarmText` variable for a detailed description of the alarm.
4. Check the `e5AlarmServiceAffecting` variable to determine if the alarm is service-affecting.
5. Verify the `e5AlarmTimeStamp` variable to understand when the alarm was triggered.
6. Analyze the `e5AlarmObjectClass` and `e5AlarmObjectInstance` variables to identify the specific object or instance associated with the alarm.

## Mitigation

To mitigate the impact of the alarm, follow these steps:

1. Immediately notify the E5 device administrators and technical support teams of the alarm.
2. Isolate the affected device or service to prevent further disruption.
3. Gather detailed logs and diagnostic data from the E5 device to aid in troubleshooting.
4. Perform a thorough analysis of the alarm data to identify the root cause of the problem.
5. Implement corrective actions to resolve the underlying issue, such as software updates, configuration changes, or hardware replacements.
6. Verify that the alarm has been cleared and services have been restored to normal.
---




