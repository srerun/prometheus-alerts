---
title: spRelayArray8-3Status
description: Troubleshooting for SNMP trap spRelayArray8-3Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spRelayArray8-3Status 

RelayArray8.3 sensor trap 


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


Here is a runbook for the SNMP trap:

## Meaning

The SPAGENT-MIB::spRelayArray8-3Status SNMP trap indicates that the RelayArray8.3 sensor has exceeded a predetermined threshold, triggering an alarm. This trap is sent when the sensor status changes to an abnormal state, requiring attention from the operations team.

## Impact

The impact of this trap is that the RelayArray8.3 sensor is not functioning within normal parameters, which may affect the overall system performance, reliability, or availability. This could lead to service disruptions, data loss, or equipment damage if not addressed promptly.

## Diagnosis

To diagnose the issue, perform the following steps:

1. Check the **spSensorStatus** variable to determine the current status of the sensor (e.g., 0 = normal, 1 = warning, 2 = critical).
2. Review the **spSensorValue** variable to see the current reading of the sensor.
3. Check the **spSensorLevelExceeded** variable to determine the threshold level that was exceeded, triggering the trap.
4. Identify the sensor using the **spSensorIndex**, **spSensorName**, and **spSensorDescription** variables.
5. Verify the sensor's configuration and settings to ensure they are correct and up-to-date.
6. Check the system logs for any related errors or issues.

## Mitigation

To resolve the issue, perform the following steps:

1. Acknowledge the trap and notify the responsible teams (e.g., operations, maintenance).
2. Investigate the root cause of the issue, using the diagnosis steps above.
3. Take corrective action to return the sensor to a normal operating state (e.g., adjust settings, perform maintenance, replace the sensor).
4. Verify that the sensor is operating within normal parameters before closing the incident.
5. Update the sensor configuration and settings as needed to prevent similar issues in the future.
6. Document the incident and root cause analysis for future reference.
---




