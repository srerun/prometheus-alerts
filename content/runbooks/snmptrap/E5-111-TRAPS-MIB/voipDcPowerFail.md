---
title: voipDcPowerFail
description: Troubleshooting for SNMP trap voipDcPowerFail
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-111-TRAPS-MIB::voipDcPowerFail 

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


Here is a runbook for the SNMP trap `E5-111-TRAPS-MIB::voipDcPowerFail`:

## Meaning

The `E5-111-TRAPS-MIB::voipDcPowerFail` trap indicates that a DC power error has been triggered on a VoIP device. This means that the device has lost power or is experiencing a power-related issue, which may affect its ability to function properly.

## Impact

The impact of this trap is that VoIP phone service may be disrupted or unavailable, leading to potential communication outages and disruptions to business operations. This may result in lost productivity, missed calls, and reputational damage.

## Diagnosis

To diagnose the issue, check the following:

* Verify the DC power supply to the VoIP device is stable and functioning correctly.
* Check the device's logs for any error messages related to power failure or electrical issues.
* Use the `ifIndex` variable to identify the specific interface associated with the trap and investigate any issues with that interface.
* Use the `voipPhoneState` variable to determine the current state of the phone line and investigate any issues that may be related to the phone's status.

## Mitigation

To mitigate the issue, follow these steps:

* Check the DC power supply and replace it if necessary.
* Verify that the device is properly configured and that all necessary connections are secure.
* If the issue is related to a specific interface, check the configuration of that interface and make any necessary adjustments.
* If the issue is related to the phone line, check the phone's configuration and make any necessary adjustments to restore service.
* Consider implementing redundant power supplies or backup power sources to minimize the risk of future power-related outages.
---




