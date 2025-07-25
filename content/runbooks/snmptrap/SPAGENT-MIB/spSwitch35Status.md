---
title: spSwitch35Status
description: Troubleshooting for SNMP trap spSwitch35Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spSwitch35Status 

Switch sensor trap 


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


Here is a runbook for the SNMP trap `SPAGENT-MIB::spSwitch35Status`:

## Meaning

The `SPAGENT-MIB::spSwitch35Status` trap is a switch sensor trap that indicates a sensor on the switch has exceeded a specified level, triggering this trap to be sent. This trap provides information about the sensor that triggered the event, including its status, value, level exceeded, index, name, and description.

## Impact

The impact of this trap depends on the specific sensor that triggered the event. Possible impacts include:

* Increased temperature or humidity levels that could affect switch performance or longevity
* Power supply or fan failures that could cause switch downtime
* Physical security breaches or unauthorized access to the switch
* Environmental monitoring issues that could affect switch operation

## Diagnosis

To diagnose the cause of this trap, follow these steps:

1. Identify the sensor that triggered the trap using the `spSensorIndex` and `spSensorName` variables.
2. Check the `spSensorStatus` variable to determine the current status of the sensor.
3. Review the `spSensorValue` variable to understand the current reading of the sensor.
4. Compare the `spSensorValue` to the `spSensorLevelExceeded` variable to determine the threshold that was exceeded.
5. Consult the switch documentation or contact the manufacturer for more information on the specific sensor and its thresholds.
6. Verify that the sensor is functioning correctly and that the reading is accurate.

## Mitigation

To mitigate the issue, follow these steps:

1. Take immediate action to address the underlying issue that triggered the trap, such as adjusting the environment to reduce temperature or humidity levels.
2. Verify that the sensor is functioning correctly and that the reading is accurate.
3. Adjust the threshold level of the sensor if necessary to prevent future traps.
4. Consider implementing additional monitoring or alerting mechanisms to detect similar issues earlier.
5. Document the incident and the actions taken to resolve it for future reference.
---




