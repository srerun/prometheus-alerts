---
title: isisSequenceNumberSkip
description: Troubleshooting for SNMP trap isisSequenceNumberSkip
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisSequenceNumberSkip 

When we receive an LSP with our System ID
and different contents, we may need to reissue
the LSP with a higher sequence number.
We send this notification if we need to increase
the sequence number by more than one.  If two
Intermediate Systems are configured with the same
System ID, this notification will fire. 


## Variables


  - isisNotificationSysLevelIndex
  - isisNotificationCircIfIndex
  - isisPduLspId 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduLspId** 
: An Octet String that uniquely identifies
a Link State PDU. 


Here is a runbook for the ISIS-MIB::isisSequenceNumberSkip SNMP trap:

## Meaning

The ISIS-MIB::isisSequenceNumberSkip trap is generated when an Intermediate System (IS) receives a Link State PDU (LSP) with its own System ID but with different contents. To ensure consistency and prevent loop formation, the IS needs to reissue the LSP with a higher sequence number. This trap is sent if the sequence number needs to be increased by more than one.

## Impact

The impact of this trap is that it may indicate a configuration error or misbehavior in the IS-IS network. If two Intermediate Systems are configured with the same System ID, this trap will fire, and the network may experience instability or routing issues. The trap may also indicate a genuine need to reissue an LSP with a higher sequence number, which can cause temporary network disruptions.

## Diagnosis

To diagnose the root cause of this trap, follow these steps:

1. Verify that the System IDs of all Intermediate Systems in the network are unique.
2. Check the IS-IS configuration for any errors or inconsistencies.
3. Analyze the LSP contents to determine the reason for the sequence number skip.
4. Examine the IS-IS logs to identify any other related errors or issues.

## Mitigation

To mitigate the effects of this trap, follow these steps:

1. Ensure that all Intermediate Systems have unique System IDs.
2. Correct any configuration errors or inconsistencies in the IS-IS network.
3. Implement measures to prevent LSP sequence number skips, such as increasing the sequence number window or implementing sequence number synchronization mechanisms.
4. Monitor the IS-IS network closely to detect and respond to any related issues promptly.

Note: The variables provided can be used to gather more information about the trap and aid in diagnosis and mitigation. These variables include:

* isisNotificationSysLevelIndex: The system level for this notification.
* isisNotificationCircIfIndex: The identifier of the circuit relevant to this notification.
* isisPduLspId: An Octet String that uniquely identifies the Link State PDU.
---




