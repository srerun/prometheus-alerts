---
title: spRelayArray3-5Status
description: Troubleshooting for SNMP trap spRelayArray3-5Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spRelayArray3-5Status 

RelayArray3.5 sensor trap 


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


Here is a runbook for the SNMP Trap description:

## Meaning

The SPAGENT-MIB::spRelayArray3-5Status trap indicates that a sensor in the RelayArray3 module has exceeded a predefined threshold level, triggering an alert. This trap provides information about the sensor that triggered the alert, including its current status, value, and the level that was exceeded.

## Impact

This trap may indicate a potential issue with the RelayArray3 module or the environment it is monitoring. The impact can vary depending on the specific sensor and the threshold level that was exceeded. Possible impacts include:

* System downtime or failure
* Data loss or corruption
* Environmental issues (e.g. temperature, humidity)
* Equipment damage or wear

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the sensor that triggered the trap using the `spSensorIndex` and `spSensorName` variables.
2. Check the current status of the sensor using the `spSensorStatus` variable.
3. Review the current value of the sensor using the `spSensorValue` variable.
4. Determine the threshold level that was exceeded using the `spSensorLevelExceeded` variable.
5. Consult the `spSensorDescription` variable for more information about the sensor and the threshold level.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate the cause of the sensor value exceeding the threshold level.
2. Take corrective action to address the underlying issue (e.g. adjust environmental settings, replace faulty equipment).
3. Monitor the sensor value to ensure it returns to a safe range.
4. Update the threshold level if necessary to prevent future false positives or false negatives.
5. Consider implementing additional monitoring or alerting mechanisms to catch similar issues earlier.
---




