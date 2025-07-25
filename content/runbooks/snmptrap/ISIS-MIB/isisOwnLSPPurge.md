---
title: isisOwnLSPPurge
description: Troubleshooting for SNMP trap isisOwnLSPPurge
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisOwnLSPPurge 

A notification sent when we receive a PDU
with our systemID and zero age.  This
notification includes the circuit Index
and router ID from the LSP, if available,
which may help a network manager
identify the source of the confusion. 


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


Here is the runbook for the SNMP trap ISIS-MIB::isisOwnLSPPurge:

## Meaning

This SNMP trap indicates that the router has received a Link State Protocol Data Unit (LSP) with its own system ID and zero age. This can cause confusion in the network and potentially lead to routing issues.

## Impact

The impact of this trap is that the router may become confused about its own state and may start to advertise incorrect routing information to its neighbors. This can lead to network instability, routing loops, and decreased network performance.

## Diagnosis

To diagnose the issue, the following steps can be taken:

* Check the circuit index and router ID from the LSP to identify the source of the confusion.
* Verify the system level and circuit interface index using the `isisNotificationSysLevelIndex` and `isisNotificationCircIfIndex` variables.
* Analyze the LSP ID using the `isisPduLspId` variable to identify the specific LSP that is causing the issue.
* Check the router's logs and configuration to see if there are any issues with the ISIS protocol or if there have been any recent changes to the network topology.

## Mitigation

To mitigate the issue, the following steps can be taken:

* Verify that the ISIS protocol is configured correctly on the router and that there are no issues with the protocol's operation.
* Check for any duplicate or conflicting LSPs in the network and remove or correct them as necessary.
* Verify that the router's system ID and circuit interface index are correctly configured and match the values in the LSP.
* Consider implementing measures to prevent LSP duplication, such as configuring ISIS to ignore LSUs with zero age.
* If the issue persists, consider restarting the ISIS protocol or the router itself to clear any corruption or incorrect state.
---




