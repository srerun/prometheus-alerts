---
title: isisProtocolsSupportedMismatch
description: Troubleshooting for SNMP trap isisProtocolsSupportedMismatch
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisProtocolsSupportedMismatch 

A notification sent when a non-pseudonode
segment 0 LSP is received that has no matching
protocols supported.  This may be because the system
does not generate the field, or because there are no
common elements.  The list of protocols supported
should be included in the notification: it may be
empty if the TLV is not supported, or if the
TLV is empty.
The agent must throttle the generation of
consecutive isisProtocolsSupportedMismatch
notifications so that there is at least a 5-second
gap between notifications of this type.  When
notifications are throttled, they are dropped, not
queued for sending at a future time. 


## Variables


  - isisNotificationSysLevelIndex
  - isisNotificationCircIfIndex
  - isisPduProtocolsSupported
  - isisPduLspId
  - isisPduFragment 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduProtocolsSupported** 
: The list of protocols supported by an
adjacent system.  This may be empty. 

**isisPduLspId** 
: An Octet String that uniquely identifies
a Link State PDU. 

**isisPduFragment** 
: Holds up to 64 initial bytes of a PDU that
triggered the notification. 


Here is a runbook for the SNMP Trap description:

## Meaning

The `ISIS-MIB::isisProtocolsSupportedMismatch` notification is sent when a non-pseudonode segment 0 LSP is received with no matching protocols supported. This can occur when the system does not generate the field or when there are no common elements. The notification includes the list of protocols supported, which may be empty if the TLV is not supported or if the TLV is empty.

## Impact

This notification may indicate a mismatch in the ISIS protocols supported between two systems, which can lead to issues in the network topology and routing. This can cause problems with network convergence, stability, and performance. It is essential to investigate and resolve this issue to ensure the reliable operation of the network.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the **isisNotificationSysLevelIndex** and **isisNotificationCircIfIndex** variables to identify the system and circuit interface involved.
2. Examine the **isisPduProtocolsSupported** variable to determine the list of protocols supported by the adjacent system.
3. Verify the **isisPduLspId** and **isisPduFragment** variables to identify the specific LSP and PDU that triggered the notification.
4. Check the network configuration and ISIS protocol settings to ensure that they match between the involved systems.
5. Analyze the network topology and routing tables to identify any potential issues or inconsistencies.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify that all systems in the network have the same ISIS protocol settings and configurations.
2. Check for any software or firmware issues that may be causing the mismatch and apply any necessary updates.
3. Ensure that the network topology and routing tables are up-to-date and consistent across all systems.
4. If the issue persists, consider isolating the problematic segment or circuit to prevent further impact on the network.
5. Monitor the network closely for any further notifications or issues related to ISIS protocol mismatches.
---




