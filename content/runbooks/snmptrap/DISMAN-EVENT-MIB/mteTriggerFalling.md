---
title: mteTriggerFalling
description: Troubleshooting for SNMP trap mteTriggerFalling
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# DISMAN-EVENT-MIB::mteTriggerFalling 

Notification that the falling threshold was met for triggers
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


Here is a runbook for the SNMP Trap description # DISMAN-EVENT-MIB::mteTriggerFalling:

## Meaning

The # DISMAN-EVENT-MIB::mteTriggerFalling notification is triggered when a falling threshold is met for triggers with mteTriggerType 'threshold'. This notification indicates that a trigger has been activated due to a decrease in value below a specified threshold.

## Impact

The impact of this notification depends on the specific trigger and the system being monitored. However, in general, it may indicate a potential issue or problem that requires attention. The trigger may be related to a critical system component, such as a network device or a server, and the falling threshold may indicate a degradation in performance or an approaching failure.

## Diagnosis

To diagnose the cause of this notification, follow these steps:

1. Identify the trigger name and target using the `mteHotTrigger` and `mteHotTargetName` variables.
2. Determine the context name and OID related to the notification using the `mteHotContextName` and `mteHotOID` variables.
3. Check the value of the object at `mteTriggerValueID` when the trigger fired, using the `mteHotValue` variable.
4. Investigate the system or device related to the trigger and assess the current status and any potential issues.
5. Review logs and monitoring data to identify any trends or patterns that may be contributing to the falling threshold.

## Mitigation

 Depending on the specific trigger and system being monitored, possible mitigation steps may include:

1. Investigate and address the root cause of the falling threshold, such as adjusting configuration or repairing hardware.
2. Implement additional monitoring and alerting to detect potential issues before they impact system performance.
3. Perform maintenance or maintenance tasks to prevent future occurrences of the falling threshold.
4. Update trigger threshold values to more accurately reflect system performance and avoid false positives.
5. Escalate the issue to a higher-level support team or subject matter expert if necessary.
---




