---
title: isisAreaMismatch
description: Troubleshooting for SNMP trap isisAreaMismatch
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisAreaMismatch 

A notification sent when we receive a Hello
PDU from an IS that does not share any
area address.  This notification includes
the header of the packet, which may help a
network manager identify the source of the
confusion.
The agent must throttle the generation of
consecutive isisAreaMismatch notifications
so that there is at least a 5-second gap
between notifications of this type.  When
notifications are throttled, they are dropped, not
queued for sending at a future time. 


## Variables


  - isisNotificationCircIfIndex
  - isisPduFragment 

### Definitions 


**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduFragment** 
: Holds up to 64 initial bytes of a PDU that
triggered the notification. 


Here is a runbook for the ISIS-MIB::isisAreaMismatch SNMP trap:

## Meaning

The ISIS-MIB::isisAreaMismatch notification is sent when the network device receives a Hello PDU from an Intermediate System (IS) that does not share any area address. This indicates a mismatch between the area addresses used by the local device and the remote IS.

## Impact

This notification can have a significant impact on the network, as it can lead to routing inconsistencies and potential network instability. If not addressed, it can cause:

* Routing loops
* Traffic congestion
* Network downtime
* Decreased network performance

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `isisNotificationCircIfIndex` variable to identify the circuit interface relevant to the notification.
2. Analyze the `isisPduFragment` variable to inspect the initial bytes of the PDU that triggered the notification.
3. Verify the area addresses used by the local device and the remote IS to identify the mismatch.
4. Check the network topology and configuration to identify any potential misconfigurations or inconsistencies.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and correct the mismatched area addresses used by the local device and the remote IS.
2. Update the network configuration to ensure consistency across all devices.
3. Verify that all devices in the network are using the correct area addresses.
4. Monitor the network for any signs of routing instability or performance issues.
5. Consider implementing additional measures to prevent similar issues in the future, such as regular network audits and configuration backups.

Note: Due to the throttling mechanism, consecutive notifications of this type may be dropped if they occur within a 5-second gap. Therefore, it is essential to respond promptly to this notification to prevent potential issues.
---




