---
title: pimInterfaceElection
description: Troubleshooting for SNMP trap pimInterfaceElection
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# PIM-MIB::pimInterfaceElection 

A pimInterfaceElection notification signifies that a new DR
or DF has been elected on a network.
This notification is generated whenever the counter
pimInterfaceElectionWinCount is incremented, subject to the
rate limit specified by
pimInterfaceElectionNotificationPeriod. 


## Variables


  - pimInterfaceAddressType
  - pimInterfaceAddress 

### Definitions 


**pimInterfaceAddressType** 
: The address type of this PIM interface. 

**pimInterfaceAddress** 
: The primary IP address of this router on this PIM
interface.  The InetAddressType is given by the
pimInterfaceAddressType object. 


Here is a sample runbook for the PIM-MIB::pimInterfaceElection SNMP trap:

## Meaning

The PIM-MIB::pimInterfaceElection notification indicates that a new Designated Router (DR) or Designated Forwarder (DF) has been elected on a network. This event is triggered when the `pimInterfaceElectionWinCount` counter is incremented, subject to a rate limit specified by `pimInterfaceElectionNotificationPeriod`.

## Impact

This event may cause a change in the multicast traffic flow in the network, as the new DR or DF will take over the responsibility of forwarding multicast traffic. This can potentially lead to temporary disruptions or changes in network performance.

## Diagnosis

To diagnose the cause of this event, follow these steps:

1. Check the PIM interface configuration to identify the affected interface and the new DR or DF elected.
2. Verify the `pimInterfaceElectionWinCount` counter to confirm that it has been incremented.
3. Review the PIM protocol logs to understand the sequence of events leading up to the election.
4. Verify the `pimInterfaceAddressType` and `pimInterfaceAddress` variables to ensure they are correctly set for the affected interface.

## Mitigation

To mitigate the impact of this event, follow these steps:

1. Verify that the new DR or DF is correctly configured and operational.
2. Monitor network performance and multicast traffic flow to ensure that it is stable and correct.
3. Consider adjusting the `pimInterfaceElectionNotificationPeriod` variable to adjust the rate at which election notifications are sent.
4. Review PIM interface configurations and routing tables to ensure that they are correctly set and aligned with the new DR or DF election.

Note: The specific mitigation steps may vary depending on the network architecture and configuration.
---




