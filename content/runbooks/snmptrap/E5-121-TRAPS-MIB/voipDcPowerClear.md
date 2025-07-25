---
title: voipDcPowerClear
description: Troubleshooting for SNMP trap voipDcPowerClear
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-121-TRAPS-MIB::voipDcPowerClear 

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


Here is a runbook for the SNMP trap `E5-121-TRAPS-MIB::voipDcPowerClear`:

## Meaning

This SNMP trap indicates that a DC power error has been cleared on a VoIP device. This trap is triggered when the device's DC power issue is resolved, and normal operation has been restored.

## Impact

The impact of this trap being triggered is minimal, as it indicates that a previous DC power error has been cleared. However, it may be an indication that there was a recent power-related issue with the device that has since been resolved.

## Diagnosis

To diagnose the cause of this trap, follow these steps:

1. Check the device's power supply unit (PSU) to ensure it is functioning properly.
2. Verify that the DC power cable is securely connected to both the device and the power source.
3. Review system logs to determine the cause of the original DC power error.
4. Check the `ifIndex` variable to identify the specific interface that was affected by the power error.
5. Verify the `voipPhoneState` variable to ensure that the phone is operational and functioning correctly.

## Mitigation

To mitigate the risk of future DC power errors, follow these steps:

1. Ensure that the device is installed in a well-ventilated area to prevent overheating.
2. Regularly inspect the DC power cable for signs of wear or damage.
3. Consider replacing the PSU if it is old or has a history of failures.
4. Implement redundant power supplies to ensure continuous operation in the event of a power failure.
5. Monitor system logs and device status to quickly identify and respond to any future power-related issues.
---




