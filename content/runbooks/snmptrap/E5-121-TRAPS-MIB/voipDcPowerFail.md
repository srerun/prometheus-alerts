---
title: voipDcPowerFail
description: Troubleshooting for SNMP trap voipDcPowerFail
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-121-TRAPS-MIB::voipDcPowerFail 

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


Here is a runbook for the SNMP trap `E5-121-TRAPS-MIB::voipDcPowerFail`:

## Meaning
The `E5-121-TRAPS-MIB::voipDcPowerFail` trap indicates that a DC power error has been triggered on a VoIP device. This trap is generated when the device detects a problem with its DC power supply, which may impact its ability to operate normally.

## Impact
The impact of this trap is that the VoIP device may not function properly or may become unavailable, leading to potential disruptions to voice communications. This can result in missed calls, dropped calls, or poor call quality, which can have a significant impact on business operations and customer satisfaction.

## Diagnosis
To diagnose the root cause of the `E5-121-TRAPS-MIB::voipDcPowerFail` trap, follow these steps:

* Check the VoIP device's power supply and cabling to ensure they are securely connected and functioning properly.
* Verify that the device is receiving the correct voltage and current.
* Check the device's logs for any error messages related to power supply issues.
* Use the `ifIndex` variable to identify the specific interface affected by the power failure.
* Check the `voipPhoneState` variable to determine the current state of the VoIP phone.

## Mitigation
To mitigate the impact of the `E5-121-TRAPS-MIB::voipDcPowerFail` trap, follow these steps:

* Replace the faulty power supply unit (PSU) or repair/replace the power cabling if necessary.
* Ensure that the VoIP device is properly configured and receiving the correct power settings.
* Implement redundant power supplies or backup power systems to minimize the risk of future power failures.
* Monitor the VoIP device's power status and logs regularly to quickly detect any potential power issues.
* Consider implementing a network monitoring system to detect and alert on power-related issues before they impact voice communications.
---




