---
title: netSnmpExampleHeartbeatNotification
description: Troubleshooting for SNMP trap netSnmpExampleHeartbeatNotification
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# NET-SNMP-EXAMPLES-MIB::netSnmpExampleHeartbeatNotification 

An example notification, used to illustrate the
definition and generation of trap and inform PDUs
(including the use of both standard and additional
varbinds in the notification payload).
This notification will typically be sent every
	 30 seconds, using the code found in the example module
agent/mibgroup/examples/notification.c 


## Variables


  - netSnmpExampleHeartbeatRate 

### Definitions 


**netSnmpExampleHeartbeatRate** 
: A simple integer object, to act as a payload for the
netSnmpExampleHeartbeatNotification.  The value has
no real meaning, but is nominally the interval (in
seconds) between successive heartbeat notifications. 


Here is a runbook for the SNMP trap description:

## Meaning

The NET-SNMP-EXAMPLES-MIB::netSnmpExampleHeartbeatNotification trap is an example notification used to illustrate the definition and generation of trap and inform PDUs. It is sent periodically, typically every 30 seconds, from the agent/mibgroup/examples/notification.c module.

## Impact

This trap has no significant impact on the network or devices. It is purely for demonstration purposes and does not indicate any errors or issues. However, if this trap is received unexpectedly or with an unusual frequency, it may indicate a misconfiguration or issue with the SNMP agent.

## Diagnosis

To diagnose the cause of this trap, check the following:

* Verify that the SNMP agent is configured correctly and is sending this trap as expected.
* Check the value of the `netSnmpExampleHeartbeatRate` variable to ensure it is within the expected range (every 30 seconds).
* Review the agent/mibgroup/examples/notification.c module code to ensure it is functioning correctly.

## Mitigation

No mitigation is required for this trap, as it is a normal and expected notification. However, if you are experiencing issues with the trap frequency or content, you may want to:

* Review and adjust the SNMP agent configuration to ensure it is sending the trap correctly.
* Debug the agent/mibgroup/examples/notification.c module code to identify and fix any issues.
* Consider disabling or modifying the trap if it is causing unnecessary noise or alerts.
---




