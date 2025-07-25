---
title: spVirtual11Status
description: Troubleshooting for SNMP trap spVirtual11Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spVirtual11Status 

Virtual11 sensor trap 


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

The SPAGENT-MIB::spVirtual11Status trap is triggered when a virtual11 sensor detects an abnormal condition. This trap provides information about the sensor that caused the trap to be sent, including its current status, value, and the level that was exceeded.

## Impact

The impact of this trap depends on the specific sensor and the threshold that was exceeded. However, in general, this trap may indicate a potential issue with the system or device being monitored, such as overheating, high voltage, or other environmental factors. If left unaddressed, this issue could lead to system downtime, data loss, or equipment damage.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the sensor that triggered the trap by checking the `spSensorIndex` and `spSensorName` variables.
2. Check the `spSensorStatus` variable to determine the current status of the sensor.
3. Review the `spSensorValue` variable to see the current value of the sensor.
4. Determine the threshold that was exceeded by checking the `spSensorLevelExceeded` variable.
5. Consult the `spSensorDescription` variable to understand the significance of the sensor reading.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate the cause of the abnormal sensor reading.
2. Take corrective action to address the underlying issue, such as adjusting environmental settings or replacing faulty equipment.
3. Verify that the sensor reading has returned to a normal range.
4. Update the sensor configuration and thresholds as necessary to prevent future instances of this trap.
5. Document the issue and resolution in accordance with your organization's incident management procedures.
---




