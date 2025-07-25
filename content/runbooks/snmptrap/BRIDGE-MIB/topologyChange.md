---
title: topologyChange
description: Troubleshooting for SNMP trap topologyChange
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# BRIDGE-MIB::topologyChange 

A topologyChange trap is sent by a bridge when any of
its configured ports transitions from the Learning state
to the Forwarding state, or from the Forwarding state to
the Blocking state.  The trap is not sent if a newRoot
trap is sent for the same transition.  Implementation of
this trap is optional. 



Here is a runbook for the SNMP trap description:

## Meaning

The BRIDGE-MIB::topologyChange trap is sent by a bridge when there is a change in the state of one of its configured ports. Specifically, this trap is triggered when a port transitions from the Learning state to the Forwarding state, or from the Forwarding state to the Blocking state. This trap is not sent if a newRoot trap is sent for the same transition.

## Impact

The impact of this trap depends on the specific network infrastructure and the role of the bridge in question. However, in general, a topology change can affect network connectivity and performance. For example:

* If a port transitions from the Forwarding state to the Blocking state, it may cause network traffic to be blocked or rerouted, leading to connectivity issues.
* If a port transitions from the Learning state to the Forwarding state, it may cause network traffic to be forwarded incorrectly, leading to performance issues.

## Diagnosis

To diagnose the cause of the topologyChange trap, perform the following steps:

1. Identify the bridge and port that triggered the trap.
2. Check the current state of the port using SNMP or other management tools.
3. Review system logs to determine the reason for the state change (e.g., link failure, configuration change, etc.).
4. Verify that the bridge is correctly configured and that the port is properly connected.

## Mitigation

To mitigate the impact of the topologyChange trap, perform the following steps:

1. Verify that the bridge is correctly configured and that the port is properly connected.
2. If the port transitioned to the Blocking state, investigate and resolve the underlying issue (e.g., link failure, configuration error, etc.).
3. If the port transitioned to the Forwarding state, verify that the bridge is correctly configured to forward traffic.
4. Consider implementing redundant links or paths to minimize the impact of topology changes on network connectivity.
5. Monitor the bridge and port closely to detect any further issues or changes.
---




