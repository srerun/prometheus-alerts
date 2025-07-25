---
title: voipClockFail
description: Troubleshooting for SNMP trap voipClockFail
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-121-TRAPS-MIB::voipClockFail 

Clock fault trigger. 


## Variables


  - voipDevId 

### Definitions 


**voipDevId** 
: The device ID of the DUT. 0 means port 1 to 24. 


Here is a runbook for the SNMP trap `E5-121-TRAPS-MIB::voipClockFail`:

## Meaning

The `E5-121-TRAPS-MIB::voipClockFail` trap indicates that a clock fault has been detected on a VoIP device. This fault may impact the device's ability to function properly, leading to disruptions in voice communication services.

## Impact

The impact of this trap is moderate to high, as it may cause issues with voice quality, call drops, or even prevent VoIP calls from being made or received. This can result in disruptions to business operations, emergency services, or other critical communications.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the affected VoIP device using the `voipDevId` variable, which indicates the device ID of the affected device (0 means port 1 to 24).
2. Check the device's clock synchronization status and configuration.
3. Verify that the device is receiving a valid clock signal from a reliable source.
4. Review device logs for any error messages related to clock faults or synchronization issues.
5. Perform a physical inspection of the device and its connections to ensure that there are no hardware issues or configuration problems.

## Mitigation

To mitigate the issue, follow these steps:

1. Reset the affected VoIP device to restore clock synchronization.
2. Verify that the device is configured to use a reliable clock source, such as a primary reference clock or a redundant clock source.
3. Check for any firmware or software updates that may address clock-related issues and apply them as necessary.
4. Perform a thorough inspection of the device's hardware and connections to ensure that there are no physical issues or configuration problems.
5. Monitor the device for further clock faults or synchronization issues and take prompt action to resolve any recurrence of the issue.

By following these steps, you should be able to diagnose and mitigate the `E5-121-TRAPS-MIB::voipClockFail` trap and restore normal operation of the affected VoIP device.
---




