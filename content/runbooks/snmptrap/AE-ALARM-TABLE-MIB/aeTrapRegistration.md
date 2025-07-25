---
title: aeTrapRegistration
description: Troubleshooting for SNMP trap aeTrapRegistration
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# AE-ALARM-TABLE-MIB::aeTrapRegistration 

aeTrapRegistration is generated periodically after the ONT boots.
It is sent more frequently before it is pinged and slows down afterwards. 


## Variables


  - aeTrapFsanSerialNumber
  - aeTrapIpAddr
  - aeTrapMacAddress
  - aeTrapOntState
  - aeTrapOntLabel
  - aeTrapRegistrationID
  - aeTrapDeviceClass
  - aeTrapDeviceModel
  - aeTrapFirmwareRevision
  - aeTrapConfigMethod
  - aeTrapConfigFilename
  - aeTrapConfigFileMarker
  - aeTrapConfigFileMIC
  - aeTrapPrimaryManagementServer
  - aeTrapDeviceStatus
  - aeTrapConfigStatus
  - aeTrapManagementVlanId
  - aeTrapSequenceNo
  - aeTrapMfgSerialNumber 

### Definitions 


**aeTrapFsanSerialNumber** 
: The FSAN Serial Number of the ONT expressed as 4 charaters and 8 hex digits. 

**aeTrapIpAddr** 
: The IP Address assigned to the ONT. 

**aeTrapMacAddress** 
: External MAC Address of the ONT. 

**aeTrapOntState** 
: ONT SNMP provisioning state. 

**aeTrapOntLabel** 
: The descriptive label assigned to the ONT. 

**aeTrapRegistrationID** 
: The Registration ID of the ONT expressed as max 10 char numerical string. 

**aeTrapDeviceClass** 
: Device class (ONT) 

**aeTrapDeviceModel** 
: Device model, i.e. ONT model number 710GX, etc 

**aeTrapFirmwareRevision** 
: This object identifies the firmware version
currently running on the ONT. 

**aeTrapConfigMethod** 
: ONT configuration method, i.e. config file, TR69, SNMP, etc 

**aeTrapConfigFilename** 
: ONT configuration filename in used 

**aeTrapConfigFileMarker** 
: ONT configuration filename marker 

**aeTrapConfigFileMIC** 
: ONT configuration filename message integrity code 

**aeTrapPrimaryManagementServer** 
: Primary management server IP address 

**aeTrapDeviceStatus** 
: Device status, i.e., acquired IP address, configured with ONT specific
file, generic file, cached file, etc 

**aeTrapConfigStatus** 
: Configuration status, i.e., no errors, configured with errors, not configured 

**aeTrapManagementVlanId** 
: Management VLAN ID 

**aeTrapSequenceNo** 
: Uniquely identifies each alarm trap that is transmitted by the ONT.
The value Increment for each alarm trap that is transmitted.
The first trap has a sequence number of one (1). 

**aeTrapMfgSerialNumber** 
: The MFG Serial Number of the ONT expressed as decimal digits. 


Here is a runbook for the SNMP Trap description `aeTrapRegistration`:

## Meaning

The `aeTrapRegistration` trap is generated periodically after an Optical Network Terminal (ONT) boots up. The trap is sent more frequently before the ONT is pinged and slows down afterwards. This trap provides information about the ONT's registration status, configuration, and device details.

## Impact

The `aeTrapRegistration` trap may cause a flood of SNMP traps, potentially overwhelming the network management system and impacting its performance. Additionally, if the trap is sent frequently, it may also cause issues with network congestion and bandwidth utilization.

## Diagnosis

To diagnose the `aeTrapRegistration` trap, follow these steps:

1. Check the ONT's boot sequence and ensure it is completing successfully.
2. Verify the ONT's configuration and registration status using the variables provided in the trap, such as `aeTrapRegistrationID`, `aeTrapIpAddr`, and `aeTrapOntState`.
3. Check the network management system's logs to identify any patterns or correlations with the trap generation.
4. Use network monitoring tools to analyze network traffic and bandwidth utilization.

## Mitigation

To mitigate the impact of the `aeTrapRegistration` trap, follow these steps:

1. Implement SNMP trap filtering or throttling to reduce the frequency of trap generation.
2. Configure the ONT to send traps at a slower rate or only when specific events occur.
3. Implement quality of service (QoS) policies to prioritize network traffic and ensure critical services are not affected.
4. Monitor network performance and adjust the mitigation strategies as needed to prevent network congestion and performance issues.
---




