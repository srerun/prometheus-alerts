---
title: ProviderFailedBecauseGetGenesisFailed
description: Troubleshooting for alert ProviderFailedBecauseGetGenesisFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ProviderFailedBecauseGetGenesisFailed

Failed to get genesis for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "ProviderFailedBecauseGetGenesisFailed" %}}

{{% comment %}}

```yaml
alert: ProviderFailedBecauseGetGenesisFailed
expr: eth_rpc_status == 2
for: 0m
labels:
    severity: critical
annotations:
    summary: Provider failed because get genesis failed (instance {{ $labels.instance }})
    description: |-
        Failed to get genesis for Provider `{{$labels.provider}}` in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/ProviderFailedBecauseGetGenesisFailed.md

```

{{% /comment %}}

</details>


Here is the runbook for the Prometheus alert rule:

## Meaning

The `ProviderFailedBecauseGetGenesisFailed` alert is triggered when the `eth_rpc_status` metric returns a value of 2, indicating that the provider failed to get the genesis block. This means that the Graph node is unable to retrieve the genesis block from the Ethereum blockchain, which is a critical component for the node's operation.

## Impact

The impact of this alert is severe, as it prevents the Graph node from functioning correctly. Without access to the genesis block, the node cannot synchronize with the blockchain, and data may become inconsistent or outdated. This can lead to:

* Inaccurate or incomplete data
* Node instability or crashes
* Disruption of dependent services or applications

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Graph node logs for errors related to getting the genesis block.
2. Verify that the Ethereum node is reachable and responding correctly.
3. Check the network connectivity between the Graph node and the Ethereum node.
4. Review the `eth_rpc_status` metric to identify the specific error code and message.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Graph node to attempt to re-initialize the connection to the Ethereum node.
2. Check the Ethereum node's status and ensure it is online and responding correctly.
3. Verify that the network connectivity between the Graph node and the Ethereum node is stable.
4. If the issue persists, investigate the underlying cause of the `eth_rpc_status` error and take corrective action.
5. Consider increasing the retry timeout or implementing a fallback mechanism to improve the resilience of the Graph node.

Remember to investigate the root cause of the issue and implement a permanent fix to prevent future occurrences.