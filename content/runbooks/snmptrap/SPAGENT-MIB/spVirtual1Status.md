---
title: spVirtual1Status
description: Troubleshooting for SNMP trap spVirtual1Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spVirtual1Status 

Virtual1 sensor trap 


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

The SPAGENT-MIB::spVirtual1Status trap indicates that a virtual sensor has triggered an alert. This trap is sent when the sensor value exceeds a configured threshold, indicating a potential issue with the system or device being monitored.

## Impact

The impact of this trap can be significant, as it may indicate a critical issue with the system or device, such as overheating, power loss, or other environmental factors. If left unaddressed, these issues can lead to system downtime, data loss, or even permanent damage to hardware.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the sensor that triggered the trap by checking the `spSensorName` and `spSensorDescription` variables.
2. Check the `spSensorStatus` variable to determine the current status of the sensor.
3. Review the `spSensorValue` variable to determine the current value of the sensor.
4. Verify the threshold value that was exceeded by checking the `spSensorLevelExceeded` variable.
5. Check the system or device logs for any error messages or other indicators of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Take immediate action to address the issue indicated by the sensor, such as shutting down the system or device if necessary.
2. Verify that the sensor is functioning correctly and that the threshold value is set appropriately.
3. Check the system or device for any other issues that may be contributing to the problem.
4. Take steps to prevent similar issues from occurring in the future, such as adjusting the threshold value or implementing additional monitoring or alerting mechanisms.
5. Clear the trap and reset the sensor status once the issue has been resolved.
---




