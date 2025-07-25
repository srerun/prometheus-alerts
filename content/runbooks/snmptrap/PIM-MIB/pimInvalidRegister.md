---
title: pimInvalidRegister
description: Troubleshooting for SNMP trap pimInvalidRegister
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# PIM-MIB::pimInvalidRegister 

A pimInvalidRegister notification signifies that an invalid
PIM Register message was received by this device.
This notification is generated whenever the counter
pimInvalidRegisterMsgsRcvd is incremented, subject to the
rate limit specified by
pimInvalidRegisterNotificationPeriod. 


## Variables


  - pimGroupMappingPimMode
  - pimInvalidRegisterAddressType
  - pimInvalidRegisterOrigin
  - pimInvalidRegisterGroup
  - pimInvalidRegisterRp 

### Definitions 


**pimGroupMappingPimMode** 
: The PIM mode to be used for groups in this group prefix. 

**pimInvalidRegisterAddressType** 
: The address type stored in pimInvalidRegisterOrigin,
pimInvalidRegisterGroup, and pimInvalidRegisterRp.
If no invalid Register messages have been received, then
this object is set to unknown(0). 

**pimInvalidRegisterOrigin** 
: The source address of the last invalid Register message
received by this device. 

**pimInvalidRegisterGroup** 
: The IP multicast group address to which the last invalid
Register message received by this device was addressed. 

**pimInvalidRegisterRp** 
: The RP address to which the last invalid Register message
received by this device was delivered. 


Here is a runbook for the SNMP trap # PIM-MIB::pimInvalidRegister:

## Meaning

The `pimInvalidRegister` trap indicates that an invalid PIM Register message was received by the device. This trap is generated when the `pimInvalidRegisterMsgsRcvd` counter is incremented, subject to a rate limit specified by `pimInvalidRegisterNotificationPeriod`. This trap suggests that there is an issue with the PIM configuration or the device receiving invalid PIM Register messages.

## Impact

The impact of this trap is that invalid PIM Register messages are being received by the device, which can cause issues with multicast routing and forwarding. This can lead to:

* Inefficiencies in multicast routing and forwarding
* Potential security risks due to malformed or malicious PIM messages
* Increased network latency and packet loss

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `pimInvalidRegisterAddressType` variable to determine the address type of the invalid Register message.
2. Check the `pimInvalidRegisterOrigin` variable to determine the source address of the invalid Register message.
3. Check the `pimInvalidRegisterGroup` variable to determine the IP multicast group address to which the invalid Register message was addressed.
4. Check the `pimInvalidRegisterRp` variable to determine the RP address to which the invalid Register message was delivered.
5. Check the `pimGroupMappingPimMode` variable to determine the PIM mode used for the group prefix.
6. Review the PIM configuration and ensure that it is correct and consistent across all devices in the network.
7. Check for any security policies or firewall rules that may be blocking or modifying PIM messages.

## Mitigation

To mitigate the issue, follow these steps:

1. Verify the PIM configuration and ensure it is correct and consistent across all devices in the network.
2. Check for any security policies or firewall rules that may be blocking or modifying PIM messages and adjust them accordingly.
3. Implement measures to prevent invalid PIM Register messages from being sent or received, such as configuring PIM message filtering or rate limiting.
4. Consider implementing additional security measures, such as encrypting PIM messages or using authentication mechanisms to validate PIM messages.
5. Monitor the `pimInvalidRegisterMsgsRcvd` counter to track the frequency of invalid PIM Register messages and adjust the mitigation measures as needed.
---




