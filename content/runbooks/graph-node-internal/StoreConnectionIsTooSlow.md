---
title: StoreConnectionIsTooSlow
description: Troubleshooting for alert StoreConnectionIsTooSlow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# StoreConnectionIsTooSlow

Store connection is too slow to `{{$labels.pool}}` pool, `{{$labels.shard}}` shard in Graph node `{{$labels.instance}}`

<details>
  <summary>Alert Rule</summary>

{{% rule "graph-node/graph-node-internal.yml" "StoreConnectionIsTooSlow" %}}

{{% comment %}}

```yaml
alert: StoreConnectionIsTooSlow
expr: store_connection_wait_time_ms > 10
for: 0m
labels:
    severity: warning
annotations:
    summary: Store connection is too slow (instance {{ $labels.instance }})
    description: |-
        Store connection is too slow to `{{$labels.pool}}` pool, `{{$labels.shard}}` shard in Graph node `{{$labels.instance}}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/graph-node-internal/StoreConnectionIsTooSlow.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "StoreConnectionIsTooSlow":

## Meaning

The "StoreConnectionIsTooSlow" alert is triggered when the wait time for the store connection exceeds 10 milliseconds. This indicates that the connection to the store is slow, which can impact the performance of the Graph node.

## Impact

A slow store connection can cause:

* Delays in data processing and querying
* Increased latency in the Graph node
* Potential data loss or inconsistencies
* Impact on overall system performance and availability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the instance and pool/shard details in the alert labels to identify the affected component.
2. Verify the current store connection wait time using the `store_connection_wait_time_ms` metric.
3. Investigate recent changes in the store configuration, network connectivity, or resource utilization that may be contributing to the slow connection.
4. Review system logs for errors or warnings related to the store connection.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the store connection configuration and optimize it if necessary.
2. Verify that the network connectivity between the Graph node and the store is stable and not congested.
3. Consider increasing resources (e.g., CPU, memory) for the Graph node or store to improve performance.
4. If the issue persists, consider restarting the affected Graph node or store component.
5. If none of the above steps resolve the issue, escalate to a senior engineer or expert for further assistance.