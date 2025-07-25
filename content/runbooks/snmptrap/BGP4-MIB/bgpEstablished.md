---
title: bgpEstablished
description: Troubleshooting for SNMP trap bgpEstablished
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# BGP4-MIB::bgpEstablished 

The BGP Established event is generated when
the BGP FSM enters the ESTABLISHED state. 


## Variables


  - bgpPeerLastError
  - bgpPeerState 

### Definitions 


**bgpPeerLastError** 
: The last error code and subcode seen by this
peer on this connection.  If no error has
occurred, this field is zero.  Otherwise, the
first byte of this two byte OCTET STRING
contains the error code, and the second byte
contains the subcode. 

**bgpPeerState** 
: The BGP peer connection state. 


Here is a runbook for the SNMP Trap `BGP4-MIB::bgpEstablished`:

## Meaning

The BGP Established event is generated when the BGP Finite State Machine (FSM) enters the ESTABLISHED state. This indicates that the BGP peer connection has been successfully established and the router is exchanging routing information with its BGP neighbor.

## Impact

The impact of this event is typically none, as it indicates a normal and desired state of the BGP connection. However, it is still important to monitor and track these events to ensure that the BGP connection remains stable and that there are no underlying issues that may cause the connection to flap or drop.

## Diagnosis

To diagnose the cause of this event, check the following:

* Verify the BGP peer connection state using the `bgpPeerState` variable.
* Check the last error code and subcode seen by the peer on this connection using the `bgpPeerLastError` variable.
* Review the BGP configuration and verify that it is correct and consistent with the intended design.

## Mitigation

No mitigation is typically required for this event, as it indicates a normal and desired state of the BGP connection. However, if you are experiencing issues with the BGP connection, such as flapping or dropped connections, you may need to take additional steps to troubleshoot and resolve the underlying issue.
---




