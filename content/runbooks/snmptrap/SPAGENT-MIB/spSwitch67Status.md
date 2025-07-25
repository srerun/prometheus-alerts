---
title: spSwitch67Status
description: Troubleshooting for SNMP trap spSwitch67Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spSwitch67Status 

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


Here is a sample runbook for the SNMP trap:

**SPAGENT-MIB::spSwitch67Status Switch Sensor Trap**

## Meaning

The SPAGENT-MIB::spSwitch67Status trap indicates that a switch sensor has exceeded a specified threshold, triggering an alert. This trap is generated when the sensor value crosses a predetermined level, indicating a potential issue with the switch.

## Impact

The impact of this trap can vary depending on the specific sensor and threshold exceeded. Possible consequences include:

* Overheating of the switch, which can lead to equipment failure or downtime
* Power supply issues, which can affect the overall network reliability
* Environmental factors, such as humidity or temperature, which can impact switch performance
* Reduced network performance or availability due to sensor-related issues

## Diagnosis

To diagnose the cause of the trap, follow these steps:

1. Identify the specific sensor that triggered the trap using the `spSensorIndex` and `spSensorName` variables.
2. Check the `spSensorValue` to determine the current reading of the sensor.
3. Compare the `spSensorValue` to the `spSensorLevelExceeded` to determine the threshold that was exceeded.
4. Review the `spSensorDescription` to understand the purpose and function of the sensor.
5. Consult the switch documentation and sensor specifications to determine the recommended course of action.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify the sensor reading and threshold values to ensure accuracy.
2. If the sensor reading indicates a critical issue (e.g., overheating), take immediate action to address the problem.
3. If the issue is environmental, take steps to stabilize the environment (e.g., adjust temperature or humidity).
4. If the issue is related to power supply, investigate and resolve any power-related problems.
5. Consider adjusting the sensor threshold or configured alerts to prevent false positives or unnecessary notifications.
6. Document the issue and resolution in the incident management system for future reference.
---




