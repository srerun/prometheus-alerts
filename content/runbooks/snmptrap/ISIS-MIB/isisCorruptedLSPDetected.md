---
title: isisCorruptedLSPDetected
description: Troubleshooting for SNMP trap isisCorruptedLSPDetected
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisCorruptedLSPDetected 

This notification is generated when we find that
an LSP that was stored in memory has become
corrupted.  The number of times this has been
generated is counted by isisSysCorrLSPs.
We forward an LSP ID.  We may have independent
knowledge of the ID, but in some implementations
there is a chance that the ID itself will be
corrupted. 


## Variables


  - isisNotificationSysLevelIndex
  - isisPduLspId 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisPduLspId** 
: An Octet String that uniquely identifies
a Link State PDU. 


Here is a runbook for the SNMP trap `ISIS-MIB::isisCorruptedLSPDetected`:

## Meaning

This SNMP trap indicates that a corrupted Link State PDU (LSP) has been detected in memory. The LSP ID is provided in the trap, but it's possible that the ID itself may be corrupted.

## Impact

The impact of this trap is that the integrity of the LSP database is compromised, which can affect the accuracy of the network topology and may lead to routing issues. This can cause network instability, packet loss, and decreased network performance.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `isisSysCorrLSPs` counter to determine the number of times this trap has been generated.
2. Verify the system level of the notification using the `isisNotificationSysLevelIndex` variable.
3. Attempt to retrieve the corrupted LSP ID from the trap using the `isisPduLspId` variable.
4. Check the LSP database for any inconsistencies or corruption.
5. Review system logs for any related errors or warnings.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the ISIS process to attempt to recover from the corrupted LSP.
2. Perform a manual LSP database consistency check and repair any inconsistencies found.
3. Check for any software or firmware updates that may address the issue.
4. Implement additional monitoring and logging to detect similar issues in the future.
5. Consider configuring additional redundancy or backup mechanisms to minimize the impact of future LSP corruption events.
---




