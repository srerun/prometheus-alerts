---
title: spSwitch21Status
description: Troubleshooting for SNMP trap spSwitch21Status
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spSwitch21Status 

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


Here is a runbook for the SNMP trap `SPAGENT-MIB::spSwitch21Status`:

## Meaning

The `SPAGENT-MIB::spSwitch21Status` trap is sent when a switch sensor reaches a critical level, indicating a potential issue with the switch's environmental conditions. This trap is triggered when a sensor, such as a temperature or voltage sensor, exceeds a predefined threshold.

## Impact

The impact of this trap can be significant, as it may indicate a hardware failure or other environmental issue that can affect the switch's performance or stability. If left unchecked, this issue can lead to device failure, network downtime, or even data loss.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Identify the affected sensor**: Check the `spSensorName` and `spSensorDescription` variables to identify which sensor triggered the trap.
2. **Check the sensor status**: Verify the current status of the sensor using the `spSensorStatus` variable.
3. **Determine the threshold exceeded**: Check the `spSensorLevelExceeded` variable to determine the level that was exceeded, triggering the trap.
4. **Check the sensor value**: Verify the current value of the sensor using the `spSensorValue` variable.
5. **Review system logs**: Review system logs to identify any other related errors or issues.

## Mitigation

To mitigate the issue, follow these steps:

1. **Investigate the environmental conditions**: Investigate the environmental conditions surrounding the switch to identify the root cause of the issue.
2. **Take corrective action**: Take corrective action to address the environmental issue, such as adjusting the temperature or voltage levels.
3. **Verify sensor readings**: Verify that the sensor readings have returned to a safe level.
4. **Restart or replace the switch**: If necessary, restart or replace the switch to prevent further damage or downtime.
5. **Update monitoring thresholds**: Update monitoring thresholds to ensure that similar issues are caught earlier in the future.
---




