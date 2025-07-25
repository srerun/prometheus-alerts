---
title: e5AlarmNotify
description: Troubleshooting for SNMP trap e5AlarmNotify
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-121-TRAPS-MIB::e5AlarmNotify 

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


Here is a runbook for the SNMP trap `E5-121-TRAPS-MIB::e5AlarmNotify`:

## Meaning

The `E5-121-TRAPS-MIB::e5AlarmNotify` trap indicates that an E5 device has generated an alarm notification. This alarm may be related to a specific object or component of the device, and the trap provides detailed information about the alarm, including its type, severity, and timestamp.

## Impact

The impact of this trap depends on the specific alarm being generated. If the alarm is service-affecting, it may indicate a critical issue that requires immediate attention to prevent service disruption. Even if the alarm is not service-affecting, it may still indicate a problem that needs to be addressed to prevent future issues.

## Diagnosis

To diagnose the issue, follow these steps:

1. Review the trap details to determine the type and severity of the alarm.
2. Check the `e5AlarmObjectClass` and `e5AlarmObjectInstance` variables to identify the specific object or component related to the alarm.
3. Review the `e5AlarmText` variable to obtain a brief description of the alarm.
4. Verify the `e5AlarmServiceAffecting` variable to determine if the alarm is service-affecting.
5. Check the `e5AlarmTimeStamp` and `e5AlarmTime` variables to determine when the alarm was generated.
6. Consult the E5 device documentation and any relevant logs to gather more information about the alarm.

## Mitigation

To mitigate the issue, follow these steps:

1. Depending on the severity and impact of the alarm, take immediate action to address the issue.
2. If the alarm is service-affecting, take steps to restore service as quickly as possible.
3. Perform a thorough investigation to determine the root cause of the alarm.
4. Implement corrective actions to prevent the alarm from occurring in the future.
5. Verify that the alarm has been cleared or resolved before considering the issue closed.
6. Update any relevant documentation or logs to reflect the mitigation steps taken.
---




