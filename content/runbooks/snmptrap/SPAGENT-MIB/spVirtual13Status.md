---
title: spVirtual13Status
description: Troubleshooting for SNMP trap spVirtual13Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spVirtual13Status 

Virtual13 sensor trap 


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

The SPAGENT-MIB::spVirtual13Status trap indicates that a virtual sensor has exceeded a predefined threshold, triggering the trap to be sent. This trap provides information about the sensor that caused the trap, including its status, value, and the level that was exceeded.

## Impact

The impact of this trap depends on the specific sensor that triggered it and the threshold that was exceeded. Possible impacts include:

* Overheating: If the sensor measures temperature, exceeding a threshold could indicate a cooling issue that needs to be addressed.
* System downtime: If the sensor measures a critical system component, exceeding a threshold could indicate a potential failure that needs to be addressed to prevent system downtime.
* Performance degradation: If the sensor measures system performance, exceeding a threshold could indicate a performance issue that needs to be addressed.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the sensor that triggered the trap by checking the `spSensorIndex` and `spSensorName` variables.
2. Check the `spSensorStatus` variable to determine the current status of the sensor.
3. Check the `spSensorValue` variable to determine the current value of the sensor.
4. Check the `spSensorLevelExceeded` variable to determine the threshold that was exceeded.
5. Review system logs and performance metrics to determine the root cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Take corrective action to address the root cause of the issue, such as adjusting system settings or performing maintenance tasks.
2. Verify that the sensor value has returned to a normal range.
3. Clear the trap to acknowledge that the issue has been addressed.
4. Consider adjusting the threshold value for the sensor to prevent future traps from being sent unnecessarily.
5. Document the incident and the steps taken to mitigate the issue for future reference.
---




