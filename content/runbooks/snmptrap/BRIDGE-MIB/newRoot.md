---
title: newRoot
description: Troubleshooting for SNMP trap newRoot
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# BRIDGE-MIB::newRoot 

The newRoot trap indicates that the sending agent has
become the new root of the Spanning Tree; the trap is
sent by a bridge soon after its election as the new
root, e.g., upon expiration of the Topology Change Timer,
immediately subsequent to its election.  Implementation
of this trap is optional. 



## Meaning

The `BRIDGE-MIB::newRoot` SNMP trap indicates that the sending agent (a bridge or switch) has become the new root of the Spanning Tree. This trap is sent by the bridge soon after its election as the new root, typically triggered by the expiration of the Topology Change Timer.

## Impact

The impact of this trap being triggered is moderate, as it may indicate a change in the network topology. If not properly diagnosed and mitigated, this change may lead to:

* Network instability
* Packet loss or duplication
* Delayed or interrupted network services
* Inefficient use of network resources

## Diagnosis

To diagnose the cause of the `BRIDGE-MIB::newRoot` trap, follow these steps:

1. Verify the Spanning Tree topology using commands such as `show spanning-tree` or `show stp`.
2. Check the bridge's log for any errors or warnings related to Spanning Tree operation.
3. Use network management tools, such as network topology maps or packet sniffers, to monitor network traffic and identify any issues.
4. Verify that the bridge's configuration is correct and up-to-date.

## Mitigation

To mitigate the effects of the `BRIDGE-MIB::newRoot` trap, follow these steps:

1. Verify that the new root bridge is configured correctly and is functioning properly.
2. Ensure that all network devices are configured to use the correct Spanning Tree protocol and timers.
3. Monitor the network for any signs of instability or packet loss.
4. Consider implementing redundant links and/or redundant bridges to minimize the impact of topology changes.
5. Review and update network documentation to reflect the new Spanning Tree topology.
---




