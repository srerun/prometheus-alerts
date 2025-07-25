---
title: spWattStatus
description: Troubleshooting for SNMP trap spWattStatus
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# SPAGENT-MIB::spWattStatus 

Energy sensor trap 


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


Here is a runbook for the SNMP Trap description:

## Meaning

The `SPAGENT-MIB::spWattStatus` trap indicates that an energy sensor has exceeded a certain level, triggering an alert. This trap is sent by the sensor agent to notify the network management system of the threshold breach.

## Impact

The impact of this trap depends on the specific sensor and the level that was exceeded. Potential consequences include:

* Increased energy consumption, leading to higher costs and potential overheating issues
* Equipment malfunction or failure due to excessive energy usage
* Inefficient use of resources, leading to reduced system performance or downtime

## Diagnosis

To diagnose the issue, gather the following information from the trap variables:

* `spSensorStatus`: The current status of the sensor
* `spSensorValue`: The current value of the sensor
* `spSensorLevelExceeded`: The level that was exceeded, triggering the trap
* `spSensorIndex`: The index of the sensor causing the trap
* `spSensorName`: The name of the sensor causing the trap
* `spSensorDescription`: The description of the sensor causing the trap

Use this information to identify the specific sensor and the threshold that was breached. Review the sensor's configuration and the surrounding environment to determine the root cause of the issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate the sensor and its surrounding environment to determine the cause of the threshold breach.
2. Take corrective action to reduce energy consumption, such as adjusting system settings, replacing components, or relocating equipment.
3. Verify that the sensor is configured correctly and accurately reporting energy usage.
4. Consider adjusting the threshold level for the sensor to prevent future breaches.
5. Monitor the sensor's performance and energy usage to ensure the issue is resolved and does not recur.
---




