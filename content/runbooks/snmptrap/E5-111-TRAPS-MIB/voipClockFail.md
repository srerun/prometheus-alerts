---
title: voipClockFail
description: Troubleshooting for SNMP trap voipClockFail
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-111-TRAPS-MIB::voipClockFail 

Clock fault trigger. 


## Variables


  - voipDevId 

### Definitions 


**voipDevId** 
: The device ID of the DUT. 0 means port 25 to 48, 1 means port 1 to 24. 


Here is a runbook for the SNMP trap "E5-111-TRAPS-MIB::voipClockFail":

## Meaning

The "E5-111-TRAPS-MIB::voipClockFail" trap indicates that a clock fault has been triggered on a VoIP device. This fault can cause issues with call quality, synchronization, and overall VoIP service reliability.

## Impact

The impact of this trap is high, as a clock fault can cause:

* Call quality issues, such as jitter, latency, and dropped calls
* Inaccurate call billing and logging
* Synchronization issues with other network devices
* Downtime and outages of VoIP services

## Diagnosis

To diagnose the issue, follow these steps:

1. **Identify the affected device**: Use the `voipDevId` variable to determine which device is experiencing the clock fault. Check the device ID to determine which port range is affected (0 for port 25-48, 1 for port 1-24).
2. **Check device logs**: Review the device logs to determine when the clock fault occurred and if there are any related error messages.
3. **Verify clock synchronization**: Check if the device is properly synchronized with an external clock source or if the clock is running in a standalone mode.
4. **Check for hardware issues**: Perform a visual inspection of the device to identify any signs of hardware failure, such as overheating or physical damage.

## Mitigation

To mitigate the issue, follow these steps:

1. **Reset the device**: Attempt to reset the device to restore clock synchronization.
2. **Replace the device**: If the issue persists, replace the faulty device to prevent further disruptions to VoIP services.
3. **Verify clock configuration**: Verify that the clock configuration is correct and that the device is properly synchronized with an external clock source.
4. **Monitor for recurrences**: Continuously monitor the device for recurrences of the clock fault and take proactive measures to prevent future occurrences.
---




