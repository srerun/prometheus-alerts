---
title: spRelayArray3-4Status
description: Troubleshooting for SNMP trap spRelayArray3-4Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spRelayArray3-4Status 

RelayArray3.4 sensor trap 


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


Here is a sample runbook for the SNMP trap description:

## Meaning

The `SPAGENT-MIB::spRelayArray3-4Status` trap is generated when the RelayArray3.4 sensor has exceeded a predetermined threshold, indicating a potential issue with the sensor or the surrounding environment.

## Impact

This trap may indicate a critical condition that requires immediate attention. Failure to address the issue may lead to equipment damage, data loss, or even downtime.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `spSensorStatus` variable to determine the current status of the sensor.
2. Verify the `spSensorValue` variable to identify the current value of the sensor.
3. Compare the `spSensorLevelExceeded` variable to determine the threshold value that was exceeded.
4. Identify the sensor causing the trap using the `spSensorIndex` variable.
5. Consult the `spSensorName` and `spSensorDescription` variables to understand the purpose and location of the sensor.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate the cause of the sensor excursion, such as environmental factors (e.g., temperature, humidity) or equipment malfunction.
2. Take corrective action to address the underlying issue, such as adjusting environmental controls or replacing faulty equipment.
3. Verify that the sensor value has returned to a normal range.
4. Clear the trap to acknowledge that the issue has been resolved.
5. Perform additional monitoring to ensure the issue does not reoccur.

Remember to update the runbook according to your organization's specific requirements and procedures.
---




