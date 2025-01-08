---
title: NatsLeafNodeConnectionIssue
description: Troubleshooting for alert NatsLeafNodeConnectionIssue
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NatsLeafNodeConnectionIssue

No leaf node connections have been established in the last 5 minutes

<details>
  <summary>Alert Rule</summary>

{{% rule "nats/nats-exporter.yml" "NatsLeafNodeConnectionIssue" %}}

{{% comment %}}

```yaml
alert: NatsLeafNodeConnectionIssue
expr: increase(gnatsd_varz_leafnodes[5m]) == 0
for: 5m
labels:
    severity: critical
annotations:
    summary: Nats leaf node connection issue (instance {{ $labels.instance }})
    description: |-
        No leaf node connections have been established in the last 5 minutes
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nats-exporter/NatsLeafNodeConnectionIssue.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "NatsLeafNodeConnectionIssue":

## Meaning

The NatsLeafNodeConnectionIssue alert is triggered when no leaf node connections have been established in the last 5 minutes. This indicates a critical issue with the NATS cluster, as leaf nodes are responsible for connecting to the NATS server and publishing/subscribing to messages. Without leaf node connections, the NATS cluster is unable to function properly.

## Impact

The impact of this issue is high, as it can lead to:

* Disruption of message processing and delivery
* Increased latency and errors in the system
* Potential data loss or corruption
* Overall system instability and unavailability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NATS server logs for any errors or issues related to leaf node connections
2. Verify that the NATS exporter is functioning correctly and reporting metrics accurately
3. Investigate the network connectivity between the NATS server and the leaf nodes
4. Check the leaf node configuration and ensure that it is correct and up-to-date
5. Verify that there are no firewall or security restrictions blocking the connection between the NATS server and the leaf nodes

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the NATS server and/or the leaf nodes to re-establish connections
2. Check and update the NATS exporter configuration to ensure it is correct and up-to-date
3. Verify and update the network configuration to ensure connectivity between the NATS server and the leaf nodes
4. Check and update the leaf node configuration to ensure it is correct and up-to-date
5. Consider increasing the monitoring and logging verbosity to gather more information about the issue

Note: If the issue persists, consider escalating to a senior engineer or subject matter expert for further assistance.