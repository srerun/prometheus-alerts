---
title: Altiplano acAlarm
description: Troubleshooting for Altiplano SNMP trap acAlarm
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# AC-SNMP-ALARM-MIB::acAlarm 

Altiplano sends a single trap type, `acAlarm` for multiple issues.  To determine the underlying problem the varibles must be consulted, especially `alarmType`, `alarmStatus`, and `alarmText`.


## Variables


  - notificationId
  - raisedTime
  - clearedTime
  - objectId
  - alarmNotificationOrigin
  - lastStatusChangeTime
  - alarmResource
  - durationOpen
  - deviceRefId
  - serviceAffecting
  - alarmType
  - alarmSeverity
  - alarmStatus
  - alarmText
  - neIpAddress
  - tl1Cause
  - eventType
  - proposedRepairAction
  - alarmTypeQualifier
  - customField1
  - customField2
  - customField3 

### Definitions 


**notificationId** 
: This is the notification id for an alarm. 

**raisedTime** 
: This object defines the time at which the alarm is raised. 

**clearedTime** 
: This object defines the time at which the alarm is cleared. 

**objectId** 
: This object defines the EP alarm's objectId. 

**alarmNotificationOrigin** 
:  

**lastStatusChangeTime** 
:  

**alarmResource** 
: For AMS alarm, Altiplano will fill SOURCE_FRIENDLY_NAME or REPORTING_FOR or NE_ALIAS_NAME 

**durationOpen** 
:  

**deviceRefId** 
:  

**serviceAffecting** 
: This field is filled only for AMS & NSP alarms, not present in AV alarms 

**alarmType** 
: Indicates the type of issue, eg. `connection-failure` 

**alarmSeverity** 
: This object is defined to assign a severity level to an alarm.
This object may be indeterminate, warning, minor, major and
critical. 

**alarmStatus** 
:  

**alarmText** 
: A description of the issue encountered 

**neIpAddress** 
: This object defines the IP address of the NE for an alarm. 

**tl1Cause** 
:  

**eventType** 
:  

**proposedRepairAction** 
: Provides suggestions on how to resolve the issue. 

**alarmTypeQualifier** 
:  

**customField1** 
: This object defines the customField1 of an alarm. 

**customField2** 
: This object defines the customField2 of an alarm. 

**customField3** 
: This object defines the customField3 of an alarm. 


## Meaning
[//]: # "Short paragraph that explains what the alert means"

Any issues found in Altiplano are sent as this type of trap.  The trap is then converted into an alertmanager alert.  The true meaning of the alarm can be found in `alarmText`.

## Impact
[//]: # "What could / will happen if the alert is not addressed"

The `serviceAffecting` variable can provide a hint as to whether or not this alarm is causing issues to the customers, but care should be taken to determine the true impact.  Often alarms are set as serviceAffecting when the issue is not impacting customers.

Depending on the actual alarm, it could indicate issues that may lead to customer impact down the road.

## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"

- Look at the values contained in `alarmText`, `alarmStatus`, and `alarmSeverity`.
- Check the Altiplano Alarm Analyzier to see if the alert is still active and/or any related issues.


## Mitigation
[//]: # "The steps necessary to resolve the alert"

See `proposedRepairAction` for hints on resolving a particular alert.

---




