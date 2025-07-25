---
title: ipv6IfStateChange
description: Troubleshooting for SNMP trap ipv6IfStateChange
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# IPV6-MIB::ipv6IfStateChange 

An ipv6IfStateChange notification signifies
that there has been a change in the state of
an ipv6 interface.  This notification should
be generated when the interface's operational
status transitions to or from the up(1) state. 


## Variables


  - ipv6IfDescr
  - ipv6IfOperStatus 

### Definitions 


**ipv6IfDescr** 
: A textual string containing information about the
interface.  This string may be set by the network
management system. 

**ipv6IfOperStatus** 
: The current operational state of the interface.
The noIfIdentifier(3) state indicates that no valid
Interface Identifier is assigned to the interface.
This state usually indicates that the link-local
interface address failed Duplicate Address Detection.
If ipv6IfAdminStatus is down(2) then ipv6IfOperStatus
should be down(2).  If ipv6IfAdminStatus is changed
to up(1) then ipv6IfOperStatus should change to up(1)
if the interface is ready to transmit and receive
network traffic; it should remain in the down(2) or
noIfIdentifier(3) state if and only if there is a
fault that prevents it from going to the up(1) state;
it should remain in the notPresent(5) state if
the interface has missing (typically, lower layer)
components. 


Here is a runbook for the IPV6-MIB::ipv6IfStateChange SNMP Trap:

## Meaning

The IPV6-MIB::ipv6IfStateChange notification indicates that there has been a change in the state of an IPv6 interface. This change is specifically related to the interface's operational status transitioning to or from the "up" state.

## Impact

This notification can have a significant impact on network operations and connectivity. If the interface's operational status has changed to "down", it may cause connectivity issues for devices relying on that interface. Conversely, if the interface's operational status has changed to "up", it may restore connectivity and allow devices to communicate again.

## Diagnosis

To diagnose the issue, you can use the following variables:

* `ipv6IfDescr`: Check the interface description to identify the specific interface that has changed state.
* `ipv6IfOperStatus`: Verify the current operational state of the interface. Possible states include:
	+ `up(1)`: The interface is operational and ready to transmit and receive network traffic.
	+ `down(2)`: The interface is not operational and cannot transmit or receive network traffic.
	+ `noIfIdentifier(3)`: The interface does not have a valid Interface Identifier assigned, possibly due to Duplicate Address Detection failure.
	+ `notPresent(5)`: The interface is missing lower-layer components and cannot operate.

Check the device's logs and monitoring systems to determine the root cause of the state change. It may be related to a configuration change, hardware failure, or other issues.

## Mitigation

To mitigate the impact of this notification, follow these steps:

1. Verify the current operational state of the interface using `ipv6IfOperStatus`.
2. Check the device's logs and monitoring systems to determine the root cause of the state change.
3. Take corrective action based on the root cause:
	* If the issue is related to a configuration change, revert the change or apply a new configuration to restore the interface to its previous state.
	* If the issue is related to a hardware failure, replace the failed hardware component and restore the interface to its previous state.
	* If the issue is related to a Duplicate Address Detection failure, resolve the address conflict and restore the interface to its previous state.
4. Monitor the interface's operational status to ensure it has returned to a stable state.
5. Notify stakeholders of the issue and its resolution.
---




