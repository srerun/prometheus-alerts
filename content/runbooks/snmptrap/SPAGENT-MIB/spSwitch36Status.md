---
title: spSwitch36Status
description: Troubleshooting for SNMP trap spSwitch36Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spSwitch36Status 

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

The SPAGENT-MIB::spSwitch36Status trap indicates that a switch sensor has exceeded a threshold, triggering an alert. This trap provides information about the sensor that exceeded the threshold, including its status, value, and the level that was exceeded.

## Impact

This trap may indicate a potential issue with the switch's environmental conditions, such as temperature, humidity, or voltage. If left unchecked, these conditions can lead to switch failure, downtime, or data loss.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `spSensorName` and `spSensorDescription` variables to determine the type of sensor that triggered the trap.
2. Check the `spSensorStatus` variable to determine the current status of the sensor.
3. Check the `spSensorValue` variable to determine the current value of the sensor.
4. Check the `spSensorLevelExceeded` variable to determine the threshold that was exceeded.
5. Check the switch's logs and sensor readings to determine if there are any trends or patterns indicating a larger issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify the sensor readings to ensure they are accurate and not a false alarm.
2. Check the switch's environmental conditions to ensure they are within normal operating ranges.
3. Take corrective action to address the underlying issue, such as adjusting the temperature or humidity in the switch's environment.
4. If the issue persists, consider replacing the sensor or the switch itself.
5. Update the switch's configuration to adjust the threshold levels or notification settings as needed to prevent future false alarms.
---




