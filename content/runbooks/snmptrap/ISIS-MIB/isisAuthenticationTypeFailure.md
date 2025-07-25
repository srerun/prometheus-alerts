---
title: isisAuthenticationTypeFailure
description: Troubleshooting for SNMP trap isisAuthenticationTypeFailure
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisAuthenticationTypeFailure 

A notification sent when we receive a PDU
with the wrong authentication type field.
This notification includes the header of the
packet, which may help a network manager
identify the source of the confusion.
The agent must throttle the generation of
consecutive isisAuthenticationTypeFailure
notifications so that there is at least a 5-second
gap between notifications of this type.  When
notifications are throttled, they are dropped, not
queued for sending at a future time. 


## Variables


  - isisNotificationSysLevelIndex
  - isisNotificationCircIfIndex
  - isisPduFragment 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduFragment** 
: Holds up to 64 initial bytes of a PDU that
triggered the notification. 


Here is a runbook for the ISIS-MIB::isisAuthenticationTypeFailure SNMP trap:

## Meaning

The ISIS-MIB::isisAuthenticationTypeFailure notification is sent when a device receives a PDU (Protocol Data Unit) with an incorrect authentication type field. This error can occur due to configuration issues, authentication protocol mismatch, or malicious activity.

## Impact

The impact of this error can be significant, as it may:

* Prevent or disrupt routing information exchange between ISIS (Intermediate System to Intermediate System) routers
* Cause routing instability or blackholes in the network
* Allow unauthorized access to the network or compromise the integrity of routing information

## Diagnosis

To diagnose the issue, follow these steps:

1. Review the trap notification to identify the system level (`isisNotificationSysLevelIndex`) and circuit interface (`isisNotificationCircIfIndex`) involved in the error.
2. Analyze the `isisPduFragment` to determine the contents of the PDU that triggered the error.
3. Verify the ISIS configuration on the affected devices to ensure that the authentication type is correctly set.
4. Check the network for any signs of malicious activity or unauthorized access.
5. Use network management tools to monitor the network and identify any routing issues or instability.

## Mitigation

To mitigate the issue, follow these steps:

1. Correct the ISIS configuration on the affected devices to ensure that the authentication type is set correctly.
2. Verify that all ISIS devices in the network are configured to use the same authentication type.
3. Implement additional security measures to prevent unauthorized access to the network.
4. Monitor the network closely for any signs of routing instability or malicious activity.
5. Consider implementing rate limiting or filtering on the affected interfaces to prevent further authentication type failures.
---




