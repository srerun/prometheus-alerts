---
title: isisVersionSkew
description: Troubleshooting for SNMP trap isisVersionSkew
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisVersionSkew 

A notification sent when we receive a Hello
PDU from an IS running a different version
of the protocol.  This notification includes
the header of the packet, which may help a
network manager identify the source of the
confusion.
The agent must throttle the generation of
consecutive isisVersionSkew notifications
so that there is at least a 5-second gap
between notifications of this type.  When
notifications are throttled, they are dropped, not
queued for sending at a future time. 


## Variables


  - isisNotificationSysLevelIndex
  - isisNotificationCircIfIndex
  - isisPduProtocolVersion
  - isisPduFragment 

### Definitions 


**isisNotificationSysLevelIndex** 
: The system level for this notification. 

**isisNotificationCircIfIndex** 
: The identifier of this circuit relevant to
this notification. 

**isisPduProtocolVersion** 
: Holds the Protocol version reported in PDU we received. 

**isisPduFragment** 
: Holds up to 64 initial bytes of a PDU that
triggered the notification. 


Here is a runbook for the ISIS-MIB::isisVersionSkew SNMP trap:

## Meaning

The ISIS-MIB::isisVersionSkew trap is generated when a network device receives a Hello PDU from another Intermediate System (IS) running a different version of the IS-IS protocol. This trap is sent to notify the network administrator of a potential configuration issue or incompatibility between devices.

## Impact

The impact of this trap is that it may indicate a problem with IS-IS protocol interoperability between devices, which can lead to connectivity issues or routing instability in the network. If left unaddressed, this issue can cause network outages, packet loss, or delays.

## Diagnosis

To diagnose the issue, the following steps can be taken:

1. **Check the protocol version**: Use the `isisPduProtocolVersion` variable to determine the protocol version reported in the PDU that triggered the trap.
2. **Identify the source of the issue**: Analyze the `isisPduFragment` variable to extract the initial bytes of the PDU that triggered the trap, which may help identify the source of the confusion.
3. **Verify device configurations**: Check the IS-IS protocol configurations on both devices to ensure they are running the same protocol version and compatible settings.
4. **Review network topology**: Verify that the network topology is correctly configured and that there are no issues with network connectivity or routing.

## Mitigation

To mitigate the issue, the following steps can be taken:

1. **Upgrade or downgrade IS-IS protocol version**: If the protocol version mismatch is the cause of the issue, upgrade or downgrade the IS-IS protocol version on the affected devices to match the version used by the other devices in the network.
2. **Configure compatible settings**: Ensure that the IS-IS protocol settings are configured consistently across all devices in the network.
3. **Verify network topology**: Re-verify the network topology to ensure that it is correctly configured and that there are no issues with network connectivity or routing.
4. **Monitor network performance**: Continuously monitor network performance to detect any potential issues or degradation in network connectivity or routing.
---




