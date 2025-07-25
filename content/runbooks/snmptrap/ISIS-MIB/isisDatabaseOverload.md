---
title: isisDatabaseOverload
description: Troubleshooting for SNMP trap isisDatabaseOverload
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisDatabaseOverload 

This notification is generated when the system
enters or leaves the Overload state.  The number
of times this has been generated and cleared is kept
track of by isisSysStatLSPDbaseOloads. 


## Variables


  - isisNotificationSysLevelIndex
  - isisSysLevelState 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisSysLevelState** 
: The state of the database at this level.
The value 'off' indicates that IS-IS is not active at
this level.
The value 'on' indicates that IS-IS is active at this
level and is not overloaded.
The value 'waiting' indicates a database that is low on
an essential resource, such as memory.
The administrator may force the state to 'overloaded'
by setting the object isisSysLevelSetOverload.
If the state is 'waiting' or 'overloaded', we
originate LSPs with the overload bit set. 


## Meaning

The ISIS-MIB::isisDatabaseOverload SNMP trap is generated when the system enters or leaves the Overload state. This trap indicates that the IS-IS database has reached a critical condition, affecting the system's ability to process Link State Packets (LSPs). This state is tracked by the isisSysStatLSPDbaseOloads counter.

## Impact

When the system enters the Overload state, it may lead to:

* Slowed or delayed LSP processing, affecting network convergence and routing decisions
* Increased memory utilization, potentially causing system instability
* Reduced network reliability and availability
* Inaccurate or incomplete routing information, leading to suboptimal network performance

## Diagnosis

To diagnose the issue, follow these steps:

1. Verify the system's current state using the `isisSysLevelState` variable. Check if the state is 'overloaded', 'waiting', or 'off'.
2. Check the `isisNotificationSysLevelIndex` variable to identify the affected system level.
3. Review the `isisSysStatLSPDbaseOloads` counter to determine the number of times the system has entered or cleared the Overload state.
4. Analyze system resources, such as memory and CPU utilization, to identify potential bottlenecks.
5. Inspect IS-IS configuration and logs for any errors or inconsistencies.

## Mitigation

To mitigate the issue, follow these steps:

1. Clear the Overload state by setting the `isisSysLevelSetOverload` object to 'off' or restarting the IS-IS process.
2. Identify and address the root cause of the Overload state, such as memory or resource constraints.
3. Consider implementing IS-IS tuning and optimization techniques to improve system performance and resilience.
4. Monitor the system closely to ensure it does not re-enter the Overload state.
5. Develop a long-term plan to upgrade or add resources to prevent similar issues in the future.
---




