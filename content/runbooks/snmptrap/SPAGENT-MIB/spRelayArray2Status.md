---
title: spRelayArray2Status
description: Troubleshooting for SNMP trap spRelayArray2Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spRelayArray2Status 

RelayArray2 sensor trap 


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


Here is a runbook for the SNMP trap `SPAGENT-MIB::spRelayArray2Status`:

## Meaning

The `SPAGENT-MIB::spRelayArray2Status` trap is generated when the RelayArray2 sensor detects an abnormal condition. This sensor is monitoring a specific aspect of the system, and the trap is sent when the sensor value exceeds a predefined threshold.

## Impact

The impact of this trap is that the system may be operating outside of normal parameters, which can lead to performance degradation, downtime, or even system failure. The specific impact will depend on the nature of the sensor and the threshold that has been exceeded.

## Diagnosis

To diagnose the issue, gather the following information from the trap:

* `spSensorStatus`: The current status of the sensor (e.g. OK, Warning, Critical)
* `spSensorValue`: The current value of the sensor
* `spSensorLevelExceeded`: The threshold value that was exceeded
* `spSensorIndex`: The index of the sensor that triggered the trap
* `spSensorName`: The name of the sensor that triggered the trap
* `spSensorDescription`: The description of the sensor that triggered the trap

Use this information to investigate the cause of the trap and determine the root cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate the cause of the trap and determine the root cause of the issue.
2. Take corrective action to address the issue, such as adjusting the sensor threshold or performing maintenance on the system.
3. Verify that the sensor value has returned to a normal range.
4. Clear the trap to indicate that the issue has been resolved.

Note: The specific mitigation steps will depend on the nature of the sensor and the issue that has been detected.
---




