---
title: isisRejectedAdjacency
description: Troubleshooting for SNMP trap isisRejectedAdjacency
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# ISIS-MIB::isisRejectedAdjacency 

A notification sent when we receive a Hello
PDU from an IS but do not establish an
adjacency for some reason.
The agent must throttle the generation of
consecutive isisRejectedAdjacency notifications
so that there is at least a 5-second gap
between notifications of this type.  When
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


Here is a sample runbook for the SNMP trap `ISIS-MIB::isisRejectedAdjacency`:

## Meaning

The `isisRejectedAdjacency` trap is generated when the ISIS (Intermediate System to Intermediate System) protocol receives a Hello PDU from an Intermediate System (IS) but is unable to establish an adjacency for some reason. This trap indicates a problem with forming a neighbor relationship between ISIS routers, which is essential for exchanging routing information and maintaining network topology.

## Impact

 Failure to establish an adjacency can lead to incomplete or inaccurate network topology information, causing routing issues, and potentially leading to network instability or outages. This can result in poor network performance, packet loss, and decreased network reliability.

## Diagnosis

To diagnose the issue, collect the following information:

* The system level for this notification (`isisNotificationSysLevelIndex`)
* The identifier of the circuit relevant to this notification (`isisNotificationCircIfIndex`)
* Up to 64 initial bytes of the PDU that triggered the notification (`isisPduFragment`)

Analyze the collected information to identify the specific reason for the adjacency rejection. Possible causes include:

* Misconfigured ISIS router IDs or circuit IDs
* Incompatible ISIS protocol versions or implementations
* Network connectivity issues or hardware faults
* Resource constraints or overload on the affected routers

## Mitigation

To mitigate the issue, perform the following steps:

1. **Verify ISIS router configuration**: Check the configuration of the affected routers to ensure that ISIS router IDs and circuit IDs are correctly configured.
2. **Check network connectivity**: Verify that there are no network connectivity issues or hardware faults affecting the routers.
3. **Review ISIS protocol versions**: Ensure that all routers are running compatible ISIS protocol versions and implementations.
4. **Monitor resource utilization**: Check the resource utilization on the affected routers to identify any potential overload or constraint issues.
5. **Take corrective action**: Based on the diagnosis, take corrective action to resolve the underlying issue, such as reconfiguring the routers, replacing faulty hardware, or upgrading software.

By following these steps, you should be able to identify and address the root cause of the adjacency rejection, restoring normal network operation and preventing future occurrences of this trap.
---




