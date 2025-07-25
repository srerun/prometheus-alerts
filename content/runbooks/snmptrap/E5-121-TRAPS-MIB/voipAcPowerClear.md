---
title: voipAcPowerClear
description: Troubleshooting for SNMP trap voipAcPowerClear
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-121-TRAPS-MIB::voipAcPowerClear 

DC power error release. 


## Variables


  - ifIndex
  - voipPhoneState 

### Definitions 


**ifIndex** 
: A unique value, greater than zero, for each interface.  It
is recommended that values are assigned contiguously
starting from 1.  The value for each interface sub-layer
must remain constant at least from one re-initialization of
the entity's network management system to the next re-
initialization. 

**voipPhoneState** 
: Phone status of the line. 


Here is a runbook for the SNMP trap `E5-121-TRAPS-MIB::voipAcPowerClear`:

## Meaning

The `E5-121-TRAPS-MIB::voipAcPowerClear` trap indicates that a DC power error has been cleared on a VoIP phone. This trap is generated when the phone's power issue is resolved and normal operation is restored.

## Impact

The impact of this trap is minimal, as it indicates that a previous power error has been resolved and the phone is now functioning normally. However, it is still important to monitor the situation to ensure that the error does not recur.

## Diagnosis

To diagnose the issue, check the following:

* Verify that the VoIP phone is powered on and functioning correctly.
* Check the phone's power cable and connections to ensure they are secure.
* Review system logs to determine the cause of the original power error.
* Check the value of `voipPhoneState` to ensure that the phone is in a normal operating state.

## Mitigation

To mitigate the issue, follow these steps:

* Verify that the phone is properly configured and powered on.
* Ensure that all power cables and connections are secure.
* Monitor the phone's status and system logs to detect any recurring power errors.
* Consider implementing redundancy or backup power systems to minimize the impact of future power errors.
* Update or replace the phone's power supply if necessary.
---




