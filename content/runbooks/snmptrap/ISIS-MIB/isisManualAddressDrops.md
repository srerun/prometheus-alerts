---
title: isisManualAddressDrops
description: Troubleshooting for SNMP trap isisManualAddressDrops
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisManualAddressDrops 

This notification is generated when one of the
manual areaAddresses assigned to this system is
ignored when computing routes.  The object
isisNotificationAreaAddress describes the area that
has been dropped.
The number of times this event has been generated
is counted by isisSysStatManAddrDropFromAreas.
The agent must throttle the generation of
consecutive isisManualAddressDrops notifications
so that there is at least a 5-second gap between
notifications of this type.  When notifications
are throttled, they are dropped, not queued for
sending at a future time. 


## Variables


  - isisNotificationAreaAddress 

### Definitions 


**isisNotificationAreaAddress** 
: An Area Address. 


Here is a runbook for the SNMP trap `ISIS-MIB::isisManualAddressDrops`:

## Meaning

The `ISIS-MIB::isisManualAddressDrops` trap is generated when one of the manual area addresses assigned to the system is ignored when computing routes. This means that the router is dropping an area address configured manually, which can impact the routing topology.

## Impact

The impact of this trap is that the routing topology may be affected, potentially leading to routing loops, black holes, or suboptimal routes. This can cause network instability, packet loss, and reduced network performance.

## Diagnosis

To diagnose the issue, check the `isisNotificationAreaAddress` variable to identify the specific area address that is being dropped. Then, investigate why the area address is being ignored, such as:

* Check the configuration of the manual area addresses to ensure they are correct and valid.
* Verify that the area address is not conflicting with other area addresses or routing configurations.
* Examine the routing tables and forwarding information bases (FIBs) to determine if there are any inconsistencies or issues.

## Mitigation

To mitigate the issue, perform the following steps:

* Review and correct the manual area address configuration to ensure it is valid and correct.
* Verify that the area address is not conflicting with other area addresses or routing configurations.
* Restart the IS-IS routing process or reload the router configuration to re-compute routes.
* Monitor the routing topology and verify that the area address is being used correctly.
* If the issue persists, consider increasing the logging level or enabling debug logs to gather more information about the issue.
---




