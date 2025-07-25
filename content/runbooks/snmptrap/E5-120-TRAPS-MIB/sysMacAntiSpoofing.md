---
title: sysMacAntiSpoofing
description: Troubleshooting for SNMP trap sysMacAntiSpoofing
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# E5-120-TRAPS-MIB::sysMacAntiSpoofing 

MAC Anti-spoofing. 


## Variables


  - sysMacAntiSpoofOrig
  - sysMacAntiSpoofNew
  - sysMacAntiSpoofMAC 

### Definitions 


**sysMacAntiSpoofOrig** 
: The Original port of Mac-AntiSpoofing. 

**sysMacAntiSpoofNew** 
: The New port of Mac-AntiSpoofing. 

**sysMacAntiSpoofMAC** 
: The MAC of Mac-AntiSpoofing. 


Here is a runbook for the SNMP trap description:

## Meaning

The E5-120-TRAPS-MIB::sysMacAntiSpoofing trap indicates that the MAC anti-spoofing feature has detected a MAC address spoofing attempt on the network. This trap is generated when the system detects a MAC address that is different from the one previously learned on a port.

## Impact

The impact of this trap is that a potential security threat has been detected on the network. MAC address spoofing can be used to launch various types of attacks, including man-in-the-middle attacks, denial-of-service attacks, and unauthorized access to network resources. If left unchecked, this could lead to a security breach and potential data loss or compromise.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the value of `sysMacAntiSpoofOrig` to determine the original port where the MAC address was learned.
2. Check the value of `sysMacAntiSpoofNew` to determine the new port where the MAC address is being spoofed.
3. Check the value of `sysMacAntiSpoofMAC` to determine the MAC address that is being spoofed.
4. Investigate the device connected to the port specified in `sysMacAntiSpoofNew` to determine the source of the spoofed MAC address.
5. Review network logs and traffic captures to determine the extent of the spoofing activity.

## Mitigation

To mitigate the issue, follow these steps:

1. Isolate the port specified in `sysMacAntiSpoofNew` to prevent further unauthorized access.
2. Identify and remove the device that is spoofing the MAC address.
3. Implement additional security measures, such as MAC address filtering or port security, to prevent future MAC address spoofing attempts.
4. Update the MAC address table to reflect the correct MAC address-port mapping.
5. Monitor the network for further signs of suspicious activity and adjust security policies as needed.
---




