---
title: spCustomStatus
description: Troubleshooting for SNMP trap spCustomStatus
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spCustomStatus 

Custom sensor trap 


## Variables


  - spSensorStatus
  - spSensorValue
  - spSensorLevelExceeded
  - spSensorIndex
  - spSensorName
  - spSensorDescription
  - spSensorType
  - spSensorStatusName
  - spSensorSubIndex
  - spBoardIndex
  - spBoardDescription
  - spEventTimeStamp
  - spEventClassNumber
  - spEventClassName 

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

**spSensorType** 
: Type of this sensor 

**spSensorStatusName** 
: The current status of the sensor causing this trap to be sent 

**spSensorSubIndex** 
: The integer sub index of the sensor causing this trap to be sent. Only for temperaturearray and relayarray 

**spBoardIndex** 
: The integer index of the sensor board causing this trap to be sent. 

**spBoardDescription** 
: The description of the sensor board causing this trap to be sent 

**spEventTimeStamp** 
: The time(MM/DD/YYYY HH:MM:SS) of an event causing this trap to be sent 

**spEventClassNumber** 
: The user-defined class number associated with this trap 

**spEventClassName** 
: The user-defined class Name associated with this trap 


Here is a runbook for the SPAGENT-MIB::spCustomStatus SNMP trap:

## Meaning

The SPAGENT-MIB::spCustomStatus trap indicates that a custom sensor has triggered an event, causing this trap to be sent. This trap provides detailed information about the sensor, including its current status, value, and type, as well as the level that was exceeded to trigger the event.

## Impact

The impact of this trap depends on the specific sensor and the event that triggered it. However, in general, it may indicate a critical condition that requires immediate attention, such as a temperature or voltage threshold being exceeded. Failure to respond to this trap may result in equipment failure, data loss, or even physical damage.

## Diagnosis

To diagnose the cause of this trap, follow these steps:

1. Identify the sensor that triggered the event by examining the `spSensorName`, `spSensorDescription`, and `spSensorType` variables.
2. Determine the current status and value of the sensor by examining the `spSensorStatus` and `spSensorValue` variables.
3. Check the `spSensorLevelExceeded` variable to determine the level that was exceeded to trigger the event.
4. Verify the sensor board index and description using the `spBoardIndex` and `spBoardDescription` variables.
5. Check the event timestamp using the `spEventTimeStamp` variable to determine when the event occurred.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately investigate the cause of the event and take corrective action to prevent further damage or data loss.
2. Verify that the sensor is functioning correctly and that the event was not a false alarm.
3. Take steps to prevent the event from occurring again in the future, such as adjusting the sensor threshold or implementing additional monitoring and alerting mechanisms.
4. Update the event class number and name associated with this trap using the `spEventClassNumber` and `spEventClassName` variables to ensure that similar events are properly categorized and responded to in the future.
---




