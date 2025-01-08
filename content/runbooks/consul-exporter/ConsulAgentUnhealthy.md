---
title: ConsulAgentUnhealthy
description: Troubleshooting for alert ConsulAgentUnhealthy
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ConsulAgentUnhealthy

A Consul agent is down

<details>
  <summary>Alert Rule</summary>

{{% rule "consul/consul-exporter.yml" "ConsulAgentUnhealthy" %}}

{{% comment %}}

```yaml
alert: ConsulAgentUnhealthy
expr: consul_health_node_status{status="critical"} == 1
for: 0m
labels:
    severity: critical
annotations:
    summary: Consul agent unhealthy (instance {{ $labels.instance }})
    description: |-
        A Consul agent is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/consul-exporter/ConsulAgentUnhealthy.md

```

{{% /comment %}}

</details>


Here is the runbook for the ConsulAgentUnhealthy alert rule:

## Meaning

The ConsulAgentUnhealthy alert is triggered when a Consul agent is down and its health status is critical. This indicates that the Consul agent is not functioning properly, which can lead to issues with service discovery, health checking, and overall cluster stability.

## Impact

The impact of an unhealthy Consul agent can be significant, as it can:

* Disrupt service discovery and communication between services
* Lead to incorrect health checking and failure detection
* Cause issues with load balancing and traffic routing
* Impact the overall availability and reliability of the system

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Consul agent logs for errors or warnings that may indicate the cause of the failure.
2. Verify that the Consul agent is running and that its process is not terminated.
3. Check the network connectivity and firewall rules to ensure that the Consul agent can communicate with other nodes in the cluster.
4. Run the `consul members` command to verify that the node is still a member of the cluster.
5. Check the system resources (CPU, memory, disk space) to ensure that they are not exhausted.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the Consul agent service to attempt to recover from the failure.
2. Investigate and resolve any underlying issues that may have caused the agent to fail (e.g. network connectivity issues, resource exhaustion).
3. If the issue persists, consider redeploying the Consul agent or replacing the node if necessary.
4. Verify that the Consul cluster is still operating correctly and that services are still registering and deregistering properly.
5. Monitor the Consul agent health closely to ensure that the issue does not recur.