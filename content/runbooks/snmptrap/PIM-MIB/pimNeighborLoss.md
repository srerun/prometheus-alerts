---
title: pimNeighborLoss
description: Troubleshooting for SNMP trap pimNeighborLoss
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# PIM-MIB::pimNeighborLoss 

A pimNeighborLoss notification signifies the loss of an
adjacency with a neighbor.  This notification should be
generated when the neighbor timer expires, and the router
has no other neighbors on the same interface with the same
IP version and a lower IP address than itself.
This notification is generated whenever the counter
pimNeighborLossCount is incremented, subject
to the rate limit specified by
pimNeighborLossNotificationPeriod. 


## Variables


  - pimNeighborUpTime 

### Definitions 


**pimNeighborUpTime** 
: The time since this PIM neighbor (last) became a neighbor
of the local router. 


## Meaning

The PIM-MIB::pimNeighborLoss trap indicates that the router has lost adjacency with a PIM neighbor. This means that the router's PIM protocol has detected that a neighboring router is no longer reachable or has stopped responding.

## Impact

The impact of this trap is that multicast traffic may not be forwarded correctly, leading to potential disruptions in network communication. This can result in:

* Loss of multicast traffic
* Delayed or failed data delivery
* Increased network latency
* Potential instability in the multicast network

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the PIM neighbor table to identify the lost neighbor.
2. Verify the status of the affected interface and ensure it is operational.
3. Review the router's PIM configuration to ensure it is correct and up-to-date.
4. Check the pimNeighborUpTime variable to determine the duration of the neighbor relationship before the loss.
5. Analyze system logs for any error messages related to PIM or the affected interface.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify the physical connectivity of the affected interface.
2. Check for any configuration changes that may have caused the neighbor loss.
3. Restart the PIM process on the router to re-establish the neighbor relationship.
4. Monitor the pimNeighborLossCount counter to ensure the issue is resolved.
5. Adjust the pimNeighborLossNotificationPeriod to adjust the rate limit of the trap generation, if necessary.

Note: The specific mitigation steps may vary depending on the network environment and PIM configuration.
---




