---
title: CassandraNodeDown
description: Troubleshooting for alert CassandraNodeDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraNodeDown

Cassandra node down

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraNodeDown" %}}

{{% comment %}}

```yaml
alert: CassandraNodeDown
expr: sum(cassandra_stats{name="org:apache:cassandra:net:failuredetector:downendpointcount"}) by (service,group,cluster,env) > 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Cassandra node down (instance {{ $labels.instance }})
    description: |-
        Cassandra node down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraNodeDown.md

```

{{% /comment %}}

</details>


Here is a sample runbook for the Prometheus alert rule "CassandraNodeDown":

## Meaning

The CassandraNodeDown alert is triggered when one or more Cassandra nodes are down, indicating that the node(s) are not responding to requests. This alert is critical, as it can impact the overall availability and performance of the Cassandra cluster.

## Impact

The impact of this alert can be significant, as a downed Cassandra node can lead to:

* Data inconsistencies and potential data loss
* Increased latency and decreased performance for applications relying on Cassandra
* Reduced availability of the Cassandra cluster, potentially causing errors or downtime for dependent services

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cassandra node's system logs for any errors or exceptions that may indicate the cause of the downtime.
2. Verify that the node is properly configured and that there are no network connectivity issues.
3. Check the Cassandra cluster's overall health and performance metrics, such as latency and throughput, to identify any other potential issues.
4. Investigate any recent changes or deployments that may have caused the node to go down.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the downed Cassandra node to attempt to bring it back online.
2. If the node does not come back online, investigate the root cause of the issue and take corrective action (e.g., fix network connectivity issues, update configuration, etc.).
3. If the issue persists, consider replacing the node or scaling out the Cassandra cluster to maintain availability.
4. Notify relevant teams and stakeholders of the issue and the actions being taken to resolve it.

Remember to consult the Cassandra documentation and relevant internal resources for more detailed troubleshooting and resolution steps.