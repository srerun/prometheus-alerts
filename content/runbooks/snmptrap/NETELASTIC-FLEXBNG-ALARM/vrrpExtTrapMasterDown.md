---
title: vrrpExtTrapMasterDown
description: Troubleshooting for SNMP trap vrrpExtTrapMasterDown
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# NETELASTIC-FLEXBNG-ALARM::vrrpExtTrapMasterDown 

The MasterDown trap indicates that the state of vrrp
has transitioned from 'Master' to other state. The other state
can be noactive(0), initialize(1) and backup(2). 


## Variables


  - vrrpMasterIpAddr
  - vrrpStatus
  - vrrpStateChangeReasonString 

### Definitions 


**vrrpMasterIpAddr** 
: The master router's real IP address.  For backup
routers, this is the IP address listed as the source in
the VRRP advertisement last received by this virtual
router. 

**vrrpStatus** 
: The current state of the virtual router.  This object
has three defined values:
- 'initialize', which indicates that the
virtual router is waiting for a startup event.
- 'backup', which indicates that the virtual router is
monitoring the availability of the master router.
- 'master', which indicates that the virtual router
is forwarding packets for IP addresses that are
associated with this router. 

**vrrpStateChangeReasonString** 
: Reasons of VRRP state transition.
Used by vrrpExtTrapMasterDown trap. 


Here is a runbook for the SNMP trap `NETELASTIC-FLEXBNG-ALARM::vrrpExtTrapMasterDown`:

## Meaning

The `vrrpExtTrapMasterDown` trap indicates that the state of the VRRP (Virtual Router Redundancy Protocol) has transitioned from 'Master' to another state, such as 'noactive', 'initialize', or 'backup'. This trap is triggered when the master router fails or becomes unavailable, and another router takes over as the master.

## Impact

The impact of this trap is that the network may experience temporary disruption or degradation of service. The transition from master to another state may cause packet loss or delay, affecting network availability and reliability. Additionally, the failure of the master router may indicate a larger network issue that needs to be addressed.

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the `vrrpStatus` variable to determine the current state of the virtual router.
2. Check the `vrrpMasterIpAddr` variable to identify the IP address of the original master router.
3. Check the `vrrpStateChangeReasonString` variable to determine the reason for the VRRP state transition.
4. Investigate the system logs and network monitoring tools to identify any related errors or issues.
5. Verify that the VRRP configuration is correct and that the backup routers are configured correctly.

## Mitigation

To mitigate the issue, perform the following steps:

1. Verify that the backup router has taken over as the master router and is functioning correctly.
2. Investigate and resolve the underlying cause of the master router failure.
3. Perform a failback to the original master router once it is available and stable.
4. Verify that network services are restored and functioning correctly.
5. Review and update the VRRP configuration to prevent similar issues in the future.
---




