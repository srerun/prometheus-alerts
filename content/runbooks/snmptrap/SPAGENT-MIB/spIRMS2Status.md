---
title: spIRMS2Status
description: Troubleshooting for SNMP trap spIRMS2Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spIRMS2Status 

IRMS sensor trap 


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


Here is a runbook for the SNMP trap # SPAGENT-MIB::spIRMS2Status:

## Meaning

The # SPAGENT-MIB::spIRMS2Status SNMP trap indicates that an IRMS (Intelligent Remote Management System) sensor has exceeded a predefined threshold, triggering an alert to notify administrators of a potential issue. The trap provides additional variables to help diagnose and mitigate the problem.

## Impact

The impact of this trap depends on the specific sensor and threshold exceeded. Possible effects include:

* Equipment malfunction or failure
* Service disruption or degradation
* Increased noise levels or vibration
* Safety risks to personnel or equipment
* Unplanned downtime or maintenance

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the sensor that triggered the trap using the `spSensorIndex` and `spSensorName` variables.
2. Check the `spSensorStatus` variable to determine the current status of the sensor (e.g., normal, warning, critical).
3. Examine the `spSensorValue` variable to determine the current reading of the sensor.
4. Compare the `spSensorValue` to the `spSensorLevelExceeded` variable to understand the threshold that was exceeded.
5. Consult the `spSensorDescription` variable for additional information about the sensor and its function.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate the root cause of the sensor exceeding the threshold (e.g., equipment malfunction, environmental factors).
2. Take immediate action to address the issue, such as:
	* Performing maintenance or repairs on the affected equipment.
	* Adjusting the sensor threshold or configuration.
	* Taking measures to reduce noise levels or vibration.
3. Verify that the issue is resolved by monitoring the sensor readings and checking for subsequent trap notifications.
4. Update the sensor configuration or threshold as needed to prevent similar issues in the future.
5. Document the incident and the steps taken to resolve it for future reference.
---




