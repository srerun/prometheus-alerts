---
title: isisIDLenMismatch
description: Troubleshooting for SNMP trap isisIDLenMismatch
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisIDLenMismatch 

A notification sent when we receive a PDU
with a different value for the System ID Length.
This notification includes an index to identify
the circuit where we saw the PDU and the header of
the PDU, which may help a network manager identify
the source of the confusion.
The agent must throttle the generation of
consecutive isisIDLenMismatch notifications
so that there is at least a 5-second gap between
notifications of this type.  When notifications
are throttled, they are dropped, not queued for
sending at a future time. 


## Variables


  - isisNotificationSysLevelIndex
  - isisPduFieldLen
  - isisNotificationCircIfIndex
  - isisPduFragment 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisPduFieldLen** 
: Holds the System ID length reported in PDU we received. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduFragment** 
: Holds up to 64 initial bytes of a PDU that
triggered the notification. 


Here is a runbook for the SNMP trap ISIS-MIB::isisIDLenMismatch:

## Meaning

The ISIS-MIB::isisIDLenMismatch trap indicates that the network device has received a Protocol Data Unit (PDU) with a different value for the System ID Length. This trap is sent to notify the network manager of a potential configuration issue or incompatibility in the network.

## Impact

The impact of this trap is that it may indicate a configuration mismatch or incompatibility between network devices, which can lead to issues with network routing and connectivity. If left unaddressed, this issue can cause network instability and affect network performance.

## Diagnosis

To diagnose the issue, follow these steps:

1. Review the trap details to identify the circuit (isisNotificationCircIfIndex) and the System ID length reported in the PDU (isisPduFieldLen).
2. Analyze the PDU fragment (isisPduFragment) to determine the source of the mismatch.
3. Verify the configuration of the network device that sent the trap to ensure that the System ID length is correctly configured.
4. Check other network devices in the same domain to ensure that they have consistent System ID length configurations.

## Mitigation

To mitigate the issue, follow these steps:

1. Correct the System ID length configuration on the network device that sent the trap.
2. Verify that all network devices in the same domain have consistent System ID length configurations.
3. Monitor the network for any further instances of this trap to ensure that the issue is fully resolved.
4. Consider implementing configuration validation and auditing procedures to prevent similar issues from occurring in the future.

Note: The agent will throttle the generation of consecutive isisIDLenMismatch notifications to ensure that there is at least a 5-second gap between notifications. This means that some notifications may be dropped and not queued for sending at a future time.
---




