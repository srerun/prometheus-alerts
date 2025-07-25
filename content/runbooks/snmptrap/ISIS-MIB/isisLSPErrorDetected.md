---
title: isisLSPErrorDetected
description: Troubleshooting for SNMP trap isisLSPErrorDetected
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisLSPErrorDetected 

This notification is generated when we receive
an LSP with a parse error.  The isisCircIfIndex
holds an index of the circuit on which the PDU
arrived.  The isisPduFragment holds the start of the
LSP, and the isisErrorOffset points to the problem.
If the problem is a malformed TLV, isisErrorOffset
points to the start of the TLV, and isisErrorTLVType
holds the value of the type.
If the problem is with the LSP header, isisErrorOffset
points to the suspicious byte.
The number of such LSPs is accumulated in
isisSysStatLSPErrors. 


## Variables


  - isisNotificationSysLevelIndex
  - isisPduLspId
  - isisNotificationCircIfIndex
  - isisPduFragment
  - isisErrorOffset
  - isisErrorTLVType 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisPduLspId** 
: An Octet String that uniquely identifies
a Link State PDU. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduFragment** 
: Holds up to 64 initial bytes of a PDU that
triggered the notification. 

**isisErrorOffset** 
: An offset to a problem in a PDU.  If the problem
is a malformed TLV, this points to the beginning
of the TLV.  If the problem is in the header, this
points to the byte that is suspicious. 

**isisErrorTLVType** 
: The type for a malformed TLV. 


## Meaning

The ISIS-MIB::isisLSPErrorDetected trap notification is generated when a Link State Protocol Data Unit (LSP) with a parse error is received. This notification provides information about the circuit on which the PDU arrived, the start of the LSP, and the location and type of the error.

## Impact

The impact of this trap notification is that the LSP with a parse error cannot be processed, which may affect the integrity and accuracy of the link-state database. This can lead to incorrect routing decisions, network instability, and potential outages.

## Diagnosis

To diagnose the issue, the following steps can be taken:

1. Check the `isisNotificationCircIfIndex` variable to identify the circuit on which the PDU arrived.
2. Use the `isisPduFragment` variable to examine the initial bytes of the PDU that triggered the notification.
3. Analyze the `isisErrorOffset` variable to determine the location of the error in the PDU.
4. If the error is a malformed TLV, check the `isisErrorTLVType` variable to determine the type of TLV that is malformed.
5. Review the `isisSysStatLSPErrors` variable to see how many LSPs with parse errors have been received.

## Mitigation

To mitigate the issue, the following steps can be taken:

1. Investigate the circuit identified by `isisNotificationCircIfIndex` to determine if there are any issues with the circuit or the device attached to it.
2. Check the device logs to see if there are any error messages related to the LSP reception.
3. Verify that the device is configured correctly and that the LSPs being received are valid.
4. If the error is a malformed TLV, check the device configuration to ensure that the TLV is correctly formatted.
5. Implement measures to prevent LSPs with parse errors from being received, such as filtering or rate-limiting incoming LSPs.
---




