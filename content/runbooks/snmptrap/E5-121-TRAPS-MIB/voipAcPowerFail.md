---
title: voipAcPowerFail
description: Troubleshooting for SNMP trap voipAcPowerFail
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-121-TRAPS-MIB::voipAcPowerFail 

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


Here is a runbook for the SNMP trap `E5-121-TRAPS-MIB::voipAcPowerFail`:

## Meaning

The `voipAcPowerFail` trap is triggered when a DC power error occurs, indicating a loss of power to the VoIP phone. This can result in the phone being unable to function correctly, leading to potential communication disruptions.

## Impact

The impact of this trap is high, as it can cause:

* Loss of phone functionality, resulting in inability to make or receive calls
* Disruption to critical communication services, potentially affecting business operations
* Increased risk of security breaches due to unavailability of secure communication channels

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `ifIndex` value to identify the specific interface affected by the power failure.
2. Verify the `voipPhoneState` to determine the current status of the phone.
3. Check the device's power supply and cabling to ensure they are functioning correctly.
4. Review system logs to identify any other related error messages or alerts.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately investigate and resolve the power failure issue to restore power to the affected phone.
2. Verify that the phone is functioning correctly and is able to make and receive calls.
3. If the issue persists, consider replacing the power supply or cabling as necessary.
4. Implement measures to prevent future occurrences, such as redundant power supplies or backup power sources.
5. Monitor the device's power supply and phone status to quickly identify and respond to any future power failures.
---




