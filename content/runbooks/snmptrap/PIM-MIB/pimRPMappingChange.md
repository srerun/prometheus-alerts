---
title: pimRPMappingChange
description: Troubleshooting for SNMP trap pimRPMappingChange
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# PIM-MIB::pimRPMappingChange 

A pimRPMappingChange notification signifies a change to the
active RP mapping on this device.
This notification is generated whenever the counter
pimRPMappingChangeCount is incremented, subject to the
rate limit specified by
pimRPMappingChangeNotificationPeriod. 


## Variables


  - pimGroupMappingPimMode
  - pimGroupMappingPrecedence 

### Definitions 


**pimGroupMappingPimMode** 
: The PIM mode to be used for groups in this group prefix. 

**pimGroupMappingPrecedence** 
: The precedence of this row, used in the algorithm that
determines which row applies to a given group address
(described above).  Numerically higher values for this
object indicate lower precedences, with the value zero
denoting the highest precedence.
The absolute values of this object have a significance only
on the local router and do not need to be coordinated with
other routers. 


## Meaning

A pimRPMappingChange notification indicates a change to the active RP (Rendezvous Point) mapping on the device. This notification is triggered when the counter pimRPMappingChangeCount is incremented, subject to a rate limit specified by pimRPMappingChangeNotificationPeriod.

## Impact

This notification can have a significant impact on the network, as a change to the active RP mapping can affect the multicast routing and traffic flow. This may lead to:

* Disruption to multicast traffic and services
* Changes to the multicast tree and group membership
* Potential for packet loss or duplication
* Increased latency or jitter in multicast traffic

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the pimRPMappingChangeCount counter to determine if it has incremented.
2. Verify the pimRPMappingChangeNotificationPeriod setting to ensure it is within the expected range.
3. Review the pimGroupMappingPimMode and pimGroupMappingPrecedence values to determine if they have changed.
4. Analyze the multicast routing table and group membership to identify any changes.
5. Check for any error messages or logs related to PIM and multicast routing.

## Mitigation

To mitigate the impact of a pimRPMappingChange notification, follow these steps:

1. Identify the root cause of the RP mapping change (e.g., configuration change, network topology change, etc.).
2. Verify that the new RP mapping is correct and optimal for the network.
3. Update the pimGroupMappingPimMode and pimGroupMappingPrecedence values accordingly.
4. Monitor the multicast traffic and services to ensure they are functioning as expected.
5. Consider implementing redundancy and backup mechanisms for the RP to minimize the impact of future changes.

Note: The specific mitigation steps may vary depending on the network architecture, topology, and multicast configuration. It is essential to understand the root cause of the issue and take targeted actions to resolve it.
---




