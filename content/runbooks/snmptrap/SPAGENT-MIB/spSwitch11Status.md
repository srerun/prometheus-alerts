---
title: spSwitch11Status
description: Troubleshooting for SNMP trap spSwitch11Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spSwitch11Status 

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


Here is a runbook for the SNMP trap description:

## Meaning

The `SPAGENT-MIB::spSwitch11Status` trap indicates that a switch sensor has exceeded a threshold, triggering an alert. This trap provides information about the sensor that caused the trap, including its current status, value, and the level that was exceeded.

## Impact

The impact of this trap depends on the specific sensor and threshold that was exceeded. However, in general, this trap may indicate a potential issue with the switch or its environment, such as:

* Overheating: If the sensor is a temperature sensor, exceeding a threshold may indicate a cooling issue or a hardware fault.
* Power supply issue: If the sensor is a power supply sensor, exceeding a threshold may indicate a power supply failure or overload.
* Environmental issue: If the sensor is an environmental sensor (e.g. humidity, airflow), exceeding a threshold may indicate an environmental issue that could impact the switch's operation.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check the sensor details**: Review the trap information to identify the specific sensor that triggered the trap, including its name, description, index, and current value.
2. **Check the threshold**: Review the `spSensorLevelExceeded` value to determine the threshold that was exceeded.
3. **Check the switch logs**: Review the switch logs to see if there are any other error messages or events related to the sensor or switch.
4. **Perform a visual inspection**: If possible, perform a visual inspection of the switch and its environment to identify any signs of issues (e.g. overheating, power supply issues).

## Mitigation

To mitigate the issue, follow these steps:

1. **Take corrective action**: Based on the diagnosis, take corrective action to address the underlying issue. For example, if the sensor is a temperature sensor, ensure that the switch is in a well-ventilated area or consider replacing the switch if it is faulty.
2. **Adjust the threshold**: Consider adjusting the threshold value for the sensor to prevent future false positives or to make the trap more sensitive to actual issues.
3. **Monitor the switch**: Continue to monitor the switch and its sensors to ensure that the issue has been resolved and to identify any other potential issues.
---




