---
title: voipAcPowerFail
description: Troubleshooting for SNMP trap voipAcPowerFail
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-111-TRAPS-MIB::voipAcPowerFail 

DC power error trigger. 


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


Here is a runbook for the SNMP trap `E5-111-TRAPS-MIB::voipAcPowerFail`:

## Meaning

The `E5-111-TRAPS-MIB::voipAcPowerFail` SNMP trap indicates that a DC power error has been triggered on a VoIP phone. This trap is generated when the phone's AC power supply fails, which can impact phone operations and availability.

## Impact

The impact of this trap is that the VoIP phone may not be functioning properly or may be unavailable due to the power failure. This can lead to disruption of voice communications and affect business operations. The severity of the impact depends on the criticality of the phone and the services it provides.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the phone's power supply status using the `voipPhoneState` variable to determine if the phone is operational.
2. Verify the interface status using the `ifIndex` variable to ensure that the interface is not down or experiencing issues.
3. Check the phone's logs for any error messages related to power failure.
4. Perform a physical inspection of the phone and its power supply to identify any signs of damage or malfunction.

## Mitigation

To mitigate the issue, follow these steps:

1. Resolve the power supply issue by replacing the faulty power supply unit or resolving any electrical issues.
2. Verify that the phone's power supply is functioning properly and the phone is operational.
3. Test the phone's functionality to ensure that it is working as expected.
4. If the issue persists, contact the phone manufacturer or vendor for further assistance or replace the phone if necessary.
---




