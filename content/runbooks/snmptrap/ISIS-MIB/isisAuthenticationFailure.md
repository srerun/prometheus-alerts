---
title: isisAuthenticationFailure
description: Troubleshooting for SNMP trap isisAuthenticationFailure
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisAuthenticationFailure 

A notification sent when we receive a PDU
with an incorrect authentication information
field.  This notification includes the header
of the packet, which may help a network manager
identify the source of the confusion.
The agent must throttle the generation of
consecutive isisAuthenticationFailure
notifications so that there is at least a 5-second
gap between notifications of this type.  When
notifications are throttled, they are dropped, not
queued for sending at a future time. 


## Variables


  - isisNotificationSysLevelIndex
  - isisNotificationCircIfIndex
  - isisPduFragment 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduFragment** 
: Holds up to 64 initial bytes of a PDU that
triggered the notification. 


## Meaning

The ISIS-MIB::isisAuthenticationFailure trap is a notification sent when the device receives a Packet Data Unit (PDU) with incorrect authentication information. This trap includes the packet header, which can help network managers identify the source of the issue.

## Impact

The impact of this trap is significant, as it indicates a potential security breach or misconfiguration in the network. Incorrect authentication information can lead to unauthorized access, data tampering, or denial-of-service attacks. If left unaddressed, this issue can compromise the integrity and availability of network services.

## Diagnosis

To diagnose the issue, follow these steps:

1. **Check the packet header**: Analyze the packet header included in the trap to identify the source of the incorrect authentication information.
2. **Verify authentication configurations**: Review the authentication configurations on the device and adjacent nodes to ensure they are correct and consistent.
3. **Check for device or software issues**: Investigate if the device or software is experiencing issues that could be causing the incorrect authentication information.
4. **Review system logs**: Examine system logs to identify any other related error messages or anomalies.

Use the following variables to gather more information:

* `isisNotificationSysLevelIndex`: The system level for this notification.
* `isisNotificationCircIfIndex`: The identifier of the circuit relevant to this notification.
* `isisPduFragment`: Holds up to 64 initial bytes of the PDU that triggered the notification.

## Mitigation

To mitigate the issue, follow these steps:

1. **Correct authentication configurations**: Update the authentication configurations on the device and adjacent nodes to ensure they are correct and consistent.
2. **Disable or isolate the affected circuit**: Temporarily disable or isolate the affected circuit to prevent further unauthorized access or data tampering.
3. **Implement additional security measures**: Consider implementing additional security measures, such as access control lists (ACLs) or rate limiting, to prevent future authentication failures.
4. **Monitor the network**: Closely monitor the network for any signs of unauthorized access or suspicious activity.

Remember to throttle the generation of consecutive `isisAuthenticationFailure` notifications to prevent flooding and ensure a 5-second gap between notifications.
---




