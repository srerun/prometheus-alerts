---
title: spRelayArray1-2Status
description: Troubleshooting for SNMP trap spRelayArray1-2Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spRelayArray1-2Status 

RelayArray1.2 sensor trap 


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

The SPAGENT-MIB::spRelayArray1-2Status trap indicates that the RelayArray1.2 sensor has exceeded a predetermined threshold, triggering an alert. This sensor is likely monitoring environmental conditions, such as temperature, humidity, or air flow, and the threshold has been breached.

## Impact

The impact of this trap depends on the specific sensor and its role in the system. If the sensor is monitoring a critical component, such as a cooling system, exceeding the threshold could lead to:

* System downtime
* Data loss
* Equipment damage
* Decreased performance

## Diagnosis

To diagnose the issue, follow these steps:

1. **Identify the sensor**: Use the `spSensorName` and `spSensorDescription` variables to determine which sensor triggered the trap.
2. **Determine the threshold**: Use the `spSensorLevelExceeded` variable to determine the threshold that was breached.
3. **Check the sensor value**: Use the `spSensorValue` variable to determine the current reading of the sensor.
4. **Check the sensor status**: Use the `spSensorStatus` variable to determine the current status of the sensor.
5. **Review system logs**: Analyze system logs to identify any other related issues or errors.

## Mitigation

To mitigate the issue, follow these steps:

1. **Acknowledge the alert**: Acknowledge the trap to indicate that it has been received and is being investigated.
2. **Investigate the cause**: Identify the root cause of the threshold breach, such as a failed cooling fan or a blocked air vent.
3. **Take corrective action**: Take corrective action to address the issue, such as replacing a failed component or clearing a blockage.
4. **Verify sensor readings**: Verify that the sensor readings have returned to a safe range.
5. **Update thresholds**: Consider updating the thresholds for the sensor to prevent future false positives or negatives.
---




