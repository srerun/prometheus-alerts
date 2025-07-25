---
title: voipClockClear
description: Troubleshooting for SNMP trap voipClockClear
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-111-TRAPS-MIB::voipClockClear 

Clock fault release. 


## Variables


  - voipDevId 

### Definitions 


**voipDevId** 
: The device ID of the DUT. 0 means port 25 to 48, 1 means port 1 to 24. 


Here is a runbook for the SNMP Trap description:

## Meaning

The `E5-111-TRAPS-MIB::voipClockClear` trap indicates that a clock fault has been released on a specific device. This trap is generated when the device's clock fault condition is cleared, and normal operation is resumed.

## Impact

The impact of this trap is that the device is now functioning normally, and any disruptions caused by the clock fault are resolved. This trap is likely to be an informational message, alerting administrators that the issue has been resolved.

## Diagnosis

To diagnose the issue, administrators should:

* Verify that the device with ID `voipDevId` is functioning normally.
* Check the device's logs for any error messages related to the clock fault.
* Confirm that the clock fault has been cleared and the device is operating within normal parameters.

## Mitigation

No mitigation steps are required, as the clock fault has already been cleared. However, administrators should:

* Monitor the device for any further issues or errors.
* Perform routine maintenance and checks to ensure the device remains operational.
* Update any documentation or incident reports to reflect the resolution of the clock fault.
---




