---
title: isisAdjacencyChange
description: Troubleshooting for SNMP trap isisAdjacencyChange
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisAdjacencyChange 

A notification sent when an adjacency changes
state, entering or leaving state up.
The first 6 bytes of the isisPduLspId are the
SystemID of the adjacent IS.
The isisAdjState is the new state of the adjacency. 


## Variables


  - isisNotificationSysLevelIndex
  - isisNotificationCircIfIndex
  - isisPduLspId
  - isisAdjState 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduLspId** 
: An Octet String that uniquely identifies
a Link State PDU. 

**isisAdjState** 
: The current state of an adjacency. 


Here is a runbook for the SNMP trap description:

## Meaning

The ISIS-MIB::isisAdjacencyChange notification is sent when the state of an adjacency changes, either entering or leaving the "up" state. This trap is generated when there is a change in the adjacency state between two Intermediate System (IS) nodes.

## Impact

The impact of this trap is moderate. The change in adjacency state can affect the routing of packets in the network, potentially leading to network instability or packet loss. If the adjacency change is not properly diagnosed and mitigated, it can result in prolonged network downtime or performance issues.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the **isisNotificationSysLevelIndex** to identify the system level that generated the notification.
2. Check the **isisNotificationCircIfIndex** to identify the circuit interface related to the notification.
3. Analyze the **isisPduLspId** to determine the unique Link State PDU (LSP) identifier associated with the adjacency change.
4. Examine the **isisAdjState** to determine the new state of the adjacency (e.g., up or down).
5. Verify the configuration of the adjacent IS nodes to ensure that they are properly set up and communicating correctly.
6. Use network management tools (e.g., network monitoring software, protocol analyzers) to inspect the network traffic and verify that packets are being routed correctly.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify that the adjacent IS nodes are properly configured and communicating correctly.
2. Check for any hardware or software issues on the affected nodes that may be causing the adjacency change.
3. If the adjacency change is due to a hardware failure, replace the faulty hardware component.
4. If the adjacency change is due to a software issue, apply the necessary patches or upgrades to resolve the problem.
5. Monitor the network closely to ensure that the adjacency change does not cause any prolonged network instability or performance issues.
6. Consider implementing redundancy or backup mechanisms to minimize the impact of future adjacency changes.
---




