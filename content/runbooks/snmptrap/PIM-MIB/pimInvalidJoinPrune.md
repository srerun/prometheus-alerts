---
title: pimInvalidJoinPrune
description: Troubleshooting for SNMP trap pimInvalidJoinPrune
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# PIM-MIB::pimInvalidJoinPrune 

A pimInvalidJoinPrune notification signifies that an
invalid PIM Join/Prune message was received by this device.
This notification is generated whenever the counter
pimInvalidJoinPruneMsgsRcvd is incremented, subject to the
rate limit specified by
pimInvalidJoinPruneNotificationPeriod. 


## Variables


  - pimGroupMappingPimMode
  - pimInvalidJoinPruneAddressType
  - pimInvalidJoinPruneOrigin
  - pimInvalidJoinPruneGroup
  - pimInvalidJoinPruneRp
  - pimNeighborUpTime 

### Definitions 


**pimGroupMappingPimMode** 
: The PIM mode to be used for groups in this group prefix. 

**pimInvalidJoinPruneAddressType** 
: The address type stored in pimInvalidJoinPruneOrigin,
pimInvalidJoinPruneGroup, and pimInvalidJoinPruneRp.
If no invalid Join/Prune messages have been received, this
object is set to unknown(0). 

**pimInvalidJoinPruneOrigin** 
: The source address of the last invalid Join/Prune message
received by this device. 

**pimInvalidJoinPruneGroup** 
: The IP multicast group address carried in the last
invalid Join/Prune message received by this device. 

**pimInvalidJoinPruneRp** 
: The RP address carried in the last invalid Join/Prune
message received by this device. 

**pimNeighborUpTime** 
: The time since this PIM neighbor (last) became a neighbor
of the local router. 


Here is a runbook for the SNMP Trap `PIM-MIB::pimInvalidJoinPrune`:

## Meaning

The `PIM-MIB::pimInvalidJoinPrune` trap indicates that the device has received an invalid PIM Join/Prune message. This trap is generated when the `pimInvalidJoinPruneMsgsRcvd` counter is incremented, subject to a rate limit specified by `pimInvalidJoinPruneNotificationPeriod`.

## Impact

Receiving invalid PIM Join/Prune messages can have several negative impacts on the network:

* Inefficient use of network resources
* Increased risk of network instability
* Potential for multicast traffic to be incorrectly forwarded or not forwarded at all
* Inability to properly manage PIM neighbors and multicast groups

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the `pimInvalidJoinPruneMsgsRcvd` counter to determine the number of invalid PIM Join/Prune messages received.
2. Verify the `pimInvalidJoinPruneNotificationPeriod` rate limit to ensure it is set correctly.
3. Use the `pimInvalidJoinPruneAddressType`, `pimInvalidJoinPruneOrigin`, `pimInvalidJoinPruneGroup`, and `pimInvalidJoinPruneRp` variables to identify the source and contents of the invalid message.
4. Check the `pimNeighborUpTime` variable to determine the uptime of the PIM neighbor that sent the invalid message.
5. Analyze PIM neighbor and group mappings using `pimGroupMappingPimMode` to ensure they are correctly configured.

## Mitigation

To mitigate the issue, perform the following steps:

1. Verify that PIM neighbors and group mappings are correctly configured.
2. Check the PIM implementation on the device and ensure it is up-to-date and correctly configured.
3. Implement filters or access lists to block invalid PIM Join/Prune messages from reaching the device.
4. Increase the `pimInvalidJoinPruneNotificationPeriod` rate limit to reduce the frequency of trap notifications.
5. Consider implementing additional monitoring and logging to detect and alert on invalid PIM Join/Prune messages.
---




