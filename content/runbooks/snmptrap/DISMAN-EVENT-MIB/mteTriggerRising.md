---
title: mteTriggerRising
description: Troubleshooting for SNMP trap mteTriggerRising
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# DISMAN-EVENT-MIB::mteTriggerRising 

Notification that the rising threshold was met for triggers
with mteTriggerType 'threshold'. 


## Variables


  - mteHotTrigger
  - mteHotTargetName
  - mteHotContextName
  - mteHotOID
  - mteHotValue 

### Definitions 


**mteHotTrigger** 
: The name of the trigger causing the notification. 

**mteHotTargetName** 
: The SNMP Target MIB's snmpTargetAddrName related to the
notification. 

**mteHotContextName** 
: The context name related to the notification.  This MUST be as
fully-qualified as possible, including filling in wildcard
information determined in processing. 

**mteHotOID** 
: The object identifier of the destination object related to the
notification.  This MUST be as fully-qualified as possible,
including filling in wildcard information determined in
processing.
For a trigger-related notification this is from
mteTriggerValueID.
For a set failure this is from mteEventSetObject. 

**mteHotValue** 
: The value of the object at mteTriggerValueID when a
trigger fired. 


Here is a runbook for the SNMP Trap # DISMAN-EVENT-MIB::mteTriggerRising:

## Meaning

This SNMP trap is generated when a rising threshold is met for triggers with a type of 'threshold'. This indicates that a monitored value has exceeded a predefined threshold, and an event has been triggered.

## Impact

The impact of this trap depends on the specific trigger and threshold that has been exceeded. It may indicate a critical issue requiring immediate attention, or it may be a warning of a potential issue. In any case, it is essential to investigate and respond promptly to prevent potential service disruptions or data losses.

## Diagnosis

To diagnose the root cause of this trap, follow these steps:

1. Identify the trigger that caused the notification by examining the `mteHotTrigger` variable.
2. Determine the target of the notification by examining the `mteHotTargetName` variable.
3. Identify the context related to the notification by examining the `mteHotContextName` variable.
4. Identify the specific object that exceeded the threshold by examining the `mteHotOID` variable.
5. Check the current value of the object that exceeded the threshold by examining the `mteHotValue` variable.
6. Review the trigger configuration and threshold settings to understand why the threshold was exceeded.
7. Check system logs and monitoring data to gather more information about the event.

## Mitigation

To mitigate the issue, follow these steps:

1. Take immediate action to address the underlying issue that caused the threshold to be exceeded.
2. Adjust the trigger configuration and threshold settings as needed to prevent similar events from occurring in the future.
3. Verify that the system or device is functioning normally and that services are available.
4. Document the event and the steps taken to resolve it in a knowledge base or incident management system.
5. Inform stakeholders and teams of the event and provide updates on the status of the mitigation efforts.
---




