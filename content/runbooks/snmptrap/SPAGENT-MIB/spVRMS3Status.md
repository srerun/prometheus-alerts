---
title: spVRMS3Status
description: Troubleshooting for SNMP trap spVRMS3Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spVRMS3Status 

VRMS sensor trap 


## Variables


  - spSensorStatus
  - spSensorValue
  - spSensorLevelExceeded
  - spSensorIndex
  - spSensorName
  - spSensorDescription 

### Definitions 


**spSensorStatus** 
: The current integer status of the sensor causing this trap to be sent 

**spSensorValue** 
: The current integer value of the sensor causing this trap to be sent 

**spSensorLevelExceeded** 
: The integer level that was exceeded causing this trap to be sent 

**spSensorIndex** 
: The integer index of the sensor causing this trap to be sent 

**spSensorName** 
: The name of the sensor causing this trap to be sent 

**spSensorDescription** 
: The description of the sensor causing this trap to be sent 


Here is a runbook for the SNMP Trap `SPAGENT-MIB::spVRMS3Status`:

## Meaning

The `SPAGENT-MIB::spVRMS3Status` trap is generated when a VRMS (Voltage Regulator Module) sensor exceeds a predefined threshold. This trap indicates that the voltage regulator module has exceeded a critical level, which may affect the system's stability and performance.

## Impact

The impact of this trap is potentially severe, as it may indicate a hardware failure or malfunction. If left unaddressed, it could lead to system crashes, data loss, or even physical damage to the equipment.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `spSensorStatus` variable to determine the current status of the sensor.
2. Review the `spSensorValue` variable to understand the current reading of the sensor.
3. Evaluate the `spSensorLevelExceeded` variable to determine the threshold that was exceeded.
4. Identify the affected sensor using the `spSensorIndex`, `spSensorName`, and `spSensorDescription` variables.
5. Verify the VRMS sensor configuration and settings to ensure they are correct and functioning as expected.
6. Check the system logs for any related errors or warnings.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately investigate the affected VRMS sensor to determine the cause of the threshold exceedance.
2. Take corrective action to return the VRMS sensor to a safe operating range. This may involve adjusting the sensor configuration, replacing the sensor, or performing maintenance on the affected hardware.
3. Verify that the system is stable and functioning as expected after resolving the issue.
4. Update the system logs to reflect the resolution of the issue.
5. Consider implementing additional monitoring and alerting to ensure that similar issues are quickly identified and addressed in the future.
---




