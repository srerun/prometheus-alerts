---
title: isisOrigLSPBuffSizeMismatch
description: Troubleshooting for SNMP trap isisOrigLSPBuffSizeMismatch
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisOrigLSPBuffSizeMismatch 

A notification sent when a Level 1 LSP or Level
2 LSP is received that is larger than the local
value for isisSysLevelOrigLSPBuffSize, or when an
LSP is received that contains the supported Buffer Size
option and the value in the PDU option field does
not match the local value for isisSysLevelOrigLSPBuffSize.
We pass up the size from the option field and the
size of the LSP when one of them exceeds our configuration.
The agent must throttle the generation of
consecutive isisOrigLSPBuffSizeMismatch notifications
so that there is at least a 5-second gap
between notifications of this type.  When
notifications are throttled, they are dropped, not
queued for sending at a future time. 


## Variables


  - isisNotificationSysLevelIndex
  - isisNotificationCircIfIndex
  - isisPduLspId
  - isisPduOriginatingBufferSize
  - isisPduBufferSize 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduLspId** 
: An Octet String that uniquely identifies
a Link State PDU. 

**isisPduOriginatingBufferSize** 
: Holds the size of isisSysLevelOrigLSPBuffSize advertised
by the peer in the originatingLSPBufferSize TLV.
If the peer does not advertise this TLV, this
value is set to 0. 

**isisPduBufferSize** 
: Holds the size of LSP received from peer. 


Here is a runbook for the SNMP Trap description:

## Meaning

The ISIS-MIB::isisOrigLSPBuffSizeMismatch notification indicates that a Level 1 or Level 2 Link State PDU (LSP) has been received that is larger than the local value for isisSysLevelOrigLSPBuffSize or contains a supported Buffer Size option that does not match the local value. This mismatch can prevent proper communication between ISIS routers and potentially cause network instability.

## Impact

The impact of this notification is that the ISIS router may not be able to process the received LSP correctly, leading to potential issues with network convergence and routing decisions. This can result in:

* Network instability and convergence issues
* Routing loops or blackholes
* Increased network latency and packet loss

## Diagnosis

To diagnose the issue, check the following:

* Verify the local value of isisSysLevelOrigLSPBuffSize and ensure it is correctly configured.
* Check the received LSP size and the supported Buffer Size option value to identify the mismatch.
* Use the variables provided in the notification to identify the affected circuit and system level:
	+ isisNotificationSysLevelIndex: The system level for this notification.
	+ isisNotificationCircIfIndex: The identifier of this circuit relevant to this notification.
* Analyze the ISIS router logs and debug outputs to identify any other related errors or issues.

## Mitigation

To mitigate the issue, perform the following steps:

* Verify and update the local configuration of isisSysLevelOrigLSPBuffSize to match the received LSP size or the supported Buffer Size option value.
* Ensure that the ISIS router is configured to handle LSPs of the correct size.
* Implement rate limiting or throttling on the ISIS router to prevent excessive notifications and potential network instability.
* Monitor the network and ISIS router logs for any further issues or errors related to LSP processing and routing decisions.
---




