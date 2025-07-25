---
title: voipAcPowerClear
description: Troubleshooting for SNMP trap voipAcPowerClear
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-111-TRAPS-MIB::voipAcPowerClear 

DC power error release. 


## Variables


  - ifIndex
  - voipPhoneState 

### Definitions 


**ifIndex** 
: A unique value, greater than zero, for each interface.  It
is recommended that values are assigned contiguously
starting from 1.  The value for each interface sub-layer
must remain constant at least from one re-initialization of
the entity's network management system to the next re-
initialization. 

**voipPhoneState** 
: Phone status of the line. 


Here is a runbook for the SNMP trap E5-111-TRAPS-MIB::voipAcPowerClear:

## Meaning

The SNMP trap E5-111-TRAPS-MIB::voipAcPowerClear indicates that a DC power error has been cleared on a VoIP phone. This trap is triggered when the phone's power issue is resolved, and it is now operational again.

## Impact

The impact of this trap is low to moderate, as it indicates that a previous power issue with the VoIP phone has been resolved. However, it is essential to monitor the phone's status to ensure that it is functioning correctly and that there are no underlying issues that may cause further problems.

## Diagnosis

To diagnose the root cause of the original power issue, follow these steps:

1. Check the phone's power source to ensure it is stable and functioning correctly.
2. Verify that the phone's configuration is correct, including the power settings.
3. Review the phone's system logs to identify any errors or warnings related to the power issue.
4. Check the ifIndex variable to determine which interface was affected by the power issue.
5. Check the voipPhoneState variable to determine the current status of the phone.

## Mitigation

To mitigate the risk of future power issues on the VoIP phone, follow these steps:

1. Ensure that the phone's power source is reliable and stable.
2. Regularly check the phone's system logs for errors or warnings related to power issues.
3. Perform routine maintenance on the phone, including firmware updates and configuration backups.
4. Consider implementing redundant power sources or UPS systems to minimize the impact of power outages.
5. Monitor the phone's status and performance regularly to identify any potential issues before they become critical.
---




