---
title: vrrpv3NewMaster
description: Troubleshooting for SNMP trap vrrpv3NewMaster
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# VRRPV3-MIB::vrrpv3NewMaster 

The newMaster notification indicates that the sending
agent has transitioned to master state. 


## Variables


  - vrrpv3OperationsMasterIpAddr
  - vrrpv3StatisticsNewMasterReason 

### Definitions 


**vrrpv3OperationsMasterIpAddr** 
: The master router's real IP address.  The master router
would set this address to vrrpv3OperationsPrimaryIpAddr
while transitioning to master state.  For backup
routers, this is the IP address listed as the source in
the VRRP advertisement last received by this virtual
router. 

**vrrpv3StatisticsNewMasterReason** 
: This indicates the reason for the virtual router to
transition to master state.  If the virtual router
never transitioned to master state, the value of this
object is notMaster(0).  Otherwise, this indicates the
reason this virtual router transitioned to master
state the last time.  Used by vrrpv3NewMaster
notification. 


## Meaning

The VRRPV3-MIB::vrrpv3NewMaster SNMP trap indicates that the sending agent has transitioned to master state. This trap is generated when a router takes over as the master virtual router in a VRRP (Virtual Router Redundancy Protocol) setup. This transition can occur due to various reasons, such as the failure of the previous master router or the expiration of the master router's timer.

## Impact

The impact of this trap depends on the reason for the transition to master state. If the transition is due to a failure of the previous master router, it may cause a brief disruption in network traffic. However, if the transition is due to a planned event, such as a maintenance window, the impact should be minimal. In either case, the new master router will take over the responsibilities of the previous master, and network traffic should continue to flow normally.

## Diagnosis

To diagnose the cause of the vrrpv3NewMaster trap, check the values of the following variables:

* vrrpv3OperationsMasterIpAddr: This variable indicates the IP address of the new master router.
* vrrpv3StatisticsNewMasterReason: This variable indicates the reason for the transition to master state.

By analyzing these variables, you can determine the cause of the transition and take appropriate actions to ensure network stability.

## Mitigation

To mitigate the impact of the vrrpv3NewMaster trap, follow these steps:

1. Verify that the new master router is functioning correctly and that network traffic is flowing normally.
2. Investigate the reason for the transition to master state using the vrrpv3StatisticsNewMasterReason variable.
3. Take corrective action to prevent future transitions due to the same reason, such as replacing a failed router or adjusting VRRP timers.
4. Monitor the network for any issues related to the transition and take prompt action to resolve them.

By following these steps, you can ensure that the network remains stable and that the transition to master state does not cause significant disruptions.
---




