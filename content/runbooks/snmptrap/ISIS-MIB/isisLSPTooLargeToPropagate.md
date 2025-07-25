---
title: isisLSPTooLargeToPropagate
description: Troubleshooting for SNMP trap isisLSPTooLargeToPropagate
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisLSPTooLargeToPropagate 

A notification sent when we attempt to propagate
an LSP that is larger than the dataLinkBlockSize
for the circuit.
The agent must throttle the generation of
consecutive isisLSPTooLargeToPropagate notifications
so that there is at least a 5-second gap
between notifications of this type.  When
notifications are throttled, they are dropped, not
queued for sending at a future time. 


## Variables


  - isisNotificationSysLevelIndex
  - isisNotificationCircIfIndex
  - isisPduLspSize
  - isisPduLspId 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduLspSize** 
: Holds the size of LSP we received that is too
big to forward. 

**isisPduLspId** 
: An Octet String that uniquely identifies
a Link State PDU. 


Here is the runbook for the SNMP trap `ISIS-MIB::isisLSPTooLargeToPropagate`:

## Meaning

The `ISIS-MIB::isisLSPTooLargeToPropagate` trap is sent when the device attempts to propagate a Link State PDU (LSP) that is larger than the maximum allowed size for the circuit. This trap indicates that the LSP is too large to be forwarded, which can cause issues with network convergence and routing.

## Impact

The impact of this trap is that the LSP will not be propagated, which can lead to incomplete or inconsistent routing information in the network. This can cause routing loops, black holes, or other issues that affect network availability and performance. Additionally, the throttling of consecutive notifications may lead to delayed detection of the issue, making it more challenging to troubleshoot and resolve.

## Diagnosis

To diagnose the issue, the following steps can be taken:

1. Verify the LSP size and compare it to the maximum allowed size for the circuit.
2. Check the circuit configuration to ensure that the maximum allowed size is correctly set.
3. Review the routing table to identify any inconsistencies or errors.
4. Analyze the network topology to identify any potential routing loops or black holes.
5. Check the device logs for any other related errors or issues.

## Mitigation

To mitigate the issue, the following steps can be taken:

1. Reduce the size of the LSP to ensure it is within the maximum allowed size for the circuit.
2. Adjust the circuit configuration to increase the maximum allowed size for the LSP.
3. Implement routing policies or filters to prevent oversized LSPs from being propagated.
4. Ensure that the network topology is correctly configured to prevent routing loops or black holes.
5. Monitor the network and device logs for any related issues or errors to quickly detect and respond to any future occurrences.
---




