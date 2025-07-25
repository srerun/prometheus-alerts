---
title: isisMaxAreaAddressesMismatch
description: Troubleshooting for SNMP trap isisMaxAreaAddressesMismatch
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisMaxAreaAddressesMismatch 

A notification sent when we receive a PDU
with a different value for the Maximum Area
Addresses.  This notification includes the
header of the packet, which may help a
network manager identify the source of the
confusion.
The agent must throttle the generation of
consecutive isisMaxAreaAddressesMismatch
notifications so that there is at least a 5-second
gap between notifications of this type.  When
notifications are throttled, they are dropped, not
queued for sending at a future time. 


## Variables


  - isisNotificationSysLevelIndex
  - isisPduMaxAreaAddress
  - isisNotificationCircIfIndex
  - isisPduFragment 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisPduMaxAreaAddress** 
: Holds the Max Area Addresses reported in a PDU
we received. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduFragment** 
: Holds up to 64 initial bytes of a PDU that
triggered the notification. 


## Meaning

The ISIS-MIB::isisMaxAreaAddressesMismatch trap is sent when the network device receives a PDU (Protocol Data Unit) with a different value for the Maximum Area Addresses. This notification helps network managers identify the source of the confusion by including the header of the packet that triggered the notification.

## Impact

The impact of this trap is:

* Inconsistent routing information: The mismatch in Maximum Area Addresses can cause routing inconsistencies, leading to network instability and potential outages.
* Network congestion: The device may drop packets or throttle notifications, leading to a delay in notification delivery, which can cause congestion in the network.
* Difficulty in troubleshooting: Without accurate routing information, it can be challenging to identify and troubleshoot issues in the network.

## Diagnosis

To diagnose the issue, follow these steps:

1. Verify the Maximum Area Addresses configured on the device: Check the device's configuration to ensure that the Maximum Area Addresses are set correctly.
2. Analyze the PDU header: Examine the header of the packet that triggered the notification to identify the source of the mismatch.
3. Check for duplicate notifications: Since the agent throttles consecutive notifications, check for duplicate notifications to ensure that the issue is not being masked.
4. Review network logs: Analyze network logs to identify any other related issues or errors that may be contributing to the problem.

## Mitigation

To mitigate the issue, follow these steps:

1. Update the device configuration: Update the device's configuration to reflect the correct Maximum Area Addresses.
2. Clear the notification: Clear the notification to ensure that it is not throttled and to allow for further troubleshooting.
3. Monitor the network: Closely monitor the network for any signs of routing inconsistencies or congestion.
4. Implement network redundancy: Implement redundant network paths to minimize the impact of routing inconsistencies.
5. Review and update network design: Review the network design and update it as necessary to prevent similar issues in the future.
---




